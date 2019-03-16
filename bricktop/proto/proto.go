//
// GENERATED BY 'go run $GOPATH/github.com/azmodb/tool/generator -go -root .'; DO NOT EDIT!
//

package proto

import "fmt"

// 9P2000.L operations. There are 30 basic operations in 9P2000.L, paired
// as requests and responses. The one special case is ERROR as there is
// no NINEP_TERROR request for clients to transmit to the server, but the
// server may respond to any other request with an NINEP_RERROR.
//
// See also: http://9p.io/sys/man/5/INDEX.html
//
const (
	_TVERSION     uint8 = 100
	_RVERSION     uint8 = 101
	_TFLUSH       uint8 = 108
	_RFLUSH       uint8 = 109
	_TWALK        uint8 = 110
	_RWALK        uint8 = 111
	_TCLUNK       uint8 = 120
	_RCLUNK       uint8 = 121
	_TREMOVE      uint8 = 122
	_RREMOVE      uint8 = 123
	_TAUTH        uint8 = 102
	_RAUTH        uint8 = 103
	_TATTACH      uint8 = 104
	_RATTACH      uint8 = 105
	_TREAD        uint8 = 116
	_RREAD        uint8 = 117
	_TWRITE       uint8 = 118
	_RWRITE       uint8 = 119
	_TSTATFS      uint8 = 8
	_RSTATFS      uint8 = 9
	_TOPEN        uint8 = 12
	_ROPEN        uint8 = 13
	_TCREATE      uint8 = 14
	_RCREATE      uint8 = 15
	_TSYMLINK     uint8 = 16
	_RSYMLINK     uint8 = 17
	_TMKNOD       uint8 = 18
	_RMKNOD       uint8 = 19
	_TRENAME      uint8 = 20
	_RRENAME      uint8 = 21
	_TREADLINK    uint8 = 22
	_RREADLINK    uint8 = 23
	_TGETATTR     uint8 = 24
	_RGETATTR     uint8 = 25
	_TSETATTR     uint8 = 26
	_RSETATTR     uint8 = 27
	_TXATTRWALK   uint8 = 30
	_RXATTRWALK   uint8 = 31
	_TXATTRCREATE uint8 = 32
	_RXATTRCREATE uint8 = 33
	_TREADDIR     uint8 = 40
	_RREADDIR     uint8 = 41
	_TFYNC        uint8 = 50
	_RFSYNC       uint8 = 51
	_TLOCK        uint8 = 52
	_RLOCK        uint8 = 53
	_TGETLOCK     uint8 = 54
	_RGETLOCK     uint8 = 55
	_TLINK        uint8 = 70
	_RLINK        uint8 = 71
	_TMKDIR       uint8 = 72
	_RMKDIR       uint8 = 73
	_TRENAMEAT    uint8 = 74
	_RRENAMEAT    uint8 = 75
	_TUNLINKAT    uint8 = 76
	_RUNLINKAT    uint8 = 77
	_RERROR       uint8 = 7
)

func unmarshal(buf *Buffer, typ uint8, tag uint16, fcall Fcall) error {
	switch typ {
	case _TVERSION:
		f, ok := fcall.(*Tversion)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Msize)
		buf.DecodeString(&f.Version)
		return buf.Err()
	case _RVERSION:
		f, ok := fcall.(*Rversion)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Msize)
		buf.DecodeString(&f.Version)
		return buf.Err()
	case _TFLUSH:
		f, ok := fcall.(*Tflush)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint16(&f.Oldtag)
		return buf.Err()
	case _RFLUSH:
		f, ok := fcall.(*Rflush)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TWALK:
		f, ok := fcall.(*Twalk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Newfid)
		return buf.Err()
	case _RWALK:
		f, ok := fcall.(*Rwalk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TCLUNK:
		f, ok := fcall.(*Tclunk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		return buf.Err()
	case _RCLUNK:
		f, ok := fcall.(*Rclunk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TREMOVE:
		f, ok := fcall.(*Tremove)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		return buf.Err()
	case _RREMOVE:
		f, ok := fcall.(*Rremove)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TAUTH:
		f, ok := fcall.(*Tauth)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Afid)
		buf.DecodeString(&f.Uname)
		buf.DecodeString(&f.Aname)
		buf.DecodeUint32(&f.Uid)
		return buf.Err()
	case _RAUTH:
		f, ok := fcall.(*Rauth)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TATTACH:
		f, ok := fcall.(*Tattach)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Afid)
		buf.DecodeString(&f.Uname)
		buf.DecodeString(&f.Aname)
		buf.DecodeUint32(&f.Uid)
		return buf.Err()
	case _RATTACH:
		f, ok := fcall.(*Rattach)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TREAD:
		f, ok := fcall.(*Tread)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint64(&f.Offset)
		buf.DecodeUint32(&f.Count)
		return buf.Err()
	case _RREAD:
		f, ok := fcall.(*Rread)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TWRITE:
		f, ok := fcall.(*Twrite)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint64(&f.Offset)
		return buf.Err()
	case _RWRITE:
		f, ok := fcall.(*Rwrite)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Count)
		return buf.Err()
	case _TSTATFS:
		f, ok := fcall.(*Tstatfs)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		return buf.Err()
	case _RSTATFS:
		f, ok := fcall.(*Rstatfs)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Type)
		buf.DecodeUint32(&f.Bsize)
		buf.DecodeUint64(&f.Blocks)
		buf.DecodeUint64(&f.Bfree)
		buf.DecodeUint64(&f.Bavail)
		buf.DecodeUint64(&f.Files)
		buf.DecodeUint64(&f.Ffree)
		buf.DecodeUint64(&f.Fsid)
		buf.DecodeUint32(&f.Namelen)
		return buf.Err()
	case _TOPEN:
		f, ok := fcall.(*Topen)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Flags)
		return buf.Err()
	case _ROPEN:
		f, ok := fcall.(*Ropen)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Iounit)
		return buf.Err()
	case _TCREATE:
		f, ok := fcall.(*Tcreate)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeString(&f.Name)
		buf.DecodeUint32(&f.Flags)
		buf.DecodeUint32(&f.Mode)
		buf.DecodeUint32(&f.Gid)
		return buf.Err()
	case _RCREATE:
		f, ok := fcall.(*Rcreate)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Iounit)
		return buf.Err()
	case _TSYMLINK:
		f, ok := fcall.(*Tsymlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeString(&f.Name)
		buf.DecodeString(&f.Symtgt)
		buf.DecodeUint32(&f.Gid)
		return buf.Err()
	case _RSYMLINK:
		f, ok := fcall.(*Rsymlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TMKNOD:
		f, ok := fcall.(*Tmknod)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Dfid)
		buf.DecodeString(&f.Name)
		buf.DecodeUint32(&f.Mode)
		buf.DecodeUint32(&f.Major)
		buf.DecodeUint32(&f.Minor)
		buf.DecodeUint32(&f.Gid)
		return buf.Err()
	case _RMKNOD:
		f, ok := fcall.(*Rmknod)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TRENAME:
		f, ok := fcall.(*Trename)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Dfid)
		buf.DecodeString(&f.Name)
		return buf.Err()
	case _RRENAME:
		f, ok := fcall.(*Rrename)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TREADLINK:
		f, ok := fcall.(*Treadlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		return buf.Err()
	case _RREADLINK:
		f, ok := fcall.(*Rreadlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeString(&f.Target)
		return buf.Err()
	case _TGETATTR:
		f, ok := fcall.(*Tgetattr)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint64(&f.RequestMask)
		return buf.Err()
	case _RGETATTR:
		f, ok := fcall.(*Rgetattr)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint64(&f.Valid)
		buf.DecodeUint32(&f.Mode)
		buf.DecodeUint32(&f.Uid)
		buf.DecodeUint32(&f.Gid)
		buf.DecodeUint64(&f.Nlink)
		buf.DecodeUint64(&f.Rdev)
		buf.DecodeUint64(&f.Size)
		buf.DecodeUint64(&f.Blksize)
		buf.DecodeUint64(&f.Blocks)
		buf.DecodeUint64(&f.AtimeSec)
		buf.DecodeUint64(&f.AtimeNsec)
		buf.DecodeUint64(&f.MtimeSec)
		buf.DecodeUint64(&f.MtimeNsec)
		buf.DecodeUint64(&f.CtimeSec)
		buf.DecodeUint64(&f.CtimeNsec)
		buf.DecodeUint64(&f.BtimeSec)
		buf.DecodeUint64(&f.BtimeNsec)
		buf.DecodeUint64(&f.Gen)
		buf.DecodeUint64(&f.DataVersion)
		return buf.Err()
	case _TSETATTR:
		f, ok := fcall.(*Tsetattr)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Valid)
		buf.DecodeUint32(&f.Mode)
		buf.DecodeUint32(&f.Uid)
		buf.DecodeUint32(&f.Gid)
		buf.DecodeUint64(&f.Size)
		buf.DecodeUint64(&f.AtimeSec)
		buf.DecodeUint64(&f.AtimeNsec)
		buf.DecodeUint64(&f.MtimeSec)
		buf.DecodeUint64(&f.MtimeNsec)
		return buf.Err()
	case _RSETATTR:
		f, ok := fcall.(*Rsetattr)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TXATTRWALK:
		f, ok := fcall.(*Txattrwalk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint32(&f.Newfid)
		buf.DecodeString(&f.Name)
		return buf.Err()
	case _RXATTRWALK:
		f, ok := fcall.(*Rxattrwalk)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint64(&f.Size)
		return buf.Err()
	case _TXATTRCREATE:
		f, ok := fcall.(*Txattrcreate)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeString(&f.Name)
		buf.DecodeUint64(&f.AttrSize)
		buf.DecodeUint32(&f.Flags)
		return buf.Err()
	case _RXATTRCREATE:
		f, ok := fcall.(*Rxattrcreate)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TREADDIR:
		f, ok := fcall.(*Treaddir)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint64(&f.Offset)
		buf.DecodeUint32(&f.Count)
		return buf.Err()
	case _RREADDIR:
		f, ok := fcall.(*Rreaddir)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TFYNC:
		f, ok := fcall.(*Tfync)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		return buf.Err()
	case _RFSYNC:
		f, ok := fcall.(*Rfsync)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TLOCK:
		f, ok := fcall.(*Tlock)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint8(&f.Type)
		buf.DecodeUint32(&f.Flags)
		buf.DecodeUint64(&f.Start)
		buf.DecodeUint64(&f.Length)
		buf.DecodeUint32(&f.ProcId)
		buf.DecodeString(&f.ClientID)
		return buf.Err()
	case _RLOCK:
		f, ok := fcall.(*Rlock)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint8(&f.Status)
		return buf.Err()
	case _TGETLOCK:
		f, ok := fcall.(*Tgetlock)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Fid)
		buf.DecodeUint8(&f.Type)
		buf.DecodeUint64(&f.Start)
		buf.DecodeUint64(&f.Length)
		buf.DecodeUint32(&f.ProcId)
		buf.DecodeString(&f.ClientID)
		return buf.Err()
	case _RGETLOCK:
		f, ok := fcall.(*Rgetlock)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint8(&f.Type)
		buf.DecodeUint64(&f.Start)
		buf.DecodeUint64(&f.Length)
		buf.DecodeUint32(&f.ProcId)
		buf.DecodeString(&f.ClientID)
		return buf.Err()
	case _TLINK:
		f, ok := fcall.(*Tlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Dfid)
		buf.DecodeUint32(&f.Fid)
		buf.DecodeString(&f.Name)
		return buf.Err()
	case _RLINK:
		f, ok := fcall.(*Rlink)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TMKDIR:
		f, ok := fcall.(*Tmkdir)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Dfid)
		buf.DecodeString(&f.Name)
		buf.DecodeUint32(&f.Mode)
		buf.DecodeUint32(&f.Gid)
		return buf.Err()
	case _RMKDIR:
		f, ok := fcall.(*Rmkdir)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TRENAMEAT:
		f, ok := fcall.(*Trenameat)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Olddirfid)
		buf.DecodeString(&f.Oldname)
		buf.DecodeUint32(&f.Newdirfid)
		buf.DecodeString(&f.Newname)
		return buf.Err()
	case _RRENAMEAT:
		f, ok := fcall.(*Rrenameat)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _TUNLINKAT:
		f, ok := fcall.(*Tunlinkat)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Dirfd)
		buf.DecodeString(&f.Name)
		buf.DecodeUint32(&f.Flags)
		return buf.Err()
	case _RUNLINKAT:
		f, ok := fcall.(*Runlinkat)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		// nothing
		return buf.Err()
	case _RERROR:
		f, ok := fcall.(*Rerror)
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		buf.DecodeUint32(&f.Ecode)
		return buf.Err()
	}
	return fmt.Errorf("unknown 9P2000.L fcall type (%d)", typ)
}

func marshal(buf *Buffer, tag uint16, fcall Fcall) error {
	switch f := fcall.(type) {
	case *Tversion:
		buf.EncodeUint8(_TVERSION)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Msize)
		buf.EncodeString(f.Version)
		return buf.Err()
	case *Rversion:
		buf.EncodeUint8(_RVERSION)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Msize)
		buf.EncodeString(f.Version)
		return buf.Err()
	case *Tflush:
		buf.EncodeUint8(_TFLUSH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint16(f.Oldtag)
		return buf.Err()
	case *Rflush:
		buf.EncodeUint8(_RFLUSH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Twalk:
		buf.EncodeUint8(_TWALK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Newfid)
		return buf.Err()
	case *Rwalk:
		buf.EncodeUint8(_RWALK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tclunk:
		buf.EncodeUint8(_TCLUNK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		return buf.Err()
	case *Rclunk:
		buf.EncodeUint8(_RCLUNK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tremove:
		buf.EncodeUint8(_TREMOVE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		return buf.Err()
	case *Rremove:
		buf.EncodeUint8(_RREMOVE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tauth:
		buf.EncodeUint8(_TAUTH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Afid)
		buf.EncodeString(f.Uname)
		buf.EncodeString(f.Aname)
		buf.EncodeUint32(f.Uid)
		return buf.Err()
	case *Rauth:
		buf.EncodeUint8(_RAUTH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tattach:
		buf.EncodeUint8(_TATTACH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Afid)
		buf.EncodeString(f.Uname)
		buf.EncodeString(f.Aname)
		buf.EncodeUint32(f.Uid)
		return buf.Err()
	case *Rattach:
		buf.EncodeUint8(_RATTACH)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tread:
		buf.EncodeUint8(_TREAD)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint64(f.Offset)
		buf.EncodeUint32(f.Count)
		return buf.Err()
	case *Rread:
		buf.EncodeUint8(_RREAD)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Twrite:
		buf.EncodeUint8(_TWRITE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint64(f.Offset)
		return buf.Err()
	case *Rwrite:
		buf.EncodeUint8(_RWRITE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Count)
		return buf.Err()
	case *Tstatfs:
		buf.EncodeUint8(_TSTATFS)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		return buf.Err()
	case *Rstatfs:
		buf.EncodeUint8(_RSTATFS)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Type)
		buf.EncodeUint32(f.Bsize)
		buf.EncodeUint64(f.Blocks)
		buf.EncodeUint64(f.Bfree)
		buf.EncodeUint64(f.Bavail)
		buf.EncodeUint64(f.Files)
		buf.EncodeUint64(f.Ffree)
		buf.EncodeUint64(f.Fsid)
		buf.EncodeUint32(f.Namelen)
		return buf.Err()
	case *Topen:
		buf.EncodeUint8(_TOPEN)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Flags)
		return buf.Err()
	case *Ropen:
		buf.EncodeUint8(_ROPEN)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Iounit)
		return buf.Err()
	case *Tcreate:
		buf.EncodeUint8(_TCREATE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeString(f.Name)
		buf.EncodeUint32(f.Flags)
		buf.EncodeUint32(f.Mode)
		buf.EncodeUint32(f.Gid)
		return buf.Err()
	case *Rcreate:
		buf.EncodeUint8(_RCREATE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Iounit)
		return buf.Err()
	case *Tsymlink:
		buf.EncodeUint8(_TSYMLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeString(f.Name)
		buf.EncodeString(f.Symtgt)
		buf.EncodeUint32(f.Gid)
		return buf.Err()
	case *Rsymlink:
		buf.EncodeUint8(_RSYMLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tmknod:
		buf.EncodeUint8(_TMKNOD)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Dfid)
		buf.EncodeString(f.Name)
		buf.EncodeUint32(f.Mode)
		buf.EncodeUint32(f.Major)
		buf.EncodeUint32(f.Minor)
		buf.EncodeUint32(f.Gid)
		return buf.Err()
	case *Rmknod:
		buf.EncodeUint8(_RMKNOD)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Trename:
		buf.EncodeUint8(_TRENAME)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Dfid)
		buf.EncodeString(f.Name)
		return buf.Err()
	case *Rrename:
		buf.EncodeUint8(_RRENAME)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Treadlink:
		buf.EncodeUint8(_TREADLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		return buf.Err()
	case *Rreadlink:
		buf.EncodeUint8(_RREADLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeString(f.Target)
		return buf.Err()
	case *Tgetattr:
		buf.EncodeUint8(_TGETATTR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint64(f.RequestMask)
		return buf.Err()
	case *Rgetattr:
		buf.EncodeUint8(_RGETATTR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint64(f.Valid)
		buf.EncodeUint32(f.Mode)
		buf.EncodeUint32(f.Uid)
		buf.EncodeUint32(f.Gid)
		buf.EncodeUint64(f.Nlink)
		buf.EncodeUint64(f.Rdev)
		buf.EncodeUint64(f.Size)
		buf.EncodeUint64(f.Blksize)
		buf.EncodeUint64(f.Blocks)
		buf.EncodeUint64(f.AtimeSec)
		buf.EncodeUint64(f.AtimeNsec)
		buf.EncodeUint64(f.MtimeSec)
		buf.EncodeUint64(f.MtimeNsec)
		buf.EncodeUint64(f.CtimeSec)
		buf.EncodeUint64(f.CtimeNsec)
		buf.EncodeUint64(f.BtimeSec)
		buf.EncodeUint64(f.BtimeNsec)
		buf.EncodeUint64(f.Gen)
		buf.EncodeUint64(f.DataVersion)
		return buf.Err()
	case *Tsetattr:
		buf.EncodeUint8(_TSETATTR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Valid)
		buf.EncodeUint32(f.Mode)
		buf.EncodeUint32(f.Uid)
		buf.EncodeUint32(f.Gid)
		buf.EncodeUint64(f.Size)
		buf.EncodeUint64(f.AtimeSec)
		buf.EncodeUint64(f.AtimeNsec)
		buf.EncodeUint64(f.MtimeSec)
		buf.EncodeUint64(f.MtimeNsec)
		return buf.Err()
	case *Rsetattr:
		buf.EncodeUint8(_RSETATTR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Txattrwalk:
		buf.EncodeUint8(_TXATTRWALK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint32(f.Newfid)
		buf.EncodeString(f.Name)
		return buf.Err()
	case *Rxattrwalk:
		buf.EncodeUint8(_RXATTRWALK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint64(f.Size)
		return buf.Err()
	case *Txattrcreate:
		buf.EncodeUint8(_TXATTRCREATE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeString(f.Name)
		buf.EncodeUint64(f.AttrSize)
		buf.EncodeUint32(f.Flags)
		return buf.Err()
	case *Rxattrcreate:
		buf.EncodeUint8(_RXATTRCREATE)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Treaddir:
		buf.EncodeUint8(_TREADDIR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint64(f.Offset)
		buf.EncodeUint32(f.Count)
		return buf.Err()
	case *Rreaddir:
		buf.EncodeUint8(_RREADDIR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tfync:
		buf.EncodeUint8(_TFYNC)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		return buf.Err()
	case *Rfsync:
		buf.EncodeUint8(_RFSYNC)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tlock:
		buf.EncodeUint8(_TLOCK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint8(f.Type)
		buf.EncodeUint32(f.Flags)
		buf.EncodeUint64(f.Start)
		buf.EncodeUint64(f.Length)
		buf.EncodeUint32(f.ProcId)
		buf.EncodeString(f.ClientID)
		return buf.Err()
	case *Rlock:
		buf.EncodeUint8(_RLOCK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint8(f.Status)
		return buf.Err()
	case *Tgetlock:
		buf.EncodeUint8(_TGETLOCK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Fid)
		buf.EncodeUint8(f.Type)
		buf.EncodeUint64(f.Start)
		buf.EncodeUint64(f.Length)
		buf.EncodeUint32(f.ProcId)
		buf.EncodeString(f.ClientID)
		return buf.Err()
	case *Rgetlock:
		buf.EncodeUint8(_RGETLOCK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint8(f.Type)
		buf.EncodeUint64(f.Start)
		buf.EncodeUint64(f.Length)
		buf.EncodeUint32(f.ProcId)
		buf.EncodeString(f.ClientID)
		return buf.Err()
	case *Tlink:
		buf.EncodeUint8(_TLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Dfid)
		buf.EncodeUint32(f.Fid)
		buf.EncodeString(f.Name)
		return buf.Err()
	case *Rlink:
		buf.EncodeUint8(_RLINK)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tmkdir:
		buf.EncodeUint8(_TMKDIR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Dfid)
		buf.EncodeString(f.Name)
		buf.EncodeUint32(f.Mode)
		buf.EncodeUint32(f.Gid)
		return buf.Err()
	case *Rmkdir:
		buf.EncodeUint8(_RMKDIR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Trenameat:
		buf.EncodeUint8(_TRENAMEAT)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Olddirfid)
		buf.EncodeString(f.Oldname)
		buf.EncodeUint32(f.Newdirfid)
		buf.EncodeString(f.Newname)
		return buf.Err()
	case *Rrenameat:
		buf.EncodeUint8(_RRENAMEAT)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Tunlinkat:
		buf.EncodeUint8(_TUNLINKAT)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Dirfd)
		buf.EncodeString(f.Name)
		buf.EncodeUint32(f.Flags)
		return buf.Err()
	case *Runlinkat:
		buf.EncodeUint8(_RUNLINKAT)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		// nothing
		return buf.Err()
	case *Rerror:
		buf.EncodeUint8(_RERROR)
		f.tag = tag
		buf.EncodeUint16(f.tag)
		buf.EncodeUint32(f.Ecode)
		return buf.Err()
	}
	return fmt.Errorf("unknown 9P2000.L fcall (%T)", fcall)
}

type Tversion struct {
	tag     uint16
	Msize   uint32
	Version string
}

//func (f Tversion) FcallType() uint8 { return _TVERSION }
//func (f Tversion) FcallMessage()    {}
//func (f Tversion) Tag() uint16      { return f.tag }

type Rversion struct {
	tag     uint16
	Msize   uint32
	Version string
}

//func (f Rversion) FcallType() uint8 { return _RVERSION }
//func (f Rversion) FcallMessage()    {}
//func (f Rversion) Tag() uint16      { return f.tag }

type Tflush struct {
	tag    uint16
	Oldtag uint16
}

//func (f Tflush) FcallType() uint8 { return _TFLUSH }
//func (f Tflush) FcallMessage()    {}
//func (f Tflush) Tag() uint16      { return f.tag }

type Rflush struct {
	tag uint16
}

//func (f Rflush) FcallType() uint8 { return _RFLUSH }
//func (f Rflush) FcallMessage()    {}
//func (f Rflush) Tag() uint16      { return f.tag }

type Twalk struct {
	tag    uint16
	Fid    uint32
	Newfid uint32
	Names  string
}

//func (f Twalk) FcallType() uint8 { return _TWALK }
//func (f Twalk) FcallMessage()    {}
//func (f Twalk) Tag() uint16      { return f.tag }

type Rwalk struct {
	tag uint16
	Qid Qid
}

//func (f Rwalk) FcallType() uint8 { return _RWALK }
//func (f Rwalk) FcallMessage()    {}
//func (f Rwalk) Tag() uint16      { return f.tag }

type Tclunk struct {
	tag uint16
	Fid uint32
}

//func (f Tclunk) FcallType() uint8 { return _TCLUNK }
//func (f Tclunk) FcallMessage()    {}
//func (f Tclunk) Tag() uint16      { return f.tag }

type Rclunk struct {
	tag uint16
}

//func (f Rclunk) FcallType() uint8 { return _RCLUNK }
//func (f Rclunk) FcallMessage()    {}
//func (f Rclunk) Tag() uint16      { return f.tag }

type Tremove struct {
	tag uint16
	Fid uint32
}

//func (f Tremove) FcallType() uint8 { return _TREMOVE }
//func (f Tremove) FcallMessage()    {}
//func (f Tremove) Tag() uint16      { return f.tag }

type Rremove struct {
	tag uint16
}

//func (f Rremove) FcallType() uint8 { return _RREMOVE }
//func (f Rremove) FcallMessage()    {}
//func (f Rremove) Tag() uint16      { return f.tag }

type Tauth struct {
	tag   uint16
	Afid  uint32
	Uname string
	Aname string
	Uid   uint32
}

//func (f Tauth) FcallType() uint8 { return _TAUTH }
//func (f Tauth) FcallMessage()    {}
//func (f Tauth) Tag() uint16      { return f.tag }

type Rauth struct {
	tag uint16
	Qid Qid
}

//func (f Rauth) FcallType() uint8 { return _RAUTH }
//func (f Rauth) FcallMessage()    {}
//func (f Rauth) Tag() uint16      { return f.tag }

type Tattach struct {
	tag   uint16
	Fid   uint32
	Afid  uint32
	Uname string
	Aname string
	Uid   uint32
}

//func (f Tattach) FcallType() uint8 { return _TATTACH }
//func (f Tattach) FcallMessage()    {}
//func (f Tattach) Tag() uint16      { return f.tag }

type Rattach struct {
	tag uint16
	Qid Qid
}

//func (f Rattach) FcallType() uint8 { return _RATTACH }
//func (f Rattach) FcallMessage()    {}
//func (f Rattach) Tag() uint16      { return f.tag }

type Tread struct {
	tag    uint16
	Fid    uint32
	Offset uint64
	Count  uint32
}

//func (f Tread) FcallType() uint8 { return _TREAD }
//func (f Tread) FcallMessage()    {}
//func (f Tread) Tag() uint16      { return f.tag }

type Rread struct {
	tag  uint16
	Data []byte
}

//func (f Rread) FcallType() uint8 { return _RREAD }
//func (f Rread) FcallMessage()    {}
//func (f Rread) Tag() uint16      { return f.tag }

type Twrite struct {
	tag    uint16
	Fid    uint32
	Offset uint64
	Data   []byte
}

//func (f Twrite) FcallType() uint8 { return _TWRITE }
//func (f Twrite) FcallMessage()    {}
//func (f Twrite) Tag() uint16      { return f.tag }

type Rwrite struct {
	tag   uint16
	Count uint32
}

//func (f Rwrite) FcallType() uint8 { return _RWRITE }
//func (f Rwrite) FcallMessage()    {}
//func (f Rwrite) Tag() uint16      { return f.tag }

type Tstatfs struct {
	tag uint16
	Fid uint32
}

//func (f Tstatfs) FcallType() uint8 { return _TSTATFS }
//func (f Tstatfs) FcallMessage()    {}
//func (f Tstatfs) Tag() uint16      { return f.tag }

type Rstatfs struct {
	tag     uint16
	Type    uint32
	Bsize   uint32
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Fsid    uint64
	Namelen uint32
}

//func (f Rstatfs) FcallType() uint8 { return _RSTATFS }
//func (f Rstatfs) FcallMessage()    {}
//func (f Rstatfs) Tag() uint16      { return f.tag }

type Topen struct {
	tag   uint16
	Fid   uint32
	Flags uint32
}

//func (f Topen) FcallType() uint8 { return _TOPEN }
//func (f Topen) FcallMessage()    {}
//func (f Topen) Tag() uint16      { return f.tag }

type Ropen struct {
	tag    uint16
	Qid    Qid
	Iounit uint32
}

//func (f Ropen) FcallType() uint8 { return _ROPEN }
//func (f Ropen) FcallMessage()    {}
//func (f Ropen) Tag() uint16      { return f.tag }

type Tcreate struct {
	tag   uint16
	Fid   uint32
	Name  string
	Flags uint32
	Mode  uint32
	Gid   uint32
}

//func (f Tcreate) FcallType() uint8 { return _TCREATE }
//func (f Tcreate) FcallMessage()    {}
//func (f Tcreate) Tag() uint16      { return f.tag }

type Rcreate struct {
	tag    uint16
	Qid    Qid
	Iounit uint32
}

//func (f Rcreate) FcallType() uint8 { return _RCREATE }
//func (f Rcreate) FcallMessage()    {}
//func (f Rcreate) Tag() uint16      { return f.tag }

type Tsymlink struct {
	tag    uint16
	Fid    uint32
	Name   string
	Symtgt string
	Gid    uint32
}

//func (f Tsymlink) FcallType() uint8 { return _TSYMLINK }
//func (f Tsymlink) FcallMessage()    {}
//func (f Tsymlink) Tag() uint16      { return f.tag }

type Rsymlink struct {
	tag uint16
	Qid Qid
}

//func (f Rsymlink) FcallType() uint8 { return _RSYMLINK }
//func (f Rsymlink) FcallMessage()    {}
//func (f Rsymlink) Tag() uint16      { return f.tag }

type Tmknod struct {
	tag   uint16
	Dfid  uint32
	Name  string
	Mode  uint32
	Major uint32
	Minor uint32
	Gid   uint32
}

//func (f Tmknod) FcallType() uint8 { return _TMKNOD }
//func (f Tmknod) FcallMessage()    {}
//func (f Tmknod) Tag() uint16      { return f.tag }

type Rmknod struct {
	tag uint16
	Qid Qid
}

//func (f Rmknod) FcallType() uint8 { return _RMKNOD }
//func (f Rmknod) FcallMessage()    {}
//func (f Rmknod) Tag() uint16      { return f.tag }

type Trename struct {
	tag  uint16
	Fid  uint32
	Dfid uint32
	Name string
}

//func (f Trename) FcallType() uint8 { return _TRENAME }
//func (f Trename) FcallMessage()    {}
//func (f Trename) Tag() uint16      { return f.tag }

type Rrename struct {
	tag uint16
}

//func (f Rrename) FcallType() uint8 { return _RRENAME }
//func (f Rrename) FcallMessage()    {}
//func (f Rrename) Tag() uint16      { return f.tag }

type Treadlink struct {
	tag uint16
	Fid uint32
}

//func (f Treadlink) FcallType() uint8 { return _TREADLINK }
//func (f Treadlink) FcallMessage()    {}
//func (f Treadlink) Tag() uint16      { return f.tag }

type Rreadlink struct {
	tag    uint16
	Target string
}

//func (f Rreadlink) FcallType() uint8 { return _RREADLINK }
//func (f Rreadlink) FcallMessage()    {}
//func (f Rreadlink) Tag() uint16      { return f.tag }

type Tgetattr struct {
	tag         uint16
	Fid         uint32
	RequestMask uint64
}

//func (f Tgetattr) FcallType() uint8 { return _TGETATTR }
//func (f Tgetattr) FcallMessage()    {}
//func (f Tgetattr) Tag() uint16      { return f.tag }

type Rgetattr struct {
	tag         uint16
	Valid       uint64
	Qid         Qid
	Mode        uint32
	Uid         uint32
	Gid         uint32
	Nlink       uint64
	Rdev        uint64
	Size        uint64
	Blksize     uint64
	Blocks      uint64
	AtimeSec    uint64
	AtimeNsec   uint64
	MtimeSec    uint64
	MtimeNsec   uint64
	CtimeSec    uint64
	CtimeNsec   uint64
	BtimeSec    uint64
	BtimeNsec   uint64
	Gen         uint64
	DataVersion uint64
}

//func (f Rgetattr) FcallType() uint8 { return _RGETATTR }
//func (f Rgetattr) FcallMessage()    {}
//func (f Rgetattr) Tag() uint16      { return f.tag }

type Tsetattr struct {
	tag       uint16
	Fid       uint32
	Valid     uint32
	Mode      uint32
	Uid       uint32
	Gid       uint32
	Size      uint64
	AtimeSec  uint64
	AtimeNsec uint64
	MtimeSec  uint64
	MtimeNsec uint64
}

//func (f Tsetattr) FcallType() uint8 { return _TSETATTR }
//func (f Tsetattr) FcallMessage()    {}
//func (f Tsetattr) Tag() uint16      { return f.tag }

type Rsetattr struct {
	tag uint16
}

//func (f Rsetattr) FcallType() uint8 { return _RSETATTR }
//func (f Rsetattr) FcallMessage()    {}
//func (f Rsetattr) Tag() uint16      { return f.tag }

type Txattrwalk struct {
	tag    uint16
	Fid    uint32
	Newfid uint32
	Name   string
}

//func (f Txattrwalk) FcallType() uint8 { return _TXATTRWALK }
//func (f Txattrwalk) FcallMessage()    {}
//func (f Txattrwalk) Tag() uint16      { return f.tag }

type Rxattrwalk struct {
	tag  uint16
	Size uint64
}

//func (f Rxattrwalk) FcallType() uint8 { return _RXATTRWALK }
//func (f Rxattrwalk) FcallMessage()    {}
//func (f Rxattrwalk) Tag() uint16      { return f.tag }

type Txattrcreate struct {
	tag      uint16
	Fid      uint32
	Name     string
	AttrSize uint64
	Flags    uint32
}

//func (f Txattrcreate) FcallType() uint8 { return _TXATTRCREATE }
//func (f Txattrcreate) FcallMessage()    {}
//func (f Txattrcreate) Tag() uint16      { return f.tag }

type Rxattrcreate struct {
	tag uint16
}

//func (f Rxattrcreate) FcallType() uint8 { return _RXATTRCREATE }
//func (f Rxattrcreate) FcallMessage()    {}
//func (f Rxattrcreate) Tag() uint16      { return f.tag }

type Treaddir struct {
	tag    uint16
	Fid    uint32
	Offset uint64
	Count  uint32
}

//func (f Treaddir) FcallType() uint8 { return _TREADDIR }
//func (f Treaddir) FcallMessage()    {}
//func (f Treaddir) Tag() uint16      { return f.tag }

type Rreaddir struct {
	tag  uint16
	Data []byte
}

//func (f Rreaddir) FcallType() uint8 { return _RREADDIR }
//func (f Rreaddir) FcallMessage()    {}
//func (f Rreaddir) Tag() uint16      { return f.tag }

type Tfync struct {
	tag uint16
	Fid uint32
}

//func (f Tfync) FcallType() uint8 { return _TFYNC }
//func (f Tfync) FcallMessage()    {}
//func (f Tfync) Tag() uint16      { return f.tag }

type Rfsync struct {
	tag uint16
}

//func (f Rfsync) FcallType() uint8 { return _RFSYNC }
//func (f Rfsync) FcallMessage()    {}
//func (f Rfsync) Tag() uint16      { return f.tag }

type Tlock struct {
	tag      uint16
	Fid      uint32
	Type     uint8
	Flags    uint32
	Start    uint64
	Length   uint64
	ProcId   uint32
	ClientID string
}

//func (f Tlock) FcallType() uint8 { return _TLOCK }
//func (f Tlock) FcallMessage()    {}
//func (f Tlock) Tag() uint16      { return f.tag }

type Rlock struct {
	tag    uint16
	Status uint8
}

//func (f Rlock) FcallType() uint8 { return _RLOCK }
//func (f Rlock) FcallMessage()    {}
//func (f Rlock) Tag() uint16      { return f.tag }

type Tgetlock struct {
	tag      uint16
	Fid      uint32
	Type     uint8
	Start    uint64
	Length   uint64
	ProcId   uint32
	ClientID string
}

//func (f Tgetlock) FcallType() uint8 { return _TGETLOCK }
//func (f Tgetlock) FcallMessage()    {}
//func (f Tgetlock) Tag() uint16      { return f.tag }

type Rgetlock struct {
	tag      uint16
	Type     uint8
	Start    uint64
	Length   uint64
	ProcId   uint32
	ClientID string
}

//func (f Rgetlock) FcallType() uint8 { return _RGETLOCK }
//func (f Rgetlock) FcallMessage()    {}
//func (f Rgetlock) Tag() uint16      { return f.tag }

type Tlink struct {
	tag  uint16
	Dfid uint32
	Fid  uint32
	Name string
}

//func (f Tlink) FcallType() uint8 { return _TLINK }
//func (f Tlink) FcallMessage()    {}
//func (f Tlink) Tag() uint16      { return f.tag }

type Rlink struct {
	tag uint16
}

//func (f Rlink) FcallType() uint8 { return _RLINK }
//func (f Rlink) FcallMessage()    {}
//func (f Rlink) Tag() uint16      { return f.tag }

type Tmkdir struct {
	tag  uint16
	Dfid uint32
	Name string
	Mode uint32
	Gid  uint32
}

//func (f Tmkdir) FcallType() uint8 { return _TMKDIR }
//func (f Tmkdir) FcallMessage()    {}
//func (f Tmkdir) Tag() uint16      { return f.tag }

type Rmkdir struct {
	tag uint16
	Qid Qid
}

//func (f Rmkdir) FcallType() uint8 { return _RMKDIR }
//func (f Rmkdir) FcallMessage()    {}
//func (f Rmkdir) Tag() uint16      { return f.tag }

type Trenameat struct {
	tag       uint16
	Olddirfid uint32
	Oldname   string
	Newdirfid uint32
	Newname   string
}

//func (f Trenameat) FcallType() uint8 { return _TRENAMEAT }
//func (f Trenameat) FcallMessage()    {}
//func (f Trenameat) Tag() uint16      { return f.tag }

type Rrenameat struct {
	tag uint16
}

//func (f Rrenameat) FcallType() uint8 { return _RRENAMEAT }
//func (f Rrenameat) FcallMessage()    {}
//func (f Rrenameat) Tag() uint16      { return f.tag }

type Tunlinkat struct {
	tag   uint16
	Dirfd uint32
	Name  string
	Flags uint32
}

//func (f Tunlinkat) FcallType() uint8 { return _TUNLINKAT }
//func (f Tunlinkat) FcallMessage()    {}
//func (f Tunlinkat) Tag() uint16      { return f.tag }

type Runlinkat struct {
	tag uint16
}

//func (f Runlinkat) FcallType() uint8 { return _RUNLINKAT }
//func (f Runlinkat) FcallMessage()    {}
//func (f Runlinkat) Tag() uint16      { return f.tag }

type Rerror struct {
	tag   uint16
	Ecode uint32
}

//func (f Rerror) FcallType() uint8 { return _RERROR }
//func (f Rerror) FcallMessage()    {}
//func (f Rerror) Tag() uint16      { return f.tag }