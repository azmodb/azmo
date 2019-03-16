#include <sys/queue.h>
#include "libc.h"
#include "tree.h"

enum {
	MAXDATASIZE = 2 * 1024 * 1024, /* Maximum read and write fcall payload */
	MAXNAMESIZE = (1 << 16) - 1,
	HEADERSIZE  = 24,

	MAXMESSAGESIZE = HEADERSIZE + MAXDATASIZE,
};

typedef struct request_t request_t;
typedef struct queue_t   queue_t;
typedef struct fid_t     fid_t;

struct request_t {
	TAILQ_ENTRY(request_t) entry; /* reserved to threadpool */

	unsigned char* data; /* request/response buffer */
	uint32_t       msize;
	fcall_t*       tx;
	fcall_t*       rx;
};

struct queue_t {
	pthread_mutex_t mutex; /* used for read/write access */
	pthread_cond_t  cond;
	TAILQ_HEAD(rq, request_t) head;

	int size;     /* number of requests in queue */
	int shutdown; /* server has told us to stop */
};

struct fid_t {
	RB_ENTRY(fid_t) entry;

	uint32_t num;
};

static request_t*
request_create(uint32_t msize)
{
	request_t* req = xmalloc(sizeof(request_t));

	req->tx    = xmalloc(sizeof(fcall_t));
	req->rx    = xmalloc(sizeof(fcall_t));
	req->data  = xmalloc(msize);
	req->msize = msize;

	return req;
}

static void
request_destroy(request_t* req)
{
	free(req->data);
	free(req->tx);
	free(req->rx);
	free(req);
}

static queue_t*
queue_create(void)
{
	queue_t* q = xmalloc(sizeof(queue_t));

	xmutex_init(&q->mutex);
	xcond_init(&q->cond);
	TAILQ_INIT(&q->head);

	q->size     = 0;
	q->shutdown = 0;

	return q;
}

static int
queue_push(queue_t* q, request_t* req)
{
	xmutex_lock(&q->mutex);
	if(q->shutdown != 0) {
		xmutex_unlock(&q->mutex);
		return -1;
	}

	TAILQ_INSERT_TAIL(&q->head, req, entry);
	q->size++;

	xcond_signal(&q->cond);
	xmutex_unlock(&q->mutex);
	return 0;
}

static request_t*
queue_pop(queue_t* q)
{
	xmutex_lock(&q->mutex);
	while(TAILQ_EMPTY(&q->head) && q->shutdown == 0)
		xcond_wait(&q->cond, &q->mutex);
	if(q->shutdown != 0) {
		xmutex_unlock(&q->mutex);
		return NULL;
	}

	request_t* req = TAILQ_FIRST(&q->head);
	TAILQ_REMOVE(&q->head, req, entry);
	q->size--;
	xmutex_unlock(&q->mutex);
	return req;
}

static void
queue_shutdown(queue_t* q)
{
	xmutex_lock(&q->mutex);
	q->shutdown = 1;
	xcond_broadcast(&q->cond);
	xmutex_unlock(&q->mutex);
}

static void
queue_destroy(queue_t* q)
{
	request_t* req;

	xmutex_lock(&q->mutex);
	while(!TAILQ_EMPTY(&q->head)) {
		req = TAILQ_FIRST(&q->head);
		TAILQ_REMOVE(&q->head, req, entry);
		q->size--;
		free(req);
	}
	xmutex_unlock(&q->mutex);

	xcond_destroy(&q->cond);
	xmutex_destroy(&q->mutex);
	free(q);
}

static int
fidcmp(fid_t* a, fid_t* b)
{
	return a->num - b->num;
}

RB_HEAD(fmap, fid_t);
RB_GENERATE(fmap, fid_t, entry, fidcmp);

struct conn_t {
	pthread_mutex_t state; /* protects internal conn state */
	backend_t* backend;
	int        worker;
	int        fd;
	size_t     msize;

	pthread_mutex_t writer; /* exclusive writer lock */
	queue_t*        queue;

	pthread_mutex_t mutex; /* protects map access */
	struct fmap     head;
};

/*
 * conn_accept accepts a 9P2000.L connection on a socket. The call returns NULL
 * on error and the global variable errno is set to indicate the error.
 */
conn_t*
conn_accept(int socket, backend_t* backend, int worker, size_t msize)
{
	struct sockaddr addr;
	socklen_t       addrlen;
	int             fd;
	conn_t*         conn;

	if((fd = accept(socket, &addr, &addrlen)) == -1)
		return NULL;

	if(msize < HEADERSIZE)
		msize = MAXMESSAGESIZE; // TODO

	struct fmap head = RB_INITIALIZER(&head);
	conn             = xmalloc(sizeof(conn_t));
	xmutex_init(&conn->writer);
	xmutex_init(&conn->mutex);
	xmutex_init(&conn->state);

	conn->backend = backend;
	conn->head    = head;
	conn->worker  = worker;
	conn->fd      = fd;
	conn->msize   = msize;

	warnx("accepted new connection (%p)", conn);
	return conn;
}

static void
fmap_insert(conn_t* conn, fid_t* fid)
{
	xmutex_lock(&conn->mutex);
	RB_INSERT(fmap, &conn->head, fid);
	xmutex_unlock(&conn->mutex);
}

static fid_t*
fmap_get(conn_t* conn, uint32_t num)
{
	fid_t f = {.num = num}, *p = NULL;

	xmutex_lock(&conn->mutex);
	p = RB_FIND(fmap, &conn->head, &f);
	xmutex_unlock(&conn->mutex);

	return p;
}

static void
fmap_destroy(conn_t* conn)
{
	fid_t *f, *nf;

	xmutex_lock(&conn->mutex);
	for(f = RB_MIN(fmap, &conn->head); f != NULL; f = nf) {
		nf = RB_NEXT(fmap, &conn->head, f);
		RB_REMOVE(fmap, &conn->head, f);
		free(f);
	}
	xmutex_unlock(&conn->mutex);
}

static void
fmap_range(conn_t* conn, void (*fn)(fid_t*))
{
	fid_t* p;

	RB_FOREACH(p, fmap, &conn->head)
		fn(p);
}

/*
 * conn_destroy closes a 9P2000.L connection. Rendering it unusable for IO.
 */
void
conn_destroy(conn_t* conn)
{
	fmap_destroy(conn);
	xmutex_destroy(&conn->writer);
	xmutex_destroy(&conn->mutex);
	xmutex_destroy(&conn->state);
	close(conn->fd);
	free(conn);
	warnx("destroyed connection (%p)", conn);
}

static ssize_t
xread(int fd, unsigned char* buf, uint32_t n)
{
	uint32_t       done = 0;
	ssize_t        m;
	unsigned char* p;

	p = buf;
	while(done < n) {
		if((m = read(fd, p + done, n - done)) <= 0) {
			if(errno == EINTR)
				continue;
			if(done == 0)
				return m;
			break;
		}
		done += m;
	}
	return done;
}

static ssize_t
xwrite(int fd, unsigned char* buf, uint32_t n)
{
	uint32_t       done = 0;
	ssize_t        m;
	unsigned char* p;

	p = buf;
	while(done < n) {
		if((m = write(fd, p + done, n - done)) <= 0) {
			if(errno == EINTR)
				continue;
			if(done == 0)
				return m;
			break;
		}
		done += n;
	}
	return done;
}

static ssize_t
conn_send(conn_t* conn, request_t* req)
{
	ssize_t        n, msize = fcall_size(req->rx);
	unsigned char* p;

	if(msize > req->msize) {
		p         = xrealloc(req->data, msize);
		req->data = p;
	}
	req->msize = msize;

	if(fcall_marshal(req->data, msize, req->rx) < 0)
		return -1;

	xmutex_lock(&conn->writer);
	n = xwrite(conn->fd, req->data, msize);
	xmutex_unlock(&conn->writer);
	if(n != msize)
		return -1;
	return n;
}

static void*
conn_worker(void* arg)
{
	conn_t*    conn    = (conn_t*)arg;
	backend_t* b       = conn->backend;
	fcall_t *  tx, *rx;
	int        ret;

	while(err >= 0 && shutdown >= 0) {
		request_t* req = queue_pop(conn->queue);
		if(req == NULL) // queue is shut down
			break;

		warnx("worker received request (%p)", req);
		tx = req->tx;
		rx = req->rx;
		rx->type = tx->type + 1;
		rx->tag  = tx->tag;

		switch(req->tx->type) {
		case TVERSION:
			ret = b->version(conn->msize, &tx->tversion, &rx->rversion);
			warnx("version func done");
			if(ret == 0) {
				xmutex_lock(&conn->state);
				conn->msize = rx->rversion.msize;
				// TODO: destroy pending requests
				xmutex_unlock(&conn->state);
			}
			break;
		case TAUTH:
			ret = b->auth(&tx->tauth, &rx->rauth);
			break;
		case TATTACH: // TODO: is conn versioned check
			ret = b->attach(&tx->tattach, &rx->rattach);
			break;
		default:
			// TODO: deregister from queue and stop conn_serve if no more
			// worker available
			req->rx->rerror.ecode = EINVAL;
			req->rx->type         = RERROR;
			// break;
		}
		if(ret != 0) {
			// TODO: send rerror_t response
		}

		// err = conn_send(conn, req);
		conn_send(conn, req);
		warnx("worker finished request (%p)", req);
		request_destroy(req);
	}
	pthread_exit(NULL);
}

void
conn_serve(conn_t* conn)
{
	pthread_t     w[conn->worker];
	int           i, done;
	unsigned char header[4];
	uint32_t      msize;

	warnx("connection (%p) creating thread pool", conn);
	conn->queue = queue_create();
	for(i    = 0; i < conn->worker; i++)
		w[i] = xthread_create(conn_worker, conn);
	done     = i;

	for(;;) {
		if(xread(conn->fd, header, 4) != 4)
			break;
		msize = header[0] | header[1] << 8 | header[2] << 16 | header[3] << 24;

		request_t* req = request_create(msize);
		memmove(req->data, header, 4);
		req->msize = msize;
		msize -= 4;

		if(xread(conn->fd, req->data + 4, msize) != msize) {
			request_destroy(req);
			break;
		}

		if(fcall_unmarshal(req->data, req->msize, req->tx) < 0) {
			request_destroy(req);
			break;
		}

		if(queue_push(conn->queue, req) < 0) {
			request_destroy(req);
			break;
		}
	}

	queue_shutdown(conn->queue);
	for(i = 0; i < done; i++)
		xthread_join(w[i]);
	queue_destroy(conn->queue);
	warnx("connection (%p) destroyed thread pool", conn);
}
