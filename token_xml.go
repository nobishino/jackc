package jackc

import (
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func TokenizeToXML(w io.Writer, r io.Reader) error {
	tokenizer := NewTokenizer(r)
	const startTag = "<tokens>"
	const endTag = "</tokens>"

	if _, err := fmt.Fprintln(w, startTag); err != nil {
		return errors.WithStack(err)
	}

	for {
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
		default:
			return errors.Errorf("unexpected type: %q", tp)
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return errors.WithStack(err)
		}
	}
	if _, err := fmt.Fprintln(w, endTag); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
