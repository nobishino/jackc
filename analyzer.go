package jackc

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// source: Xxx.jack / directory name
// output: Xxx.xml
// create JackTokenizer
// use Compilation Engine

type Analyzer struct{}

// ExecTokenize
func (a Analyzer) ExecTokenize(path string) error {
	// pathがhoge.jackだったらそれをhogeT.xmlに吐き出す
	info, err := os.Stat(path)
	if err != nil {
		return errors.WithStack(err)
	}
	switch {
	case info.IsDir():
		return nil // TODO
	default:
		return a.TokenizeFile(path)
	}
}

// pathがhoge.jackだったらhogeT.xmlに結果を書き出す
func (a Analyzer) TokenizeFile(path string) error {
	log.Println("TokenizeFile:", path)
	if filepath.Ext(path) != ".jack" {
		return errors.Errorf("source file must have ext .jack. got path: %q", path)
	}
	src, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer src.Close()

	t, err := NewTokenizer(src)
	if err != nil {
		return err
	}

	for t.HasMoreTokens() {
		switch typ := t.TokenType(); typ {
		case KEYWORD:
			log.Println(typ)
		case SYMBOL:
			log.Println(typ)
		case IDENTIFIER:
			log.Println(typ)
		case INT_CONST:
			log.Println(typ)
		case STRING_CONST:
			log.Println(typ)
		default:
			return errors.Errorf("unexpected token type: %q", typ)
		}
	}
	// ここでtの機能を使っていろいろやる

	return nil
}

func (a Analyzer) ExecParse(path string) error {
	// pathがhoge.jackだったらそれをhoge.xml
	return nil
}
