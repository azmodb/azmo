// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/azmodb/bricktop/internal/parser"
)

func main() {
	var (
		proto = []string{"github.com", "azmodb", "bricktop", "internal", "protocol"}

		jsonsource = flag.Bool("json", false, "generate json source-code")
		root       = flag.String("root", "/tmp", "generated file root")

		gosource = flag.Bool("go", false, "generate go source-code")
		csource  = flag.Bool("c", false, "generate c source-code")
		hsource  = flag.Bool("h", false, "generate h source-code")
	)
	flag.Parse()

	protoFile := make([]string, len(proto)+2)
	protoFile[0] = os.Getenv("GOPATH")
	protoFile[1] = "src"
	copy(protoFile[2:], proto)

	f, err := os.Open(filepath.Join(protoFile...))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	expr, err := parser.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	if *gosource {
		file := filepath.Join(*root, "proto.go")
		if err = genGoSource(file, expr); err != nil {
			log.Fatal(err)
		}
	} else if *csource {
		file := filepath.Join(*root, "fcall.c")
		if err = genCSource(file, expr); err != nil {
			log.Fatal(err)
		}
	} else if *hsource {
		file := filepath.Join(*root, "fcall.h")
		if err = genCHeader(file, expr); err != nil {
			log.Fatal(err)
		}
	} else if *jsonsource {
		data, _ := json.MarshalIndent(expr, "", "    ")
		//if file == "" {
		os.Stdout.Write(data)
		//	return
		//}
		//return ioutil.WriteFile(file, data, 0644)
	}
}

const (
	goTmpl = `package proto

import "fmt"

// 9P2000.L operations. There are 30 basic operations in 9P2000.L, paired
// as requests and responses. The one special case is ERROR as there is
// no NINEP_TERROR request for clients to transmit to the server, but the
// server may respond to any other request with an NINEP_RERROR.
//
// See also: http://9p.io/sys/man/5/INDEX.html
//
const ( {{range $k, $v := .}}
	_{{toUpper .Type}} uint8 = {{typeNum .Type}}{{end}}
)

func unmarshal(buf *Buffer, typ uint8, tag uint16, fcall Fcall) error {
	switch typ { {{range $k, $v := .}}
	case _{{toUpper .Type}}:
		f, ok := fcall.(*{{toPublic .Type}})
		if !ok {
			return fmt.Errorf("malformed fcall (num:%d) (type:%T)", typ, f)
		}
		f.tag = tag
		{{fromWireGo .}}	
		return buf.Err() {{end}}
	}
	return fmt.Errorf("unknown 9P2000.L fcall type (%d)", typ)
}

func marshal(buf *Buffer, tag uint16, fcall Fcall) error {
	switch f := fcall.(type) { {{range $k, $v := .}}
	case *{{toPublic .Type}}:
		buf.EncodeUint8(_{{toUpper .Type}})
		f.tag = tag
		buf.EncodeUint16(f.tag)
		{{toWireGo .}}	
		return buf.Err() {{end}}
	}
	return fmt.Errorf("unknown 9P2000.L fcall (%T)", fcall)
}

{{range $k, $v := .}}
type {{toPublic .Type}} struct {
	tag uint16 {{range $k1, $v1 := .Ident}}
	{{toPublic .Name}} {{toGo .Type}}{{end}}
}


//func (f {{toPublic .Type}}) FcallType() uint8 { return _{{toUpper .Type}} }
//func (f {{toPublic .Type}}) FcallMessage()    {}
//func (f {{toPublic .Type}}) Tag() uint16      { return f.tag }

{{end}}
`

	hTmpl = `#ifndef __FCALL_H
#define __FCALL_H

#include "binary.h"

/*
 * 9P2000.L operations. There are 30 basic operations in 9P2000.L, paired
 * as requests and responses. The one special case is ERROR as there is
 * no NINEP_TERROR request for clients to transmit to the server, but the
 * server may respond to any other request with an NINEP_RERROR.
 *
 * See also: http://9p.io/sys/man/5/INDEX.html
 */
enum { {{range $k, $v := .}}
	{{toUpper .Type}} = {{typeNum .Type}},{{end}}
};

/*
 * Qids are identifiers used by 9P2000.L servers to track file system
 * entities. The type is used to differentiate semantics for operations
 * on the entity. The path provides a server unique index for an entity
 * (roughly analogous to an inode number), while the version is updated
 * every time a file is modified and can be used to maintain cache
 * coherency between clients and serves.
 *
 * See also: http://9p.io/magic/man2html/2/stat
 */
typedef struct qid_t qid_t;
struct qid_t {
	uint8_t  type;
	uint32_t version;
	uint64_t path;
};

{{range $k, $v := .}}
typedef struct {{toLower .Type}}_t {{toLower .Type}}_t;{{end}}

typedef struct fcall_t fcall_t;

{{range $k, $v := .}}
struct {{toLower .Type}}_t { {{range $k1, $v1 := .Ident}}
	{{toC .Type}} {{.Name}};{{end}}
};
{{end}}

struct fcall_t {
	uint8_t  type;
	uint16_t tag;

	union { {{range $k, $v := .}}
		{{toLower .Type}}_t {{toLower .Type}};{{end}}
	};
};

ssize_t fcall_unmarshal(unsigned char*, size_t, fcall_t*);
ssize_t fcall_marshal(unsigned char*, size_t, fcall_t*);
ssize_t fcall_size(fcall_t*);

#endif /* __FCALL_H */
`

	cTmpl = `#include "fcall.h"

ssize_t
fcall_unmarshal(unsigned char* data, size_t msize, fcall_t* f)
{
	uint32_t      size = 0;
	unsigned char *p, *ep;

	p = data;
	ep = p + msize;
	if(p + 7 > ep)
		return -1;

	p = guint32(p, ep, &size);
	if(size < 7)
		return 0;
	p = guint8(p, ep, &f->type);
	p = guint16(p, ep, &f->tag);

	switch(f->type) { {{range $k, $v := .}}
	case {{toUpper .Type}}:
		{{fromWireC .}}
		break;{{end}}
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

	p = data;
	ep = p + msize;
	if(p + 7 > ep)
		return -1;

	p = puint32(p, ep, msize);
	p = puint8(p, ep, f->type);
	p = puint16(p, ep, f->tag);
	if(p == NULL)
		return -1;

	switch(f->type) { {{range $k, $v := .}}
	case {{toUpper .Type}}:
		{{toWireC .}}
		break;{{end}}
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
fcall_size(fcall_t* f) {
	uint32_t n = 4 + 1 + 2; // msize + type + tag

	switch(f->type) { {{range $k, $v := .}}
	case {{toUpper .Type}}:
		{{toSize .}}
		break;{{end}}
	default:
		return -1;
	}
	return n;
}
`
)

var gomain = "../internal/generator"

func genGoSource(file string, expr []parser.Expr) (err error) {
	b := &bytes.Buffer{}
	b.WriteString(fmt.Sprintf("//\n// GENERATED BY 'go run %s %s'; DO NOT EDIT!\n//\n\n",
		gomain, strings.Join(os.Args[1:], " ")))

	t := template.Must(template.New("go").Funcs(funcMap).Parse(goTmpl))
	if err = t.Execute(b, expr); err != nil {
		return err
	}

	data, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	if file == "" {
		os.Stdout.Write(data)
		return
	}
	return ioutil.WriteFile(file, data, 0644)
	//os.Stdout.Write(b.Bytes())
	//return nil
}

func genCHeader(file string, expr []parser.Expr) (err error) {
	b := &bytes.Buffer{}
	b.WriteString(fmt.Sprintf("/*\n * GENERATED BY 'go run %s %s'; DO NOT EDIT!\n */\n\n",
		gomain, strings.Join(os.Args[1:], " ")))

	t := template.Must(template.New("h").Funcs(funcMap).Parse(hTmpl))
	if err = t.Execute(b, expr); err != nil {
		log.Fatal(err)
	}

	data, err := cformat(b)
	if err != nil {
		log.Fatal(err)
	}
	if file == "" {
		os.Stdout.Write(data)
		return
	}
	return ioutil.WriteFile(file, data, 0644)
}

func genCSource(file string, expr []parser.Expr) (err error) {
	b := &bytes.Buffer{}
	b.WriteString(fmt.Sprintf("/*\n * GENERATED BY 'go run %s %s'; DO NOT EDIT!\n */\n\n",
		gomain, strings.Join(os.Args[1:], " ")))

	t := template.Must(template.New("c").Funcs(funcMap).Parse(cTmpl))
	if err = t.Execute(b, expr); err != nil {
		log.Fatal(err)
	}

	data, err := cformat(b)
	if err != nil {
		log.Fatal(err)
	}
	if file == "" {
		os.Stdout.Write(data)
		return
	}
	return ioutil.WriteFile(file, data, 0644)
}

func typeNum(t string) uint8 {
	return parser.TypeNum[t]
}

func toGo(t string) string {
	switch t {
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "string":
		return "string"
	case "[]byte":
		return "[]byte"
	case "qid":
		return "Qid"
	}
	return "unknown"
}

func toC(t string) string {
	switch t {
	case "uint8":
		return "uint8_t"
	case "uint16":
		return "uint16_t"
	case "uint32":
		return "uint32_t"
	case "uint64":
		return "uint64_t"
	case "string", "[]byte":
		return "char*"
	case "qid":
		return "qid_t"
	}
	return "unknown"
}

func toSize(expr parser.Expr) string {
	var s, size string
	for _, i := range expr.Ident {
		if i.IdentType == "basic" {
			switch i.Size {
			case "count":
				size = fmt.Sprintf("data_size(f->%s.%s)",
					strings.ToLower(expr.Type), i.Name)
			case "s":
				size = fmt.Sprintf("string_size(f->%s.%s)",
					strings.ToLower(expr.Type), i.Name)
			default:
				size = i.Size
			}
			s += size + "+"
		} else {

		}
	}
	if len(s) >= 1 {
		return fmt.Sprintf("n += %s;", s[:len(s)-1])
	} else {
		return "// nothing"
	}
}

func toWireGo(expr parser.Expr) string {
	var s string
	for _, i := range expr.Ident {
		if i.IdentType == "basic" {
			switch i.Type {
			case "uint8":
				s += fmt.Sprintf("buf.EncodeUint8(f.%s)\n", toPublic(i.Name))
			case "uint16":
				s += fmt.Sprintf("buf.EncodeUint16(f.%s)\n", toPublic(i.Name))
			case "uint32":
				s += fmt.Sprintf("buf.EncodeUint32(f.%s)\n", toPublic(i.Name))
			case "uint64":
				s += fmt.Sprintf("buf.EncodeUint64(f.%s)\n", toPublic(i.Name))
			case "string":
				s += fmt.Sprintf("buf.EncodeString(f.%s)\n", toPublic(i.Name))
			case "[]byte":
				s += fmt.Sprintf("buf.EncodeBytes(f.%s)\n", toPublic(i.Name))
			}
		} else {
			switch i.Type {
			case "string":
				//s += fmt.Sprintf("buf.EncodeStrings(f.%s)\n", toPublic(i.Name))
			}
		}
	}
	if len(s) >= 1 {
		s = s[:len(s)-1]
	}
	if len(s) == 0 {
		s = "// nothing"
	}
	return s
}

func fromWireGo(expr parser.Expr) string {
	var s string
	for _, i := range expr.Ident {
		if i.IdentType == "basic" {
			switch i.Type {
			case "uint8":
				s += fmt.Sprintf("buf.DecodeUint8(&f.%s)\n", toPublic(i.Name))
			case "uint16":
				s += fmt.Sprintf("buf.DecodeUint16(&f.%s)\n", toPublic(i.Name))
			case "uint32":
				s += fmt.Sprintf("buf.DecodeUint32(&f.%s)\n", toPublic(i.Name))
			case "uint64":
				s += fmt.Sprintf("buf.DecodeUint64(&f.%s)\n", toPublic(i.Name))
			case "string":
				s += fmt.Sprintf("buf.DecodeString(&f.%s)\n", toPublic(i.Name))
			case "[]byte":
				s += fmt.Sprintf("buf.DecodeBytes(&f.%s)\n", toPublic(i.Name))
			}
		} else {
			switch i.Type {
			case "string":
				//s += fmt.Sprintf("buf.DecodeStrings(&f.%s)\n", toPublic(i.Name))
			}
		}
	}
	if len(s) >= 1 {
		s = s[:len(s)-1]
	}
	if len(s) == 0 {
		s = "// nothing"
	}
	return s
}

func fromWireC(expr parser.Expr) string {
	var s string
	for _, i := range expr.Ident {
		if i.IdentType == "basic" {
			switch i.Type {
			case "uint8":
				s += fmt.Sprintf("p = guint8(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint16":
				s += fmt.Sprintf("p = guint16(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint32":
				s += fmt.Sprintf("p = guint32(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint64":
				s += fmt.Sprintf("p = guint64(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "string":
				s += fmt.Sprintf("p = gstring(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "qid":
				s += fmt.Sprintf("p = guint8(p, ep, &f->%s.%s.type);\n"+
					"p = guint32(p, ep, &f->%s.%s.version);\n"+
					"p = guint64(p, ep, &f->%s.%s.path);\n",
					strings.ToLower(expr.Type), i.Name,
					strings.ToLower(expr.Type), i.Name,
					strings.ToLower(expr.Type), i.Name)
			}
		} else {
			switch i.Type {
			case "string":
				s += fmt.Sprintf("//p = gstrings(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "[]byte":
				s += fmt.Sprintf("p = gdata(p, ep, &f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			}
		}
	}
	if len(s) >= 2 {
		s = s[:len(s)-2]
	}
	if len(s) > 0 {
		s += ";"
	} else {
		s = "// nothing"
	}
	return s
}

func toWireC(expr parser.Expr) string {
	var s string
	for _, i := range expr.Ident {
		if i.IdentType == "basic" {
			switch i.Type {
			case "uint8":
				s += fmt.Sprintf("p = puint8(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint16":
				s += fmt.Sprintf("p = puint16(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint32":
				s += fmt.Sprintf("p = puint32(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "uint64":
				s += fmt.Sprintf("p = puint64(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "string":
				s += fmt.Sprintf("p = pstring(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "[]byte":
				s += fmt.Sprintf("p = pdata(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "qid":
				s += fmt.Sprintf("p = puint8(p, ep, f->%s.%s.type);\n"+
					"p = puint32(p, ep, f->%s.%s.version);\n"+
					"p = puint64(p, ep, f->%s.%s.path);\n",
					strings.ToLower(expr.Type), i.Name,
					strings.ToLower(expr.Type), i.Name,
					strings.ToLower(expr.Type), i.Name)
			}
		} else {
			switch i.Type {
			case "string":
				s += fmt.Sprintf("//p = pstrings(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			case "[]byte":
				s += fmt.Sprintf("p = pdata(p, ep, f->%s.%s);\n",
					strings.ToLower(expr.Type), i.Name)
			}
		}
	}
	if len(s) >= 2 {
		s = s[:len(s)-2]
	}
	if len(s) > 0 {
		s += ";"
	} else {
		s = "// nothing"
	}
	return s
}

func upperInitial(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}
	return ""
}

func toPublic(s string) string {
	var toupper bool
	var r []rune

	for _, v := range s {
		if v == '_' {
			toupper = true
			continue
		}
		if toupper {
			v = unicode.ToUpper(v)
		}
		r = append(r, v)
		toupper = false
	}

	s = upperInitial(string(r))
	if s == "ClientId" { // hacky cleanup
		s = "ClientID"
	}
	return s
}

var funcMap = template.FuncMap{
	"toUpper":    strings.ToUpper,
	"toLower":    strings.ToLower,
	"toPublic":   toPublic,
	"typeNum":    typeNum,
	"toGo":       toGo,
	"toC":        toC,
	"toSize":     toSize,
	"toWireGo":   toWireGo,
	"fromWireGo": fromWireGo,
	"toWireC":    toWireC,
	"fromWireC":  fromWireC,
}

func cformat(buf *bytes.Buffer) ([]byte, error) {
	cmd := exec.Command("clang-format")
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	_, err = in.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	in.Close()
	formated, err := ioutil.ReadAll(out)
	if err != nil {
		return nil, err
	}
	if err = cmd.Wait(); err != nil {
		return nil, err
	}
	return formated, nil
}
