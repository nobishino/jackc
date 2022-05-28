package jackc

import "io"

type Tokenizer struct {
	count int
}

// 入力ファイル/ストリームを開き、トークン化を行う準備をする
func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{}
}

// 入力にまだトークンは存在するか?
func (t *Tokenizer) HasMoreTokens() bool {
	return t.count == 0
}

// 入力から次のトークンを取得し、それを現在のトークン(現トークン)する。
// このルーチンは、hasMoreTokens()がTrueの場合のみ呼び出すことができる。
// また、最初は現トークンは設定されていない
func (t *Tokenizer) Advance() {
	t.count = t.count + 1
}

// 現トークンの種類を返す
func (t *Tokenizer) TokenType() TokenType {
	return KEYWORD
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
