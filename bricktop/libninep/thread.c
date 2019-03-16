#include "libc.h"

pthread_t
xthread_create(void* (*thread_func)(void*), void* arg)
{
	pthread_t thread;

	if(pthread_create(&thread, NULL, thread_func, arg) != 0)
		abort();
	return thread;
}

void
xthread_join(pthread_t thread)
{
	if(pthread_join(thread, NULL) != 0)
		abort();
}

void
xthread_detach(pthread_t thread)
{
	if(pthread_detach(thread) != 0)
		abort();
}

uint64_t
xthread_id(void)
{
	uint64_t tid;
	if(pthread_threadid_np(pthread_self(), &tid) != 0)
		abort();
	return tid;
}

void
xmutex_init(pthread_mutex_t* mutex)
{
	pthread_mutexattr_t attr;

	pthread_mutexattr_init(&attr);
	pthread_mutexattr_settype(&attr, PTHREAD_MUTEX_NORMAL);
	if(pthread_mutex_init(mutex, &attr) != 0)
		abort();
	pthread_mutexattr_destroy(&attr);
}

void
xmutex_lock(pthread_mutex_t* mutex)
{
	if(pthread_mutex_lock(mutex) != 0)
		abort();
}

void
xmutex_unlock(pthread_mutex_t* mutex)
{
	if(pthread_mutex_unlock(mutex) != 0)
		abort();
}

void
xmutex_destroy(pthread_mutex_t* mutex)
{
	if(pthread_mutex_destroy(mutex) != 0)
		abort();
}

void
xcond_init(pthread_cond_t* cond)
{
	if(pthread_cond_init(cond, NULL) != 0)
		abort();
}

void
xcond_wait(pthread_cond_t* cond, pthread_mutex_t* mutex)
{
	if(pthread_cond_wait(cond, mutex) != 0)
		abort();
}

void
xcond_signal(pthread_cond_t* cond)
{
	if(pthread_cond_signal(cond) != 0)
		abort();
}

void
xcond_broadcast(pthread_cond_t* cond)
{
	if(pthread_cond_broadcast(cond) != 0)
		abort();
}

void
xcond_destroy(pthread_cond_t* cond)
{
	if(pthread_cond_destroy(cond) != 0)
		abort();
}

void*
xcalloc(size_t count, size_t size)
{
	void* v;

	if((v = calloc(count, size)) == NULL)
		err(1, "xcalloc: out of memory");
	return v;
}

void*
xmalloc(size_t size)
{
	void* v;

	if((v = malloc(size)) == NULL)
		err(1, "xmalloc: out of memory");
	memset(v, 0, size);
	return v;
}

void*
xrealloc(void* p, size_t size)
{
	void* v;

	if((v = realloc(p, size)) == NULL)
		err(1, "xrealloc: out of memory");
	return v;
}

char*
xstrdup(const char* s)
{
	char* r;

	if((r = strdup(s)) == NULL)
		err(1, "xstrdup: out of memory");
	return r;
}
