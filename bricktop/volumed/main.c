#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <err.h>
#include "ninep.h"

char* progname;

static void
usage(void)
{
	fprintf(stderr, "Usage: %s [options]\n\n", progname ? progname : "volumed");
	fprintf(stderr, "Options:\n");
	fprintf(stderr, "  -D print each 9P2000.L message to standard output\n");
	fprintf(stderr, "  -h address\n");
	fprintf(stderr, "      service listen address (default \"0.0.0.0\")\n");
	fprintf(stderr, "  -p port\n");
	fprintf(stderr, "      service port number (default \"5640\")\n");
	fprintf(stderr, "");
	exit(2);
}

int
main(int argc, char** argv)
{
	char *    addr = "0.0.0.0", *port = "5640", *chroot = "/tmp";
	char*     network = "tcp";
	int       c, socket, worker = 16;
	uint32_t  msize = 64 * 1024;
	server_t* srv;

	progname = argv[0];
	while((c = getopt(argc, argv, "Dha:u:p:")) != -1)
		switch(c) {
		case 'D':
			break;
		case 'u':
			network = "unix";
			addr    = optarg;
		case 'a':
			network = "tcp";
			addr    = optarg;
			break;
		case 'p':
			port = optarg;
			break;
		case 'h':
			usage();
		default:
			usage();
		}

	if((socket = tcp_announce(addr, port)) == -1)
		err(1, "announce socket %s:%s", addr, port);
	warnx("announce '9P2000.L' service '%s://%s:%s%s' (msize:%d)", network,
	    addr, port, chroot, msize);

	if((srv = server_announce(socket, chroot, msize)) == NULL)
		errx(1, "announce server %s:%s", addr, port);

	server_serve(srv, worker);
	server_destroy(srv);

	return 0;
}
