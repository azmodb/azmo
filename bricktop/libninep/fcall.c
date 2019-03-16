/*
 * GENERATED BY 'go run ../internal/generator -c -root .'; DO NOT EDIT!
 */

#include "fcall.h"

ssize_t
fcall_unmarshal(unsigned char* data, size_t msize, fcall_t* f)
{
	uint32_t       size = 0;
	unsigned char *p, *ep;

	p  = data;
	ep = p + msize;
	if(p + 7 > ep)
		return -1;

	p = guint32(p, ep, &size);
	if(size < 7)
		return 0;
	p = guint8(p, ep, &f->type);
	p = guint16(p, ep, &f->tag);

	switch(f->type) {
	case TVERSION:
		p = guint32(p, ep, &f->tversion.msize);
		p = gstring(p, ep, &f->tversion.version);
		break;
	case RVERSION:
		p = guint32(p, ep, &f->rversion.msize);
		p = gstring(p, ep, &f->rversion.version);
		break;
	case TFLUSH:
		p = guint16(p, ep, &f->tflush.oldtag);
		break;
	case RFLUSH:
		// nothing
		break;
	case TWALK:
		p = guint32(p, ep, &f->twalk.fid);
		p = guint32(p, ep, &f->twalk.newfid);
		// p = gstrings(p, ep, &f->twalk.names);
		break;
	case RWALK:
		// nothing
		break;
	case TCLUNK:
		p = guint32(p, ep, &f->tclunk.fid);
		break;
	case RCLUNK:
		// nothing
		break;
	case TREMOVE:
		p = guint32(p, ep, &f->tremove.fid);
		break;
	case RREMOVE:
		// nothing
		break;
	case TAUTH:
		p = guint32(p, ep, &f->tauth.afid);
		p = gstring(p, ep, &f->tauth.uname);
		p = gstring(p, ep, &f->tauth.aname);
		p = guint32(p, ep, &f->tauth.uid);
		break;
	case RAUTH:
		p = guint8(p, ep, &f->rauth.qid.type);
		p = guint32(p, ep, &f->rauth.qid.version);
		p = guint64(p, ep, &f->rauth.qid.path);
		break;
	case TATTACH:
		p = guint32(p, ep, &f->tattach.fid);
		p = guint32(p, ep, &f->tattach.afid);
		p = gstring(p, ep, &f->tattach.uname);
		p = gstring(p, ep, &f->tattach.aname);
		p = guint32(p, ep, &f->tattach.uid);
		break;
	case RATTACH:
		p = guint8(p, ep, &f->rattach.qid.type);
		p = guint32(p, ep, &f->rattach.qid.version);
		p = guint64(p, ep, &f->rattach.qid.path);
		break;
	case TREAD:
		p = guint32(p, ep, &f->tread.fid);
		p = guint64(p, ep, &f->tread.offset);
		p = guint32(p, ep, &f->tread.count);
		break;
	case RREAD:
		p = gdata(p, ep, &f->rread.data);
		break;
	case TWRITE:
		p = guint32(p, ep, &f->twrite.fid);
		p = guint64(p, ep, &f->twrite.offset);
		p = gdata(p, ep, &f->twrite.data);
		break;
	case RWRITE:
		p = guint32(p, ep, &f->rwrite.count);
		break;
	case TSTATFS:
		p = guint32(p, ep, &f->tstatfs.fid);
		break;
	case RSTATFS:
		p = guint32(p, ep, &f->rstatfs.type);
		p = guint32(p, ep, &f->rstatfs.bsize);
		p = guint64(p, ep, &f->rstatfs.blocks);
		p = guint64(p, ep, &f->rstatfs.bfree);
		p = guint64(p, ep, &f->rstatfs.bavail);
		p = guint64(p, ep, &f->rstatfs.files);
		p = guint64(p, ep, &f->rstatfs.ffree);
		p = guint64(p, ep, &f->rstatfs.fsid);
		p = guint32(p, ep, &f->rstatfs.namelen);
		break;
	case TOPEN:
		p = guint32(p, ep, &f->topen.fid);
		p = guint32(p, ep, &f->topen.flags);
		break;
	case ROPEN:
		p = guint8(p, ep, &f->ropen.qid.type);
		p = guint32(p, ep, &f->ropen.qid.version);
		p = guint64(p, ep, &f->ropen.qid.path);
		p = guint32(p, ep, &f->ropen.iounit);
		break;
	case TCREATE:
		p = guint32(p, ep, &f->tcreate.fid);
		p = gstring(p, ep, &f->tcreate.name);
		p = guint32(p, ep, &f->tcreate.flags);
		p = guint32(p, ep, &f->tcreate.mode);
		p = guint32(p, ep, &f->tcreate.gid);
		break;
	case RCREATE:
		p = guint8(p, ep, &f->rcreate.qid.type);
		p = guint32(p, ep, &f->rcreate.qid.version);
		p = guint64(p, ep, &f->rcreate.qid.path);
		p = guint32(p, ep, &f->rcreate.iounit);
		break;
	case TSYMLINK:
		p = guint32(p, ep, &f->tsymlink.fid);
		p = gstring(p, ep, &f->tsymlink.name);
		p = gstring(p, ep, &f->tsymlink.symtgt);
		p = guint32(p, ep, &f->tsymlink.gid);
		break;
	case RSYMLINK:
		p = guint8(p, ep, &f->rsymlink.qid.type);
		p = guint32(p, ep, &f->rsymlink.qid.version);
		p = guint64(p, ep, &f->rsymlink.qid.path);
		break;
	case TMKNOD:
		p = guint32(p, ep, &f->tmknod.dfid);
		p = gstring(p, ep, &f->tmknod.name);
		p = guint32(p, ep, &f->tmknod.mode);
		p = guint32(p, ep, &f->tmknod.major);
		p = guint32(p, ep, &f->tmknod.minor);
		p = guint32(p, ep, &f->tmknod.gid);
		break;
	case RMKNOD:
		p = guint8(p, ep, &f->rmknod.qid.type);
		p = guint32(p, ep, &f->rmknod.qid.version);
		p = guint64(p, ep, &f->rmknod.qid.path);
		break;
	case TRENAME:
		p = guint32(p, ep, &f->trename.fid);
		p = guint32(p, ep, &f->trename.dfid);
		p = gstring(p, ep, &f->trename.name);
		break;
	case RRENAME:
		// nothing
		break;
	case TREADLINK:
		p = guint32(p, ep, &f->treadlink.fid);
		break;
	case RREADLINK:
		p = gstring(p, ep, &f->rreadlink.target);
		break;
	case TGETATTR:
		p = guint32(p, ep, &f->tgetattr.fid);
		p = guint64(p, ep, &f->tgetattr.request_mask);
		break;
	case RGETATTR:
		p = guint64(p, ep, &f->rgetattr.valid);
		p = guint8(p, ep, &f->rgetattr.qid.type);
		p = guint32(p, ep, &f->rgetattr.qid.version);
		p = guint64(p, ep, &f->rgetattr.qid.path);
		p = guint32(p, ep, &f->rgetattr.mode);
		p = guint32(p, ep, &f->rgetattr.uid);
		p = guint32(p, ep, &f->rgetattr.gid);
		p = guint64(p, ep, &f->rgetattr.nlink);
		p = guint64(p, ep, &f->rgetattr.rdev);
		p = guint64(p, ep, &f->rgetattr.size);
		p = guint64(p, ep, &f->rgetattr.blksize);
		p = guint64(p, ep, &f->rgetattr.blocks);
		p = guint64(p, ep, &f->rgetattr.atime_sec);
		p = guint64(p, ep, &f->rgetattr.atime_nsec);
		p = guint64(p, ep, &f->rgetattr.mtime_sec);
		p = guint64(p, ep, &f->rgetattr.mtime_nsec);
		p = guint64(p, ep, &f->rgetattr.ctime_sec);
		p = guint64(p, ep, &f->rgetattr.ctime_nsec);
		p = guint64(p, ep, &f->rgetattr.btime_sec);
		p = guint64(p, ep, &f->rgetattr.btime_nsec);
		p = guint64(p, ep, &f->rgetattr.gen);
		p = guint64(p, ep, &f->rgetattr.data_version);
		break;
	case TSETATTR:
		p = guint32(p, ep, &f->tsetattr.fid);
		p = guint32(p, ep, &f->tsetattr.valid);
		p = guint32(p, ep, &f->tsetattr.mode);
		p = guint32(p, ep, &f->tsetattr.uid);
		p = guint32(p, ep, &f->tsetattr.gid);
		p = guint64(p, ep, &f->tsetattr.size);
		p = guint64(p, ep, &f->tsetattr.atime_sec);
		p = guint64(p, ep, &f->tsetattr.atime_nsec);
		p = guint64(p, ep, &f->tsetattr.mtime_sec);
		p = guint64(p, ep, &f->tsetattr.mtime_nsec);
		break;
	case RSETATTR:
		// nothing
		break;
	case TXATTRWALK:
		p = guint32(p, ep, &f->txattrwalk.fid);
		p = guint32(p, ep, &f->txattrwalk.newfid);
		p = gstring(p, ep, &f->txattrwalk.name);
		break;
	case RXATTRWALK:
		p = guint64(p, ep, &f->rxattrwalk.size);
		break;
	case TXATTRCREATE:
		p = guint32(p, ep, &f->txattrcreate.fid);
		p = gstring(p, ep, &f->txattrcreate.name);
		p = guint64(p, ep, &f->txattrcreate.attr_size);
		p = guint32(p, ep, &f->txattrcreate.flags);
		break;
	case RXATTRCREATE:
		// nothing
		break;
	case TREADDIR:
		p = guint32(p, ep, &f->treaddir.fid);
		p = guint64(p, ep, &f->treaddir.offset);
		p = guint32(p, ep, &f->treaddir.count);
		break;
	case RREADDIR:
		p = gdata(p, ep, &f->rreaddir.data);
		break;
	case TFYNC:
		p = guint32(p, ep, &f->tfync.fid);
		break;
	case RFSYNC:
		// nothing
		break;
	case TLOCK:
		p = guint32(p, ep, &f->tlock.fid);
		p = guint8(p, ep, &f->tlock.type);
		p = guint32(p, ep, &f->tlock.flags);
		p = guint64(p, ep, &f->tlock.start);
		p = guint64(p, ep, &f->tlock.length);
		p = guint32(p, ep, &f->tlock.proc_id);
		p = gstring(p, ep, &f->tlock.client_id);
		break;
	case RLOCK:
		p = guint8(p, ep, &f->rlock.status);
		break;
	case TGETLOCK:
		p = guint32(p, ep, &f->tgetlock.fid);
		p = guint8(p, ep, &f->tgetlock.type);
		p = guint64(p, ep, &f->tgetlock.start);
		p = guint64(p, ep, &f->tgetlock.length);
		p = guint32(p, ep, &f->tgetlock.proc_id);
		p = gstring(p, ep, &f->tgetlock.client_id);
		break;
	case RGETLOCK:
		p = guint8(p, ep, &f->rgetlock.type);
		p = guint64(p, ep, &f->rgetlock.start);
		p = guint64(p, ep, &f->rgetlock.length);
		p = guint32(p, ep, &f->rgetlock.proc_id);
		p = gstring(p, ep, &f->rgetlock.client_id);
		break;
	case TLINK:
		p = guint32(p, ep, &f->tlink.dfid);
		p = guint32(p, ep, &f->tlink.fid);
		p = gstring(p, ep, &f->tlink.name);
		break;
	case RLINK:
		// nothing
		break;
	case TMKDIR:
		p = guint32(p, ep, &f->tmkdir.dfid);
		p = gstring(p, ep, &f->tmkdir.name);
		p = guint32(p, ep, &f->tmkdir.mode);
		p = guint32(p, ep, &f->tmkdir.gid);
		break;
	case RMKDIR:
		p = guint8(p, ep, &f->rmkdir.qid.type);
		p = guint32(p, ep, &f->rmkdir.qid.version);
		p = guint64(p, ep, &f->rmkdir.qid.path);
		break;
	case TRENAMEAT:
		p = guint32(p, ep, &f->trenameat.olddirfid);
		p = gstring(p, ep, &f->trenameat.oldname);
		p = guint32(p, ep, &f->trenameat.newdirfid);
		p = gstring(p, ep, &f->trenameat.newname);
		break;
	case RRENAMEAT:
		// nothing
		break;
	case TUNLINKAT:
		p = guint32(p, ep, &f->tunlinkat.dirfd);
		p = gstring(p, ep, &f->tunlinkat.name);
		p = guint32(p, ep, &f->tunlinkat.flags);
		break;
	case RUNLINKAT:
		// nothing
		break;
	case RERROR:
		p = guint32(p, ep, &f->rerror.ecode);
		break;
	}

	if(p == NULL || p > ep)
		return -1;
	if(p == data + size)
		return size;
	return -1;
}

ssize_t
fcall_marshal(unsigned char* data, size_t msize, fcall_t* f)
{
	unsigned char *p, *ep;

	p  = data;
	ep = p + msize;
	if(p + 7 > ep)
		return -1;

	p = puint32(p, ep, msize);
	p = puint8(p, ep, f->type);
	p = puint16(p, ep, f->tag);
	if(p == NULL)
		return -1;

	switch(f->type) {
	case TVERSION:
		p = puint32(p, ep, f->tversion.msize);
		p = pstring(p, ep, f->tversion.version);
		break;
	case RVERSION:
		p = puint32(p, ep, f->rversion.msize);
		p = pstring(p, ep, f->rversion.version);
		break;
	case TFLUSH:
		p = puint16(p, ep, f->tflush.oldtag);
		break;
	case RFLUSH:
		// nothing
		break;
	case TWALK:
		p = puint32(p, ep, f->twalk.fid);
		p = puint32(p, ep, f->twalk.newfid);
		// p = pstrings(p, ep, f->twalk.names);
		break;
	case RWALK:
		// nothing
		break;
	case TCLUNK:
		p = puint32(p, ep, f->tclunk.fid);
		break;
	case RCLUNK:
		// nothing
		break;
	case TREMOVE:
		p = puint32(p, ep, f->tremove.fid);
		break;
	case RREMOVE:
		// nothing
		break;
	case TAUTH:
		p = puint32(p, ep, f->tauth.afid);
		p = pstring(p, ep, f->tauth.uname);
		p = pstring(p, ep, f->tauth.aname);
		p = puint32(p, ep, f->tauth.uid);
		break;
	case RAUTH:
		p = puint8(p, ep, f->rauth.qid.type);
		p = puint32(p, ep, f->rauth.qid.version);
		p = puint64(p, ep, f->rauth.qid.path);
		break;
	case TATTACH:
		p = puint32(p, ep, f->tattach.fid);
		p = puint32(p, ep, f->tattach.afid);
		p = pstring(p, ep, f->tattach.uname);
		p = pstring(p, ep, f->tattach.aname);
		p = puint32(p, ep, f->tattach.uid);
		break;
	case RATTACH:
		p = puint8(p, ep, f->rattach.qid.type);
		p = puint32(p, ep, f->rattach.qid.version);
		p = puint64(p, ep, f->rattach.qid.path);
		break;
	case TREAD:
		p = puint32(p, ep, f->tread.fid);
		p = puint64(p, ep, f->tread.offset);
		p = puint32(p, ep, f->tread.count);
		break;
	case RREAD:
		p = pdata(p, ep, f->rread.data);
		break;
	case TWRITE:
		p = puint32(p, ep, f->twrite.fid);
		p = puint64(p, ep, f->twrite.offset);
		p = pdata(p, ep, f->twrite.data);
		break;
	case RWRITE:
		p = puint32(p, ep, f->rwrite.count);
		break;
	case TSTATFS:
		p = puint32(p, ep, f->tstatfs.fid);
		break;
	case RSTATFS:
		p = puint32(p, ep, f->rstatfs.type);
		p = puint32(p, ep, f->rstatfs.bsize);
		p = puint64(p, ep, f->rstatfs.blocks);
		p = puint64(p, ep, f->rstatfs.bfree);
		p = puint64(p, ep, f->rstatfs.bavail);
		p = puint64(p, ep, f->rstatfs.files);
		p = puint64(p, ep, f->rstatfs.ffree);
		p = puint64(p, ep, f->rstatfs.fsid);
		p = puint32(p, ep, f->rstatfs.namelen);
		break;
	case TOPEN:
		p = puint32(p, ep, f->topen.fid);
		p = puint32(p, ep, f->topen.flags);
		break;
	case ROPEN:
		p = puint8(p, ep, f->ropen.qid.type);
		p = puint32(p, ep, f->ropen.qid.version);
		p = puint64(p, ep, f->ropen.qid.path);
		p = puint32(p, ep, f->ropen.iounit);
		break;
	case TCREATE:
		p = puint32(p, ep, f->tcreate.fid);
		p = pstring(p, ep, f->tcreate.name);
		p = puint32(p, ep, f->tcreate.flags);
		p = puint32(p, ep, f->tcreate.mode);
		p = puint32(p, ep, f->tcreate.gid);
		break;
	case RCREATE:
		p = puint8(p, ep, f->rcreate.qid.type);
		p = puint32(p, ep, f->rcreate.qid.version);
		p = puint64(p, ep, f->rcreate.qid.path);
		p = puint32(p, ep, f->rcreate.iounit);
		break;
	case TSYMLINK:
		p = puint32(p, ep, f->tsymlink.fid);
		p = pstring(p, ep, f->tsymlink.name);
		p = pstring(p, ep, f->tsymlink.symtgt);
		p = puint32(p, ep, f->tsymlink.gid);
		break;
	case RSYMLINK:
		p = puint8(p, ep, f->rsymlink.qid.type);
		p = puint32(p, ep, f->rsymlink.qid.version);
		p = puint64(p, ep, f->rsymlink.qid.path);
		break;
	case TMKNOD:
		p = puint32(p, ep, f->tmknod.dfid);
		p = pstring(p, ep, f->tmknod.name);
		p = puint32(p, ep, f->tmknod.mode);
		p = puint32(p, ep, f->tmknod.major);
		p = puint32(p, ep, f->tmknod.minor);
		p = puint32(p, ep, f->tmknod.gid);
		break;
	case RMKNOD:
		p = puint8(p, ep, f->rmknod.qid.type);
		p = puint32(p, ep, f->rmknod.qid.version);
		p = puint64(p, ep, f->rmknod.qid.path);
		break;
	case TRENAME:
		p = puint32(p, ep, f->trename.fid);
		p = puint32(p, ep, f->trename.dfid);
		p = pstring(p, ep, f->trename.name);
		break;
	case RRENAME:
		// nothing
		break;
	case TREADLINK:
		p = puint32(p, ep, f->treadlink.fid);
		break;
	case RREADLINK:
		p = pstring(p, ep, f->rreadlink.target);
		break;
	case TGETATTR:
		p = puint32(p, ep, f->tgetattr.fid);
		p = puint64(p, ep, f->tgetattr.request_mask);
		break;
	case RGETATTR:
		p = puint64(p, ep, f->rgetattr.valid);
		p = puint8(p, ep, f->rgetattr.qid.type);
		p = puint32(p, ep, f->rgetattr.qid.version);
		p = puint64(p, ep, f->rgetattr.qid.path);
		p = puint32(p, ep, f->rgetattr.mode);
		p = puint32(p, ep, f->rgetattr.uid);
		p = puint32(p, ep, f->rgetattr.gid);
		p = puint64(p, ep, f->rgetattr.nlink);
		p = puint64(p, ep, f->rgetattr.rdev);
		p = puint64(p, ep, f->rgetattr.size);
		p = puint64(p, ep, f->rgetattr.blksize);
		p = puint64(p, ep, f->rgetattr.blocks);
		p = puint64(p, ep, f->rgetattr.atime_sec);
		p = puint64(p, ep, f->rgetattr.atime_nsec);
		p = puint64(p, ep, f->rgetattr.mtime_sec);
		p = puint64(p, ep, f->rgetattr.mtime_nsec);
		p = puint64(p, ep, f->rgetattr.ctime_sec);
		p = puint64(p, ep, f->rgetattr.ctime_nsec);
		p = puint64(p, ep, f->rgetattr.btime_sec);
		p = puint64(p, ep, f->rgetattr.btime_nsec);
		p = puint64(p, ep, f->rgetattr.gen);
		p = puint64(p, ep, f->rgetattr.data_version);
		break;
	case TSETATTR:
		p = puint32(p, ep, f->tsetattr.fid);
		p = puint32(p, ep, f->tsetattr.valid);
		p = puint32(p, ep, f->tsetattr.mode);
		p = puint32(p, ep, f->tsetattr.uid);
		p = puint32(p, ep, f->tsetattr.gid);
		p = puint64(p, ep, f->tsetattr.size);
		p = puint64(p, ep, f->tsetattr.atime_sec);
		p = puint64(p, ep, f->tsetattr.atime_nsec);
		p = puint64(p, ep, f->tsetattr.mtime_sec);
		p = puint64(p, ep, f->tsetattr.mtime_nsec);
		break;
	case RSETATTR:
		// nothing
		break;
	case TXATTRWALK:
		p = puint32(p, ep, f->txattrwalk.fid);
		p = puint32(p, ep, f->txattrwalk.newfid);
		p = pstring(p, ep, f->txattrwalk.name);
		break;
	case RXATTRWALK:
		p = puint64(p, ep, f->rxattrwalk.size);
		break;
	case TXATTRCREATE:
		p = puint32(p, ep, f->txattrcreate.fid);
		p = pstring(p, ep, f->txattrcreate.name);
		p = puint64(p, ep, f->txattrcreate.attr_size);
		p = puint32(p, ep, f->txattrcreate.flags);
		break;
	case RXATTRCREATE:
		// nothing
		break;
	case TREADDIR:
		p = puint32(p, ep, f->treaddir.fid);
		p = puint64(p, ep, f->treaddir.offset);
		p = puint32(p, ep, f->treaddir.count);
		break;
	case RREADDIR:
		p = pdata(p, ep, f->rreaddir.data);
		break;
	case TFYNC:
		p = puint32(p, ep, f->tfync.fid);
		break;
	case RFSYNC:
		// nothing
		break;
	case TLOCK:
		p = puint32(p, ep, f->tlock.fid);
		p = puint8(p, ep, f->tlock.type);
		p = puint32(p, ep, f->tlock.flags);
		p = puint64(p, ep, f->tlock.start);
		p = puint64(p, ep, f->tlock.length);
		p = puint32(p, ep, f->tlock.proc_id);
		p = pstring(p, ep, f->tlock.client_id);
		break;
	case RLOCK:
		p = puint8(p, ep, f->rlock.status);
		break;
	case TGETLOCK:
		p = puint32(p, ep, f->tgetlock.fid);
		p = puint8(p, ep, f->tgetlock.type);
		p = puint64(p, ep, f->tgetlock.start);
		p = puint64(p, ep, f->tgetlock.length);
		p = puint32(p, ep, f->tgetlock.proc_id);
		p = pstring(p, ep, f->tgetlock.client_id);
		break;
	case RGETLOCK:
		p = puint8(p, ep, f->rgetlock.type);
		p = puint64(p, ep, f->rgetlock.start);
		p = puint64(p, ep, f->rgetlock.length);
		p = puint32(p, ep, f->rgetlock.proc_id);
		p = pstring(p, ep, f->rgetlock.client_id);
		break;
	case TLINK:
		p = puint32(p, ep, f->tlink.dfid);
		p = puint32(p, ep, f->tlink.fid);
		p = pstring(p, ep, f->tlink.name);
		break;
	case RLINK:
		// nothing
		break;
	case TMKDIR:
		p = puint32(p, ep, f->tmkdir.dfid);
		p = pstring(p, ep, f->tmkdir.name);
		p = puint32(p, ep, f->tmkdir.mode);
		p = puint32(p, ep, f->tmkdir.gid);
		break;
	case RMKDIR:
		p = puint8(p, ep, f->rmkdir.qid.type);
		p = puint32(p, ep, f->rmkdir.qid.version);
		p = puint64(p, ep, f->rmkdir.qid.path);
		break;
	case TRENAMEAT:
		p = puint32(p, ep, f->trenameat.olddirfid);
		p = pstring(p, ep, f->trenameat.oldname);
		p = puint32(p, ep, f->trenameat.newdirfid);
		p = pstring(p, ep, f->trenameat.newname);
		break;
	case RRENAMEAT:
		// nothing
		break;
	case TUNLINKAT:
		p = puint32(p, ep, f->tunlinkat.dirfd);
		p = pstring(p, ep, f->tunlinkat.name);
		p = puint32(p, ep, f->tunlinkat.flags);
		break;
	case RUNLINKAT:
		// nothing
		break;
	case RERROR:
		p = puint32(p, ep, f->rerror.ecode);
		break;
	}

	if(p == NULL || p > ep)
		return -1;
	if(p == data + msize)
		return msize;
	return -1;
}

static size_t
string_size(char* s)
{
	if(s == NULL)
		return 2;
	return 2 + strlen(s);
}

static size_t
data_size(char* s)
{
	if(s == NULL)
		return 4;
	return 4 + strlen(s);
}

ssize_t
fcall_size(fcall_t* f)
{
	uint32_t n = 4 + 1 + 2; // msize + type + tag

	switch(f->type) {
	case TVERSION:
		n += 4 + string_size(f->tversion.version);
		break;
	case RVERSION:
		n += 4 + string_size(f->rversion.version);
		break;
	case TFLUSH:
		n += 2;
		break;
	case RFLUSH:
		// nothing
		break;
	case TWALK:
		n += 4 + 4;
		break;
	case RWALK:
		// nothing
		break;
	case TCLUNK:
		n += 4;
		break;
	case RCLUNK:
		// nothing
		break;
	case TREMOVE:
		n += 4;
		break;
	case RREMOVE:
		// nothing
		break;
	case TAUTH:
		n += 4 + string_size(f->tauth.uname) + string_size(f->tauth.aname) + 4;
		break;
	case RAUTH:
		n += 13;
		break;
	case TATTACH:
		n += 4 + 4 + string_size(f->tattach.uname) +
		     string_size(f->tattach.aname) + 4;
		break;
	case RATTACH:
		n += 13;
		break;
	case TREAD:
		n += 4 + 8 + 4;
		break;
	case RREAD:
		// nothing
		break;
	case TWRITE:
		n += 4 + 8;
		break;
	case RWRITE:
		n += 4;
		break;
	case TSTATFS:
		n += 4;
		break;
	case RSTATFS:
		n += 4 + 4 + 8 + 8 + 8 + 8 + 8 + 8 + 4;
		break;
	case TOPEN:
		n += 4 + 4;
		break;
	case ROPEN:
		n += 13 + 4;
		break;
	case TCREATE:
		n += 4 + string_size(f->tcreate.name) + 4 + 4 + 4;
		break;
	case RCREATE:
		n += 13 + 4;
		break;
	case TSYMLINK:
		n += 4 + string_size(f->tsymlink.name) +
		     string_size(f->tsymlink.symtgt) + 4;
		break;
	case RSYMLINK:
		n += 13;
		break;
	case TMKNOD:
		n += 4 + string_size(f->tmknod.name) + 4 + 4 + 4 + 4;
		break;
	case RMKNOD:
		n += 13;
		break;
	case TRENAME:
		n += 4 + 4 + string_size(f->trename.name);
		break;
	case RRENAME:
		// nothing
		break;
	case TREADLINK:
		n += 4;
		break;
	case RREADLINK:
		n += string_size(f->rreadlink.target);
		break;
	case TGETATTR:
		n += 4 + 8;
		break;
	case RGETATTR:
		n += 8 + 13 + 4 + 4 + 4 + 8 + 8 + 8 + 8 + 8 + 8 + 8 + 8 + 8 + 8 + 8 +
		     8 + 8 + 8 + 8;
		break;
	case TSETATTR:
		n += 4 + 4 + 4 + 4 + 4 + 8 + 8 + 8 + 8 + 8;
		break;
	case RSETATTR:
		// nothing
		break;
	case TXATTRWALK:
		n += 4 + 4 + string_size(f->txattrwalk.name);
		break;
	case RXATTRWALK:
		n += 8;
		break;
	case TXATTRCREATE:
		n += 4 + string_size(f->txattrcreate.name) + 8 + 4;
		break;
	case RXATTRCREATE:
		// nothing
		break;
	case TREADDIR:
		n += 4 + 8 + 4;
		break;
	case RREADDIR:
		// nothing
		break;
	case TFYNC:
		n += 4;
		break;
	case RFSYNC:
		// nothing
		break;
	case TLOCK:
		n += 4 + 1 + 4 + 8 + 8 + 4 + string_size(f->tlock.client_id);
		break;
	case RLOCK:
		n += 1;
		break;
	case TGETLOCK:
		n += 4 + 1 + 8 + 8 + 4 + string_size(f->tgetlock.client_id);
		break;
	case RGETLOCK:
		n += 1 + 8 + 8 + 4 + string_size(f->rgetlock.client_id);
		break;
	case TLINK:
		n += 4 + 4 + string_size(f->tlink.name);
		break;
	case RLINK:
		// nothing
		break;
	case TMKDIR:
		n += 4 + string_size(f->tmkdir.name) + 4 + 4;
		break;
	case RMKDIR:
		n += 13;
		break;
	case TRENAMEAT:
		n += 4 + string_size(f->trenameat.oldname) + 4 +
		     string_size(f->trenameat.newname);
		break;
	case RRENAMEAT:
		// nothing
		break;
	case TUNLINKAT:
		n += 4 + string_size(f->tunlinkat.name) + 4;
		break;
	case RUNLINKAT:
		// nothing
		break;
	case RERROR:
		n += 4;
		break;
	default:
		return -1;
	}
	return n;
}