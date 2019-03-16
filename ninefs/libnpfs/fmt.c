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
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include "npfs.h"
#include "npfsimpl.h"

static int
np_printperm(FILE *f, int perm)
{
	int n;
	char b[10];

	n = 0;
	if (perm & DMDIR)
		b[n++] = 'd';
	if (perm & DMAPPEND)
		b[n++] = 'a';
	if (perm & DMAUTH)
		b[n++] = 'A';
	if (perm & DMEXCL)
		b[n++] = 'l';
	if (perm & DMTMP)
		b[n++] = 't';
	if (perm & DMDEVICE)
		b[n++] = 'D';
	if (perm & DMSOCKET)
		b[n++] = 'S';
	if (perm & DMNAMEDPIPE)
		b[n++] = 'P';
        if (perm & DMSYMLINK)
                b[n++] = 'L';
        b[n] = '\0';

        return fprintf(f, "%s%03o", b, perm&0777);
}             

static int
np_printqid(FILE *f, Npqid *q)
{
	int n;
	char buf[10];

	n = 0;
	if (q->type & QTDIR)
		buf[n++] = 'd';
	if (q->type & QTAPPEND)
		buf[n++] = 'a';
	if (q->type & QTAUTH)
		buf[n++] = 'A';
	if (q->type & QTEXCL)
		buf[n++] = 'l';
	if (q->type & QTTMP)
		buf[n++] = 't';
	if (q->type & QTSYMLINK)
		buf[n++] = 'L';
	buf[n] = '\0';

	return fprintf(f, " (%.16llx %x '%s')", (unsigned long long)q->path, q->version, buf);
}

int
np_printstat(FILE *f, Npstat *st, int dotu)
{
	int n;

	n = fprintf(f, "'%.*s' '%.*s' '%.*s' '%.*s' q ", 
		st->name.len, st->name.str, st->uid.len, st->uid.str,
		st->gid.len, st->gid.str, st->muid.len, st->muid.str);

	n += np_printqid(f, &st->qid);
	n += fprintf(f, " m ");
	n += np_printperm(f, st->mode);
	n += fprintf(f, " at %d mt %d l %llu t %d d %d",
		st->atime, st->mtime, (unsigned long long)st->length, st->type, st->dev);
	if (dotu)
		n += fprintf(f, " ext '%.*s'", st->extension.len, 
			st->extension.str);

	return n;
}

int
np_dump(FILE *f, u8 *data, int datalen)
{
	int i, n;

	i = n = 0;
	while (i < datalen) {
		n += fprintf(f, "%02x", data[i]);
		if (i%4 == 3)
			n += fprintf(f, " ");
		if (i%32 == 31)
			n += fprintf(f, "\n");

		i++;
	}
	if(n > 0)
		n += fprintf(f, "\n");

	return n;
}

static int
np_printdata(FILE *f, u8 *buf, int buflen)
{
	return np_dump(f, buf, buflen<64?buflen:64);
}

int
np_dumpdata(u8 *buf, int buflen)
{
	return np_dump(stderr, buf, buflen);
}

int
np_printfcall(FILE *f, Npfcall *fc, int dotu) 
{
	int i, ret, type, fid, tag;

	if (!fc)
		return fprintf(f, "NULL");

	type = fc->type;
	fid = fc->fid;
	tag = fc->tag;

	ret = 0;
	switch (type) {
	case TVERSION:
		ret += fprintf(f, "TVERSION tag %u msize %u version '%.*s'", 
			tag, fc->msize, fc->version.len, fc->version.str);
		break;

	case RVERSION:
		ret += fprintf(f, "RVERSION tag %u msize %u version '%.*s'", 
			tag, fc->msize, fc->version.len, fc->version.str);
		break;

	case TAUTH:
		ret += fprintf(f, "TAUTH tag %u afid %d uname %.*s aname %.*s",
			tag, fc->afid, fc->uname.len, fc->uname.str, 
			fc->aname.len, fc->aname.str);
		break;

	case RAUTH:
		ret += fprintf(f, "RAUTH tag %u qid ", tag); 
		np_printqid(f, &fc->qid);
		break;

	case TATTACH:
		ret += fprintf(f, "TATTACH tag %u fid %d afid %d uname %.*s aname %.*s",
			tag, fid, fc->afid, fc->uname.len, fc->uname.str, 
			fc->aname.len, fc->aname.str);
		break;

	case RATTACH:
		ret += fprintf(f, "RATTACH tag %u qid ", tag); 
		np_printqid(f, &fc->qid);
		break;

	case RERROR:
		ret += fprintf(f, "RERROR tag %u ename %.*s", tag, 
			fc->ename.len, fc->ename.str);
		if (dotu)
			ret += fprintf(f, " ecode %d", fc->ecode);
		break;

	case TFLUSH:
		ret += fprintf(f, "TFLUSH tag %u oldtag %u", tag, fc->oldtag);
		break;

	case RFLUSH:
		ret += fprintf(f, "RFLUSH tag %u", tag);
		break;

	case TWALK:
		ret += fprintf(f, "TWALK tag %u fid %d newfid %d nwname %d", 
			tag, fid, fc->newfid, fc->nwname);
		for(i = 0; i < fc->nwname; i++)
			ret += fprintf(f, " '%.*s'", fc->wnames[i].len, 
				fc->wnames[i].str);
		break;
		
	case RWALK:
		ret += fprintf(f, "RWALK tag %u nwqid %d", tag, fc->nwqid);
		for(i = 0; i < fc->nwqid; i++)
			ret += np_printqid(f, &fc->wqids[i]);
		break;
		
	case TOPEN:
		ret += fprintf(f, "TOPEN tag %u fid %d mode %d", tag, fid, 
			fc->mode);
		break;
		
	case ROPEN:
		ret += fprintf(f, "ROPEN tag %u", tag);
		ret += np_printqid(f, &fc->qid);
		ret += fprintf(f, " iounit %d", fc->iounit);
		break;
		
	case TCREATE:
		ret += fprintf(f, "TCREATE tag %u fid %d name %.*s perm ",
			tag, fid, fc->name.len, fc->name.str);
		ret += np_printperm(f, fc->perm);
		ret += fprintf(f, " mode %d", fc->mode);
		if (dotu)
			ret += fprintf(f, " ext %.*s", fc->extension.len,
				fc->extension.str);
		break;
		
	case RCREATE:
		ret += fprintf(f, "RCREATE tag %u", tag);
		ret += np_printqid(f, &fc->qid);
		ret += fprintf(f, " iounit %d", fc->iounit);
		break;
		
	case TREAD:
		ret += fprintf(f, "TREAD tag %u fid %d offset %llu count %u", 
			tag, fid, (unsigned long long)fc->offset, fc->count);
		break;
		
	case RREAD:
		ret += fprintf(f, "RREAD tag %u count %u data ", tag, fc->count);
		ret += np_printdata(f, fc->data, fc->count);
		break;
		
	case TWRITE:
		ret += fprintf(f, "TWRITE tag %u fid %d offset %llu count %u data ",
			tag, fid, (unsigned long long)fc->offset, fc->count);
		ret += np_printdata(f, fc->data, fc->count);
		break;
		
	case RWRITE:
		ret += fprintf(f, "RWRITE tag %u count %u", tag, fc->count);
		break;
		
	case TCLUNK:
		ret += fprintf(f, "TCLUNK tag %u fid %d", tag, fid);
		break;
		
	case RCLUNK:
		ret += fprintf(f, "RCLUNK tag %u", tag);
		break;
		
	case TREMOVE:
		ret += fprintf(f, "TREMOVE tag %u fid %d", tag, fid);
		break;
		
	case RREMOVE:
		ret += fprintf(f, "RREMOVE tag %u", tag);
		break;
		
	case TSTAT:
		ret += fprintf(f, "TSTAT tag %u fid %d", tag, fid);
		break;
		
	case RSTAT:
		ret += fprintf(f, "RSTAT tag %u ", tag);
		ret += np_printstat(f, &fc->stat, dotu);
		break;
		
	case TWSTAT:
		ret += fprintf(f, "TWSTAT tag %u fid %d ", tag, fid);
		ret += np_printstat(f, &fc->stat, dotu);
		break;
		
	case RWSTAT:
		ret += fprintf(f, "RWSTAT tag %u", tag);
		break;

	default:
		ret += fprintf(f, "unknown type %d", type);
		break;
	}

	return ret;
}
