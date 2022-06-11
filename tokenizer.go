package jackc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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
func NewTokenizer(r io.Reader) (*Tokenizer, error) {
	b, err := FilterComments(r)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(b)
	return &Tokenizer{
		reader:        buf,
		hasMoreTokens: true,
	}, nil
}

// 入力にまだトークンは存在するか?
// tの状態を変えない
func (t *Tokenizer) HasMoreTokens() bool {
	return t.hasMoreTokens
}

// 入力から次のトークンを取得し、それを現在のトークン(現トークン)する。
// このルーチンは、hasMoreTokens()がTrueの場合のみ呼び出すことができる。
// また、最初は現トークンは設定されていない
//
// 呼び出し時の条件: 次トークンの先頭文字をload済みであるか、まだ1文字もloadしていない
//
// 完了時の条件: そのさらに次トークンの先頭の先頭文字をload済みであるか、もしくはt.eof == trueである
func (t *Tokenizer) Advance() {
	if !t.HasMoreTokens() {
		panic("Advance must not be called if !HasMoreTokens")
	}
	defer func() {
		// Advanceが終わった時点で次のtokenの先頭かもしくはt.eof == true
		t.skipDelimiters()
		if t.eof {
			t.hasMoreTokens = false
		}
	}()
	if t.currentRune == 0 { // 1文字もloadしていない状態に対応する
		t.advanceRune()
	}
	r := t.currentRune // rは次トークンの先頭文字になっている

	// TokenTypeを分類する
	// symbol
	if t.isSymbol(r) {
		t.currentToken = Token{
			TokenType: SYMBOL,
			Value:     string(r),
		}
		if !t.eof {
			t.advanceRune()
		}
		return
	}
	tk := t.readWord()
	switch {
	// intergerConstant
	// stringConstant
	// keyword
	case isKeyword(tk):
		t.currentToken = Token{
			TokenType: KEYWORD,
			Value:     tk,
		}
	case isIdentifier(tk):
		// identifiers
		t.currentToken = Token{
			TokenType: IDENTIFIER,
			Value:     tk,
		}
	default:
		panic(fmt.Sprintf("current token value = %q", tk))
	}
}

func isIdentifier(s string) bool {
	return len(s) > 0
}

func isKeyword(s string) bool {
	switch s {
	case "class", "constructor", "function", "method", "field", "static",
		"var", "int", "char", "boolean", "void", "true",
		"false", "null", "this", "let", "do", "if", "else", "while", "return":
		return true
	}
	return false
}

// should not be called if t.eof
func (t *Tokenizer) advanceRune() {
	if t.eof {
		panic("advanceRune invoked at EOF")
	}
	r, _, err := t.reader.ReadRune()
	if err == io.EOF {
		t.eof = true
		t.hasMoreTokens = false
		t.currentRune = 0
		return
	}
	if err != nil {
		panic(err)
	}
	t.currentRune = r
}

// readWord reads current word
// at the end of this, t.isDelimiters() == true or t.isSymbol(t.currentRune)
func (t *Tokenizer) readWord() string {
	var result []rune
	for !t.eof && !t.isDelimiters() && !t.isSymbol(t.currentRune) {
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
	return Keyword(strings.ToUpper(t.currentToken.Value))
}

// 現トークンの文字を返す。このルーチンは、tokenType()がSYMBOLの場合に呼び出すことができる
func (t *Tokenizer) Symbol() Symbol {
	return Symbol([]rune(t.currentToken.Value)[0])
}

// Identifierは、現トークンの識別子を返す。tokenType()がÍIDENTIFIERの場合のみ呼び出せる
func (t *Tokenizer) Identifier() string {
	if tp := t.currentToken.TokenType; tp != IDENTIFIER {
		panic(fmt.Sprintf("current token type should be %q, but got %q", IDENTIFIER, tp))
	}
	return t.currentToken.Value
}

// Intval は、現トークンの整数の値を返す。tokenType() == INT_VALのときのみ呼び出せる。
func (t *Tokenizer) IntVal() int {
	if tp := t.currentToken.TokenType; tp != INT_CONST {
		panic(fmt.Sprintf("current token type should be %q, but got %q", INT_CONST, tp))
	}
	v, err := strconv.Atoi(t.currentToken.Value)
	if err != nil {
		panic(err)
	}
	return v
}

// StringVal は、現トークンの文字列値をかえす。tokenType() == STRING_VALのときのみ呼び出せる。
func (t *Tokenizer) StringVal() string {
	if tp := t.currentToken.TokenType; tp != STRING_CONST {
		panic(fmt.Sprintf("current token type should be %q, but got %q", STRING_CONST, tp))
	}
	return t.currentToken.Value
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
