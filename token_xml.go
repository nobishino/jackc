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

	tokenizer.Advance()
	for tokenizer.HasMoreTokens() {
		var line string
		switch tp := tokenizer.TokenType(); tp {
		case KEYWORD:
			kw := strings.ToLower(string(tokenizer.KeyWord()))
			line = fmt.Sprintf("<keyword> %s </keyword>", kw)
		default:
			return errors.Errorf("unexpected type: %q", tp)
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return errors.WithStack(err)
		}
		tokenizer.Advance()
	}
	if _, err := fmt.Fprintln(w, endTag); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
