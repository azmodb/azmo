/*
 * Copyright (C) 2005 by Latchesar Ionkov <lucho@ionkov.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice (including the next
 * paragraph) shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
 * LATCHESAR IONKOV AND/OR ITS SUPPLIERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */
#include <sys/types.h>
#include <pthread.h>
#include <stdint.h>

typedef uint8_t   u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;

typedef struct Npstr Npstr;
typedef struct Npqid Npqid;
typedef struct Npstat Npstat;
typedef struct Npwstat Npwstat;
typedef struct Npfcall Npfcall;
typedef struct Npfid Npfid;
typedef struct Npfidpool Npfidpool;
typedef struct Nptrans Nptrans;
typedef struct Npconn Npconn;
typedef struct Npreq Npreq;
typedef struct Npwthread Npwthread;
typedef struct Npauth Npauth;
typedef struct Npsrv Npsrv;
typedef struct Npuser Npuser;
typedef struct Npgroup Npgroup;
typedef struct Npuserpool Npuserpool;

/* message types */
enum {
	TFIRST		= 100,
	TVERSION	= 100,
	RVERSION,
	TAUTH		= 102,
	RAUTH,
	TATTACH		= 104,
	RATTACH,
	TERROR		= 106,
	RERROR,
	TFLUSH		= 108,
	RFLUSH,
	TWALK		= 110,
	RWALK,
	TOPEN		= 112,
	ROPEN,
	TCREATE		= 114,
	RCREATE,
	TREAD		= 116,
	RREAD,
	TWRITE		= 118,
	RWRITE,
	TCLUNK		= 120,
	RCLUNK,
	TREMOVE		= 122,
	RREMOVE,
	TSTAT		= 124,
	RSTAT,
	TWSTAT		= 126,
	RWSTAT,
	RLAST
};

/* modes */
enum {
	OREAD		= 0x00,
	OWRITE		= 0x01,
	ORDWR		= 0x02,
	OEXEC		= 0x03,
	OEXCL		= 0x04,
	OTRUNC		= 0x10,
	OREXEC		= 0x20,
	ORCLOSE		= 0x40,
	OAPPEND		= 0x80,

	OUSPECIAL	= 0x100,	/* internal use */
};

/* permissions */
enum {
	DMDIR		= 0x80000000,
	DMAPPEND	= 0x40000000,
	DMEXCL		= 0x20000000,
	DMMOUNT		= 0x10000000,
	DMAUTH		= 0x08000000,
	DMTMP		= 0x04000000,
	DMSYMLINK	= 0x02000000,
	DMLINK		= 0x01000000,

	/* 9P2000.u extensions */
	DMDEVICE	= 0x00800000,
	DMNAMEDPIPE	= 0x00200000,
	DMSOCKET	= 0x00100000,
	DMSETUID	= 0x00080000,
	DMSETGID	= 0x00040000,
};

/* qid.types */
enum {
	QTDIR		= 0x80,
	QTAPPEND	= 0x40,
	QTEXCL		= 0x20,
	QTMOUNT		= 0x10,
	QTAUTH		= 0x08,
	QTTMP		= 0x04,
	QTSYMLINK	= 0x02,
	QTLINK		= 0x01,
	QTFILE		= 0x00,
};

#define NOTAG		(u16)(~0)
#define NOFID		(u32)(~0)
#define MAXWELEM	16
#define IOHDRSZ		24
#define FID_HTABLE_SIZE 64

struct Npstr {
	u16		len;
	char*		str;
};

struct Npqid {
	u8		type;
	u32		version;
	u64		path;
};

struct Npstat {
	u16 		size;
	u16 		type;
	u32 		dev;
	Npqid		qid;
	u32 		mode;
	u32 		atime;
	u32 		mtime;
	u64 		length;
	Npstr		name;
	Npstr		uid;
	Npstr		gid;
	Npstr		muid;

	/* 9P2000.u extensions */
	Npstr		extension;
	u32 		n_uid;
	u32 		n_gid;
	u32 		n_muid;
};
 
/* file metadata (stat) structure used to create TWSTAT message
   It is similar to Npstat, but the strings don't point to 
   the same memory block and should be freed separately
*/
struct Npwstat {
	u16 		size;
	u16 		type;
	u32 		dev;
	Npqid		qid;
	u32 		mode;
	u32 		atime;
	u32 		mtime;
	u64 		length;
	char*		name;
	char*		uid;
	char*		gid;
	char*		muid;
	char*		extension;	/* 9p2000.u extensions */
	u32 		n_uid;		/* 9p2000.u extensions */
	u32 		n_gid;		/* 9p2000.u extensions */
	u32 		n_muid;		/* 9p2000.u extensions */
};

struct Npfcall {
	u32		size;
	u8		type;
	u16		tag;
	u8*		pkt;

	u32		fid;
	u32		msize;			/* TVERSION, RVERSION */
	Npstr		version;		/* TVERSION, RVERSION */
	u32		afid;			/* TAUTH, TATTACH */
	Npstr		uname;			/* TAUTH, TATTACH */
	Npstr		aname;			/* TAUTH, TATTACH */
	Npqid		qid;			/* RAUTH, RATTACH, ROPEN, RCREATE */
	Npstr		ename;			/* RERROR */
	u16		oldtag;			/* TFLUSH */
	u32		newfid;			/* TWALK */
	u16		nwname;			/* TWALK */
	Npstr		wnames[MAXWELEM];	/* TWALK */
	u16		nwqid;			/* RWALK */
	Npqid		wqids[MAXWELEM];	/* RWALK */
	u8		mode;			/* TOPEN, TCREATE */
	u32		iounit;			/* ROPEN, RCREATE */
	Npstr		name;			/* TCREATE */
	u32		perm;			/* TCREATE */
	u64		offset;			/* TREAD, TWRITE */
	u32		count;			/* TREAD, RREAD, TWRITE, RWRITE */
	u8*		data;			/* RREAD, TWRITE */
	Npstat		stat;			/* RSTAT, TWSTAT */

	/* 9P2000.u extensions */
	u32		ecode;			/* RERROR */
	Npstr		extension;		/* TCREATE */
	u32		n_uname;

	Npfcall*	next;
};

struct Npfid {
	pthread_mutex_t	lock;
	Npconn*		conn;
	u32		fid;
	int		refcount;
	u16		omode;
	u8		type;
	u32		diroffset;
	Npuser*		user;
	void*		aux;

	Npfid*		next;	/* list of fids within a bucket */
	Npfid*		prev;
};

struct Nptrans {
	void*		aux;
	int		(*read)(u8 *, u32, void *);
	int		(*write)(u8 *, u32, void *);
	void		(*destroy)(void *);
};

struct Npfidpool {
	pthread_mutex_t	lock;
	int		size;
	Npfid**		htable;
};

struct Npconn {
	pthread_mutex_t	lock;
	pthread_mutex_t	wlock;
	int		refcount;

	int		resetting;
	pthread_cond_t	resetcond;
	pthread_cond_t	resetdonecond;

	u32		msize;
	int		dotu;
	int		shutdown;
	Npsrv*		srv;
	Nptrans*	trans;
	Npfidpool*	fidpool;
	int		freercnum;
	Npfcall*	freerclist;
	void*		aux;
	pthread_t	rthread;

	Npconn*		next;	/* list of connections within a server */
};

struct Npreq {
	pthread_mutex_t	lock;
	int		refcount;
	Npconn*		conn;
	u16		tag;
	Npfcall*	tcall;
	Npfcall*	rcall;
	int		cancelled;
	int		responded;
	Npreq*		flushreq;
	Npfid*		fid;

	Npreq*		next;	/* list of all outstanding requests */
	Npreq*		prev;	/* used for requests that are worked on */
	Npwthread*	wthread;/* for requests that are worked on */
};

struct Npwthread {
	Npsrv*		srv;
	int		shutdown;
	pthread_t	thread;

	Npwthread	*next;
};

struct Npauth {
	int	(*startauth)(Npfid *afid, char *aname, Npqid *aqid);
	int	(*checkauth)(Npfid *fid, Npfid *afid, char *aname);
	int	(*read)(Npfid *fid, u64 offset, u32 count, u8 *data);
	int	(*write)(Npfid *fid, u64 offset, u32 count, u8 *data);
	int	(*clunk)(Npfid *fid);
};	

struct Npsrv {
	u32		msize;
	int		dotu;		/* 9P2000.u support flag */
	void*		srvaux;
	void*		treeaux;
	int		debuglevel;
	Npauth*		auth;
	Npuserpool*	upool;

	int			(*start)(Npsrv *);
	void		(*shutdown)(Npsrv *);
	void		(*destroy)(Npsrv *);
	void		(*connopen)(Npconn *);
	void		(*connclose)(Npconn *);
	void		(*fiddestroy)(Npfid *);

	Npfcall*	(*version)(Npconn *conn, u32 msize, Npstr *version);
//	Npfcall*	(*auth)(Npfid *afid, Npstr *uname, Npstr *aname);
	Npfcall*	(*attach)(Npfid *fid, Npfid *afid, Npstr *uname, 
				Npstr *aname);
	void		(*flush)(Npreq *req);
	int		(*clone)(Npfid *fid, Npfid *newfid);
	int		(*walk)(Npfid *fid, Npstr *wname, Npqid *wqid);
	Npfcall*	(*open)(Npfid *fid, u8 mode);
	Npfcall*	(*create)(Npfid *fid, Npstr* name, u32 perm, u8 mode, 
				Npstr* extension);
	Npfcall*	(*read)(Npfid *fid, u64 offset, u32 count, Npreq *req);
	Npfcall*	(*write)(Npfid *fid, u64 offset, u32 count, u8 *data, 
				Npreq *req);
	Npfcall*	(*clunk)(Npfid *fid);
	Npfcall*	(*remove)(Npfid *fid);
	Npfcall*	(*stat)(Npfid *fid);
	Npfcall*	(*wstat)(Npfid *fid, Npstat *stat);

	/* implementation specific */
	pthread_mutex_t	lock;
	pthread_cond_t	reqcond;
	int		shuttingdown;
	Npconn*		conns;
	int		nwthread;
	Npwthread*	wthreads;
	Npreq*		reqs_first;
	Npreq*		reqs_last;
	Npreq*		workreqs;
};

struct Npuser {
	pthread_mutex_t	lock;
	int			refcount;
	Npuserpool*	upool;
	char*		uname;
	uid_t		uid;
	Npgroup*	dfltgroup;
	int			ngroups;	
	Npgroup**	groups;
	void*		aux;

	Npuser*		next;
};

struct Npgroup {
	pthread_mutex_t	lock;
	int			refcount;
	Npuserpool* upool;
	char*		gname;
	gid_t		gid;
	void*		aux;

	Npgroup*	next;
};

struct Npuserpool {
	void*		aux;
	Npuser*		(*uname2user)(Npuserpool *, char *uname);
	Npuser*		(*uid2user)(Npuserpool *, u32 uid);
	Npgroup*	(*gname2group)(Npuserpool *, char *gname);
	Npgroup*	(*gid2group)(Npuserpool *, u32 gid);
	int		(*ismember)(Npuserpool *, Npuser *u, Npgroup *g);
	void		(*udestroy)(Npuserpool *, Npuser *u);
	void		(*gdestroy)(Npuserpool *, Npgroup *g);
};

extern char *Eunknownfid;
extern char *Ennomem;
extern char *Enoauth;
extern char *Enotimpl;
extern char *Einuse;
extern char *Ebadusefid;
extern char *Enotdir;
extern char *Etoomanywnames;
extern char *Eperm;
extern char *Etoolarge;
extern char *Ebadoffset;
extern char *Edirchange;
extern char *Enotfound;
extern char *Eopen;
extern char *Eexist;
extern char *Enotempty;
extern char *Eunknownuser;
extern Npuserpool *np_default_users;

Npsrv *np_srv_create(int nwthread);
void np_srv_remove_conn(Npsrv *, Npconn *);
int np_srv_start(Npsrv *);
void np_srv_shutdown(Npsrv *, int wait);
int np_srv_add_conn(Npsrv *, Npconn *);

Npconn *np_conn_create(Npsrv *, Nptrans *);
void np_conn_incref(Npconn *);
void np_conn_decref(Npconn *);
void np_conn_reset(Npconn *, u32, int);
void np_conn_shutdown(Npconn *);
void np_conn_respond(Npreq *req);
void np_respond(Npreq *, Npfcall *);

Npfidpool *np_fidpool_create(void);
void np_fidpool_destroy(Npfidpool *);
Npfid *np_fid_find(Npconn *, u32);
Npfid *np_fid_create(Npconn *, u32, void *);
void np_fid_destroy(Npfid *);
void np_fid_incref(Npfid *);
void np_fid_decref(Npfid *);

Nptrans *np_trans_create(void *aux, int (*read)(u8 *, u32, void *),
	int (*write)(u8 *, u32, void *), void (*destroy)(void *));
void np_trans_destroy(Nptrans *);
int np_trans_read(Nptrans *, u8 *, u32);
int np_trans_write(Nptrans *, u8 *, u32);

int np_deserialize(Npfcall*, u8*, int dotu);
int np_serialize_stat(Npwstat *wstat, u8* buf, int buflen, int dotu);
int np_deserialize_stat(Npstat *stat, u8* buf, int buflen, int dotu);

void np_strzero(Npstr *str);
char *np_strdup(Npstr *str);
int np_strcmp(Npstr *str, char *cs);
int np_strncmp(Npstr *str, char *cs, int len);

void np_set_tag(Npfcall *, u16);
Npfcall *np_create_tversion(u32 msize, char *version);
Npfcall *np_create_rversion(u32 msize, char *version);
Npfcall *np_create_tauth(u32 fid, char *uname, char *aname, u32 n_uname, int dotu);
Npfcall *np_create_rauth(Npqid *aqid);
Npfcall *np_create_rerror(char *ename, int ecode, int dotu);
Npfcall *np_create_rerror1(Npstr *ename, int ecode, int dotu);
Npfcall *np_create_tflush(u16 oldtag);
Npfcall *np_create_rflush(void);
Npfcall *np_create_tattach(u32 fid, u32 afid, char *uname, char *aname, u32 n_uname, int dotu);
Npfcall *np_create_rattach(Npqid *qid);
Npfcall *np_create_twalk(u32 fid, u32 newfid, u16 nwname, char **wnames);
Npfcall *np_create_rwalk(int nwqid, Npqid *wqids);
Npfcall *np_create_topen(u32 fid, u8 mode);
Npfcall *np_create_ropen(Npqid *qid, u32 iounit);
Npfcall *np_create_tcreate(u32 fid, char *name, u32 perm, u8 mode);
Npfcall *np_create_rcreate(Npqid *qid, u32 iounit);
Npfcall *np_create_tread(u32 fid, u64 offset, u32 count);
Npfcall *np_create_rread(u32 count, u8* data);
Npfcall *np_create_twrite(u32 fid, u64 offset, u32 count, u8 *data);
Npfcall *np_create_rwrite(u32 count);
Npfcall *np_create_tclunk(u32 fid);
Npfcall *np_create_rclunk(void);
Npfcall *np_create_tremove(u32 fid);
Npfcall *np_create_rremove(void);
Npfcall *np_create_tstat(u32 fid);
Npfcall *np_create_rstat(Npwstat *stat, int dotu);
Npfcall *np_create_twstat(u32 fid, Npwstat *wstat, int dotu);
Npfcall *np_create_rwstat(void);
Npfcall * np_alloc_rread(u32);
void np_set_rread_count(Npfcall *, u32);

int np_printstat(FILE *f, Npstat *st, int dotu);
int np_printfcall(FILE *f, Npfcall *fc, int dotu);

void np_user_incref(Npuser *);
void np_user_decref(Npuser *);
void np_group_incref(Npgroup *);
void np_group_decref(Npgroup *);
int np_change_user(Npuser *u);
Npuser* np_current_user(void);

Npuserpool *np_priv_userpool_create();
Npuser *np_priv_user_add(Npuserpool *up, char *uname, u32 uid, void *aux);
void np_priv_user_del(Npuser *u);
int np_priv_user_setdfltgroup(Npuser *u, Npgroup *g);
Npgroup *np_priv_group_add(Npuserpool *up, char *gname, u32 gid);
void np_priv_group_del(Npgroup *g);
int np_priv_group_adduser(Npgroup *g, Npuser *u);
int np_priv_group_deluser(Npgroup *g, Npuser *u);

Nptrans *np_fdtrans_create(int, int);
Npsrv *np_socksrv_create_tcp(int, int*);
Npsrv *np_pipesrv_create(int nwthreads);
int np_pipesrv_mount(Npsrv *srv, char *mntpt, char *user, int mntflags, char *opts);

Npsrv *np_rdmasrv_create(int nwthreads, int *port);

void np_werror(char *ename, int ecode, ...);
void np_rerror(char **ename, int *ecode);
int np_haserror(void);
void np_uerror(int ecode);
void np_suerror(char *s, int ecode);

void npfile_fiddestroy(Npfid *fid);
Npfcall *npfile_attach(Npfid *fid, Npfid *afid, Npstr *uname, Npstr *aname);
int npfile_clone(Npfid *fid, Npfid *newfid);
int npfile_walk(Npfid *fid, Npstr *wname, Npqid *wqid);
Npfcall *npfile_open(Npfid *fid, u8 mode);
Npfcall *npfile_create(Npfid *fid, Npstr* name, u32 perm, u8 mode, Npstr* extension);
Npfcall *npfile_read(Npfid *fid, u64 offset, u32 count, Npreq *req);
Npfcall *npfile_write(Npfid *fid, u64 offset, u32 count, u8 *data, Npreq *req);
Npfcall *npfile_clunk(Npfid *fid);
Npfcall *npfile_remove(Npfid *fid);
Npfcall *npfile_stat(Npfid *fid);
Npfcall *npfile_wstat(Npfid *fid, Npstat *stat);

void *np_malloc(int);
