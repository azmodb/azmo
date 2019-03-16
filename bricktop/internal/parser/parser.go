package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Ident struct {
	IdentType string
	LenType   string

	Type string
	Name string
	Size string
}

type Expr struct {
	Type  string
	Ident []Ident
}

func Parse(r io.Reader) (expr []Expr, err error) {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadBytes('\n')
		if err != nil {
			break
		}
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		s, e := NewScanner(line), Expr{}
		e.Type, err = parseType(s)
		if err != nil {
			return expr, err
		}
		e.Ident, err = parseIdent(s)
		if err != nil && err != io.EOF {
			return expr, err
		}
		expr = append(expr, e)
	}
	return expr, err
}

func parseType(s *Scanner) (typ string, err error) {
	tok, lit := s.Scan()
	if tok == EOF {
		return "", io.ErrUnexpectedEOF
	}
	if tok != IDENT {
		return "", fmt.Errorf("malformed input")
	}
	return string(lit), nil
}

func parseIdent(s *Scanner) (ident []Ident, err error) {
	for {
		tok, lit := s.Scan()
		if tok == EOF {
			return ident, io.EOF
		}
		if string(lit) == "tag" {
			tok, lit = s.Scan() // TODO
			tok, lit = s.Scan() // TODO
			tok, lit = s.Scan() // TODO
			continue
		}

		i := Ident{}
		switch tok {
		case IDENT:
			i.IdentType = "basic"
			i.Name = string(lit)
			tok, lit = s.Scan() // TODO
			tok, lit = s.Scan()
			i.Type = toType(lit)
			i.Size = string(lit)
			tok, lit = s.Scan() // TODO

		case LBRACK:
			i.IdentType = "array"
			tok, lit = s.Scan()
			i.LenType = toType(lit)
			tok, lit = s.Scan() // TODO

			tok, lit = s.Scan()
			i.Name = string(lit)
			tok, lit = s.Scan() // TODO
			tok, lit = s.Scan()
			i.Type = toType(lit)
			i.Size = string(lit)
			tok, lit = s.Scan() // TODO
		}
		ident = append(ident, i)
	}
	return ident, err
}

func toType(lit []byte) string {
	switch string(lit) {
	case "1":
		return "uint8"
	case "2":
		return "uint16"
	case "4":
		return "uint32"
	case "8":
		return "uint64"
	case "s":
		return "string"
	case "13":
		return "qid"
	case "count":
		return "[]byte"
	default:
		return "unknown"
	}
}
