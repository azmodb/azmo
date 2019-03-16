#include "ninep.h"
#include "libc.h"

int
tcp_announce(const char* addr, const char* port)
{
	struct addrinfo hints, *airoot, *ai;
	int             fd, ret, flags;

	memset(&hints, 0, sizeof(hints));
	hints.ai_family   = PF_UNSPEC;
	hints.ai_socktype = SOCK_STREAM;
	hints.ai_flags    = AI_PASSIVE;
	if(getaddrinfo(addr, port, &hints, &airoot) != 0)
		return -1;

	for(ai = airoot; ai; ai = ai->ai_next) {
		fd = socket(ai->ai_family, ai->ai_socktype, ai->ai_protocol);
		if(fd == -1)
			continue;

		flags = 1;

		ret = setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &flags, sizeof(flags));
		if(ret != 0) {
			close(fd);
			continue;
		}
		ret = setsockopt(fd, SOL_SOCKET, SO_KEEPALIVE, &flags, sizeof(flags));
		if(ret != 0) {
			close(fd);
			continue;
		}

		if(bind(fd, ai->ai_addr, ai->ai_addrlen) == -1) {
			close(fd);
			return -1;
		}
		if(listen(fd, 1024) == -1) {
			close(fd);
			return -1;
		}
		break;
	}

	freeaddrinfo(airoot);
	if(ai == NULL)
		return -1;
	return fd;
}

int
unix_announce(const char* path)
{
	return -1;
}

struct server_t {
	backend_t* backend;
	int        socket;
	size_t     msize;
};

server_t*
server_announce(int socket, const char* aname, size_t msize)
{
	char*     root;
	server_t* srv;

	if((root = realpath(aname, NULL)) == NULL)
		err(1, "cannot resolve attach name \"%s\"", root);

	srv          = xmalloc(sizeof(server_t));
	srv->backend = backend_create(root);
	srv->socket  = socket;
	srv->msize   = msize;
	return srv;
}

void
server_destroy(server_t* srv)
{
	close(srv->socket);
	free(srv);
}

static void*
server_worker(void* arg)
{
	uint64_t id = xthread_id();

	warnx("starting service connection thread: %lld", id);
	conn_t* conn = (conn_t*)arg;
	conn_serve(conn);
	conn_destroy(conn);
	warnx("closing service connection thread: %lld", id);
	pthread_exit(NULL);
}

void
server_serve(server_t* srv, int worker)
{
	int socket = srv->socket;

	for(;;) {
		conn_t* conn;
		conn = conn_accept(socket, srv->backend, worker, srv->msize);
		if(conn == NULL)
			break;

		pthread_t worker = xthread_create(server_worker, conn);
		xthread_detach(worker);
	}
	warnx("shut down server connection (fd:%d)", srv->socket);
}
