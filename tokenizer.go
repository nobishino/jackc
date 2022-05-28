package jackc

import (
	"bufio"
	"io"
	"strings"
)

type Token struct {
	TokenType
	Value string
}

type Tokenizer struct {
	reader        *bufio.Reader
	hasMoreTokens bool
	currentToken  Token
	currentRune   rune
	eof           bool
}

// 入力ファイル/ストリームを開き、トークン化を行う準備をする
func NewTokenizer(r io.Reader) *Tokenizer {
	buf := bufio.NewReader(r)
	return &Tokenizer{
		reader:        buf,
		hasMoreTokens: true,
	}
}

// 入力にまだトークンは存在するか?
func (t *Tokenizer) HasMoreTokens() bool {
	return t.hasMoreTokens
}

// 入力から次のトークンを取得し、それを現在のトークン(現トークン)する。
// このルーチンは、hasMoreTokens()がTrueの場合のみ呼び出すことができる。
// また、最初は現トークンは設定されていない
func (t *Tokenizer) Advance() {
	if !t.HasMoreTokens() {
		panic("Advance must not be called if !HasMoreTokens")
	}
	if t.eof {
		t.hasMoreTokens = false
		return
	}
	t.skipDelimiters()
	r := t.currentRune
	// symbol
	if t.isSymbol(r) {
		t.currentToken = Token{
			TokenType: SYMBOL,
			Value:     string(r),
		}
		return
	}
	tk := t.readWord()
	// intergerConstant
	// stringConstant
	// keyword
	t.currentToken = Token{
		TokenType: KEYWORD,
		Value:     tk,
	}
}

// should not be called if !t.eof
func (t *Tokenizer) advanceRune() {
	if t.eof {
		panic("advanceRune invoked at EOF")
	}
	r, _, err := t.reader.ReadRune()
	if err == io.EOF {
		t.eof = true
		return
	}
	if err != nil {
		panic(err)
	}
	t.currentRune = r
}

// readWord reads current word
// at the end of this, t.isDelimiters() == true
func (t *Tokenizer) readWord() string {
	var result []rune
	for !t.eof && !t.isDelimiters() {
		result = append(result, t.currentRune)
		t.advanceRune()
	}
	return string(result)
}

func (t *Tokenizer) skipDelimiters() {
	for !t.eof && t.isDelimiters() {
		t.advanceRune()
	}
}

func (t *Tokenizer) isDelimiters() bool {
	return strings.ContainsRune(" \n\t\r\x00", t.currentRune)
}

func (t *Tokenizer) isSymbol(r rune) bool {
	return strings.ContainsRune("{}()[].,;+-*/&|<>=~", r)
}

// 現トークンの種類を返す
func (t *Tokenizer) TokenType() TokenType {
	return t.currentToken.TokenType
}

// 現トークンのキーワードを返す。このルーチンは、tokenType()がKEYWORDの場合のみ呼び出すことができる
func (t *Tokenizer) KeyWord() Keyword {
	var v Keyword
	return v
}

// 現トークンの文字を返す。このルーチンは、tokenType()がSYMBOLの場合に呼び出すことができる
func (t *Tokenizer) Symbol() Symbol {
	var s Symbol
	return s
}

// Identifierは、現トークンの識別子を返す。tokenType()がÍIDENTIFIERの場合のみ呼び出せる
func (t *Tokenizer) Identifier() string {
	return ""
}

// Intval は、現トークンの整数の値を返す。tokenType() == INT_VALのときのみ呼び出せる。
func (t *Tokenizer) IntVal() int {
	return 0
}

// StringVal は、現トークンの文字列値をかえす。tokenType() == STRING_VALのときのみ呼び出せる。
func (t *Tokenizer) StringVal() string {
	return ""
}

// TokenType represents token types
type TokenType string

const (
	KEYWORD      TokenType = "KEYWORD"
	SYMBOL       TokenType = "SYMBOL"
	IDENTIFIER   TokenType = "IDENTIFIER"
	INT_CONST    TokenType = "INT_CONST"
	STRING_CONST TokenType = "STRING_CONST"
)

// Keyword represents keywords
type Keyword string

const (
	CLASS       Keyword = "CLASS"
	METHOD      Keyword = "METHOD"
	FUNCTION    Keyword = "FUNCTION"
	CONSTRUCTOR Keyword = "CONSTRUCTOR"
	INT         Keyword = "INT"
	BOOLEAN     Keyword = "BOOLEAN"
	CHAR        Keyword = "CHAR"
	VOID        Keyword = "VOID"
	VAR         Keyword = "VAR"
	STATIC      Keyword = "STATIC"
	FIELD       Keyword = "FIELD"
	LET         Keyword = "LET"
	DO          Keyword = "DO"
	IF          Keyword = "IF"
	ELSE        Keyword = "ELSE"
	WHILE       Keyword = "WHILE"
	RETURN      Keyword = "RETURN"
	TRUE        Keyword = "TRUE"
	FALSE       Keyword = "FALSE"
	NULL        Keyword = "NULL"
	THIS        Keyword = "THIS"
)

// Symbol represents symbols
type Symbol rune
