package parser

import (
	"unicode"
	"unicode/utf8"
)

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' ||
		ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' ||
		ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

const bom = 0xFEFF // byte order mark, only permitted as very first char

type Scanner struct {
	rdOffset int // reading offset (position after current character)
	offset   int // character offset
	src      []byte
	ch       rune // current character
}

func NewScanner(src []byte) *Scanner {
	s := &Scanner{src: src}
	s.next()
	return s
}

type Token int

const (
	ILLEGAL Token = iota
	EOF

	LBRACK // [
	RBRACK // ]

	IDENT
)

func (t Token) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case LBRACK:
		return "LBRACK"
	case RBRACK:
		return "RBRACK"
	case IDENT:
		return "IDENT"
	default:
		return "UNKNOWN"
	}
}

// Read the next Unicode char into s.ch. s.ch < 0 means end-of-file.
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		//if s.ch == '\n' {
		//	s.lineOffset = s.offset
		//	s.file.AddLine(s.offset)
		//}
		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		case r == 0:
			//s.error(s.offset, "illegal character NUL")
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				//s.error(s.offset, "illegal UTF-8 encoding")
			} else if r == bom && s.offset > 0 {
				//s.error(s.offset, "illegal byte order mark")
			}
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		//if s.ch == '\n' {
		//	s.lineOffset = s.offset
		//	s.file.AddLine(s.offset)
		//}
		s.ch = -1 // eof
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() []byte {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return s.src[offs:s.offset]
}

func (s *Scanner) Scan() (tok Token, lit []byte) {
	s.skipWhitespace()
	switch ch := s.ch; {
	case isLetter(ch):
		tok = IDENT
		lit = s.scanIdentifier()
	case isDigit(ch):
		tok = IDENT
		lit = s.scanIdentifier()
	default:
		s.next() // always make progress
		switch ch {
		case -1:
			tok = EOF
		case '[':
			tok = LBRACK
		case ']':
			tok = RBRACK
		}
	}
	return tok, lit
}
