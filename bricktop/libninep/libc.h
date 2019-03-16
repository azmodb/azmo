#ifndef __INTERNAL_LIBC_H
#define __INTERNAL_LIBC_H

#include <sys/socket.h>
#include <sys/param.h>
#include <pthread.h>
#include <stdlib.h>
#include <unistd.h>
#include <netdb.h>
#include <errno.h>
#include <err.h>
#include "fcall.h"

#define VERSION "9P2000.L"

/* 9P2000.L magic numbers */
#define NOTAG (uint16_t)(~0)
#define NOFID (uint32_t)(~0)
#define NONUNAME (uint32_t)(~0)
#define MAXWELEM 16

/* backend.c */
typedef struct backend_t backend_t;
struct backend_t {
	/*
	 * If root is a directory client attaches are not allowed to go outside the
	 * directory represented by root. If root is a block device it is served
	 * directly.
	 */
	const char* root;

	int (*version)(uint32_t, tversion_t*, rversion_t*);
	int (*auth)(tauth_t*, rauth_t*);
	int (*attach)(tattach_t*, rattach_t*);

	int (*flush)(tflush_t*, rflush_t*);
	int (*walk)(twalk_t*, rwalk_t*);
	int (*clunk)(tclunk_t*, rclunk_t*);
	int (*remove)(tremove_t*, rremove_t*);
	int (*read)(tread_t*, rread_t*);
	int (*write)(twrite_t*, rwrite_t*);
	int (*statfs)(tstatfs_t*, rstatfs_t*);
	int (*open)(topen_t*, ropen_t*);
	int (*create)(tcreate_t*, rcreate_t*);
	int (*symlink)(tsymlink_t*, rsymlink_t*);
	int (*mknod)(tmknod_t*, rmknod_t*);
	int (*rename)(trename_t*, rrename_t*);
	int (*readlink)(treadlink_t*, rreadlink_t*);
	int (*getattr)(tgetattr_t*, rgetattr_t*);
	int (*setattr)(tsetattr_t*, rsetattr_t*);
	int (*xattrwalk)(txattrwalk_t*, rxattrwalk_t*);
	int (*xattrcreate)(txattrcreate_t*, rxattrcreate_t*);
	int (*readdir)(treaddir_t*, rreaddir_t*);
	int (*fync)(tfync_t*, rfsync_t*);
	int (*lock)(tlock_t*, rlock_t*);
	int (*getlock)(tgetlock_t*, rgetlock_t*);
	int (*link)(tlink_t*, rlink_t*);
	int (*mkdir)(tmkdir_t*, rmkdir_t*);
	int (*renameat)(trenameat_t*, rrenameat_t*);
	int (*unlinkat)(tunlinkat_t*, runlinkat_t*);
};

backend_t* backend_create(const char*);
void       backend_destroy(backend_t*);

/* conn.c */
typedef struct conn_t conn_t;

conn_t* conn_accept(int, backend_t*, int, size_t);
void    conn_serve(conn_t*);
void    conn_destroy(conn_t*);

/* thread.c */
pthread_t xthread_create(void* (*thread_func)(void*), void* arg);
void     xthread_join(pthread_t);
uint64_t xthread_id(void);
void     xthread_detach(pthread_t);

void xmutex_init(pthread_mutex_t*);
void xmutex_lock(pthread_mutex_t*);
void xmutex_unlock(pthread_mutex_t*);
void xmutex_destroy(pthread_mutex_t*);

void xcond_init(pthread_cond_t*);
void xcond_wait(pthread_cond_t*, pthread_mutex_t*);
void xcond_signal(pthread_cond_t*);
void xcond_broadcast(pthread_cond_t*);
void xcond_destroy(pthread_cond_t*);

void* xmalloc(size_t);
void* xrealloc(void*, size_t);
void* xcalloc(size_t, size_t);
char* xstrdup(const char*);

#endif /* __INTERNAL_LIBC_H */
