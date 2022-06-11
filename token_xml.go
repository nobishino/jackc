package jackc

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/pkg/errors"
)

func TokenizeToXML(w io.Writer, r io.Reader) error {
	tokenizer, err := NewTokenizer(r)
	if err != nil {
		return err
	}
	const startTag = "<tokens>"
	const endTag = "</tokens>"

	if _, err := fmt.Fprintln(w, startTag); err != nil {
		return errors.WithStack(err)
	}

	for {
		log.Println(tokenizer.HasMoreTokens())
		if !tokenizer.HasMoreTokens() {
			break
		}
		tokenizer.Advance()
		var line string
		switch tp := tokenizer.TokenType(); tp {
		case KEYWORD:
			kw := strings.ToLower(string(tokenizer.KeyWord()))
			line = fmt.Sprintf("<keyword> %s </keyword>", kw)
		case SYMBOL:
			symbol := string(tokenizer.Symbol())
			line = fmt.Sprintf("<symbol> %s </symbol>", symbol)
		case IDENTIFIER:
			ident := tokenizer.Identifier()
			line = fmt.Sprintf("<identifier> %s </identifier>", ident)
		default:
			return errors.Errorf("unexpected type: %q", tp)
		}
		log.Println(line)
		if _, err := fmt.Fprintln(w, line); err != nil {
			return errors.WithStack(err)
		}
	}
	if _, err := fmt.Fprintln(w, endTag); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
