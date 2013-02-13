%{
// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz


package parser

import (
    "bufio"
    "log"
    "fmt"
)

type yylexer struct{
    src *bufio.Reader
    buf []byte
    empty bool
    current byte
    col int
}

func newLexer(src *bufio.Reader) (y *yylexer) {
    y = &yylexer{src: src}
    if b, err := src.ReadByte(); err == nil {
        y.current = b
    }
    return
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if b, err := y.src.ReadByte(); err == nil {
        y.current = b
    }
    return y.current
}

func (y yylexer) Error(e string) {
    log.Printf("Error in tokenizer\n")
    log.Printf("%s\n", y.buf)
    log.Printf("%s\n", y.current)
    log.Fatal(e)
}

func (y *yylexer) Lex(lval *yySymType) int {
    var err error
    _ = err
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

%%
    buf = buf[:0]

[ \t\r\n]+

[A-Z][a-zA-Z0-9_]*
    fmt.Printf("Lexed variable as\n%s\n", string(y.buf))
    lval.string = string(y.buf)
    return VARIABLE

[+\-*^/0-9a-z][+\-*^/0-9a-zA-Z]*
    lval.string = string(y.buf)
    return FUNCTION

";" { return SEMICOLON }
"=>" { return REWRITE }

[,] { return COMMA }
[\/\/].+ { }

"(" { return LEFT_PAREN }
")" { return RIGHT_PAREN }

%%
    y.empty = true
    return int(c)
}