#include "libc.h"

static int
xversion(uint32_t msize, tversion_t* tx, rversion_t* rx)
{
	if(strcmp(tx->version, VERSION) != 0)
		return EINVAL;

	rx->version = VERSION;
	if(tx->msize > msize)
		rx->msize = msize;
	if(tx->msize < msize)
		rx->msize = msize;

	return 0;
}

static int
xauth(tauth_t* tx, rauth_t* rx)
{
	return EINVAL; // auth not implemented
}

static int
xattach(tattach_t* tx, rattach_t* rx)
{
	if(tx->afid != NOFID) // auth not implemented
		return EINVAL;

	// TODO: ...
	return 0;
}

static int
xflush(tflush_t* tx, rflush_t* rx)
{
	return ENOTSUP;
}

static int
xwalk(twalk_t* tx, rwalk_t* rx)
{
	return ENOTSUP;
}

static int
xclunk(tclunk_t* tx, rclunk_t* rx)
{
	return ENOTSUP;
}

static int
xremove(tremove_t* tx, rremove_t* rx)
{
	return ENOTSUP;
}

static int
xread(tread_t* tx, rread_t* rx)
{
	return ENOTSUP;
}

static int
xwrite(twrite_t* tx, rwrite_t* rx)
{
	return ENOTSUP;
}

static int
xstatfs(tstatfs_t* tx, rstatfs_t* rx)
{
	return ENOTSUP;
}

static int
xopen(topen_t* tx, ropen_t* rx)
{
	return ENOTSUP;
}

static int
xcreate(tcreate_t* tx, rcreate_t* rx)
{
	return ENOTSUP;
}

static int
xsymlink(tsymlink_t* tx, rsymlink_t* rx)
{
	return ENOTSUP;
}

static int
xmknod(tmknod_t* tx, rmknod_t* rx)
{
	return ENOTSUP;
}

static int
xrename(trename_t* tx, rrename_t* rx)
{
	return ENOTSUP;
}

static int
xreadlink(treadlink_t* tx, rreadlink_t* rx)
{
	return ENOTSUP;
}

static int
xgetattr(tgetattr_t* tx, rgetattr_t* rx)
{
	return ENOTSUP;
}

static int
xsetattr(tsetattr_t* tx, rsetattr_t* rx)
{
	return ENOTSUP;
}

static int
xxattrwalk(txattrwalk_t* tx, rxattrwalk_t* rx)
{
	return ENOTSUP;
}

static int
xxattrcreate(txattrcreate_t* tx, rxattrcreate_t* rx)
{
	return ENOTSUP;
}

static int
xreaddir(treaddir_t* tx, rreaddir_t* rx)
{
	return ENOTSUP;
}

static int
xfync(tfync_t* tx, rfsync_t* rx)
{
	return ENOTSUP;
}

static int
xlock(tlock_t* tx, rlock_t* rx)
{
	return ENOTSUP;
}

static int
xgetlock(tgetlock_t* tx, rgetlock_t* rx)
{
	return ENOTSUP;
}

static int
xlink(tlink_t* tx, rlink_t* rx)
{
	return ENOTSUP;
}

static int
xmkdir(tmkdir_t* tx, rmkdir_t* rx)
{
	return ENOTSUP;
}

static int
xrenameat(trenameat_t* tx, rrenameat_t* rx)
{
	return ENOTSUP;
}

static int
xunlinkat(tunlinkat_t* tx, runlinkat_t* rx)
{
	return ENOTSUP;
}

backend_t*
backend_create(const char* root)
{
	backend_t* b = xmalloc(sizeof(backend_t));

	b->root    = root;
	b->version = xversion;
	b->auth    = xauth;
	b->attach  = xattach;

	b->flush       = xflush;
	b->walk        = xwalk;
	b->clunk       = xclunk;
	b->remove      = xremove;
	b->read        = xread;
	b->write       = xwrite;
	b->statfs      = xstatfs;
	b->open        = xopen;
	b->create      = xcreate;
	b->symlink     = xsymlink;
	b->mknod       = xmknod;
	b->rename      = xrename;
	b->readlink    = xreadlink;
	b->getattr     = xgetattr;
	b->setattr     = xsetattr;
	b->xattrwalk   = xxattrwalk;
	b->xattrcreate = xxattrcreate;
	b->readdir     = xreaddir;
	b->fync        = xfync;
	b->lock        = xlock;
	b->getlock     = xgetlock;
	b->link        = xlink;
	b->mkdir       = xmkdir;
	b->renameat    = xrenameat;
	b->unlinkat    = xunlinkat;

	return b;
}

void
backend_destroy(backend_t* b)
{
	free((char*)b->root); /* allocated by realpath() */
	free(b);
}
