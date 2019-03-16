//go:generate make generate

package proto

import "fmt"

const (
	MaxMessageSize = MaxHeaderSize + MaxDataSize
	MaxNameSize    = 1<<16 - 1
	MaxDataSize    = 2 * 1024 * 1024
	MaxHeaderSize  = 24
	MaxElement     = 16
)

type Header struct {
	msize uint32
	Type  uint8
	Tag   uint16
}

type Fcall interface{}

// Qids are identifiers used by 9P2000.L servers to track file system
// entities. The type is used to differentiate semantics for operations
// on the entity. The path provides a server unique index for an entity
// (roughly analogous to an inode number), while the version is updated
// every time a file is modified and can be used to maintain cache
// coherency between clients and serves.
//
// See also: http://9p.io/magic/man2html/2/stat
//
type Qid struct {
	Type    uint8
	Version uint32
	Path    uint64
}

const (
	_QTDIR     = 0x80 // directories
	_QTAPPEND  = 0x40 // append only files
	_QTEXCL    = 0x20 // exclusive use files
	_QTMOUNT   = 0x10 // mounted channel
	_QTAUTH    = 0x08 // authentication file
	_QTTMP     = 0x04 // non-backed-up file
	_QTSYMLINK = 0x02 // symbolic link (Unix, 9P2000.u)
	_QTLINK    = 0x01 // hard link (Unix, 9P2000.u)
	_QTFILE    = 0x00
)

func (q Qid) String() string {
	t := ""
	if q.Type&_QTDIR != 0 {
		t += "d"
	}
	if q.Type&_QTAPPEND != 0 {
		t += "a"
	}
	if q.Type&_QTEXCL != 0 {
		t += "l"
	}
	if q.Type&_QTAUTH != 0 {
		t += "A"
	}
	if q.Type&_QTSYMLINK != 0 {
		t += "L"
	}
	return fmt.Sprintf("(%.16x %d %q)", q.Path, q.Version, t)
}
