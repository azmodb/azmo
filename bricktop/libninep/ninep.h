#ifndef __NINEP_H
#define __NINEP_H

#include <sys/types.h>

typedef struct server_t server_t;

server_t* server_announce(int, const char*, size_t);
void      server_serve(server_t*, int);
void      server_destroy(server_t*);

int tcp_announce(const char*, const char*);
int unix_announce(const char*);

#endif /* __NINEP_H */
