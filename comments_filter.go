package jackc

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/pkg/errors"
)

// FilterComments は、rのソースコードから//...コメントと/*...*/コメントを取り除く
func FilterComments(r io.Reader) (*bytes.Buffer, error) {
	var result bytes.Buffer
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line, _, _ = strings.Cut(line, "//")
		line = strings.TrimRight(line, "\r\t ")
		if _, err := result.WriteString(line + "\n"); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	// return &result, nil
	result2, err := filterLongComments(&result)
	if err != nil {
		return nil, err
	}
	return result2, nil
}

func filterLongComments(rd io.Reader) (*bytes.Buffer, error) {
	var result bytes.Buffer
	br := bufio.NewReader(rd)
	var inComment bool
	var r [4]rune
	begin := func() bool {
		return r[0] == '*' && r[1] == '/'
	}
	end := func() bool {
		return r[2] == '/' && r[3] == '*'
	}
	shouldWrite := func() bool {
		return !inComment && r[2] != 0
	}
	advance := func(c rune) error {
		r[0], r[1], r[2], r[3] = c, r[0], r[1], r[2]
		log.Println(inComment, string(r[2]))
		if shouldWrite() {
			_, err := result.WriteRune(r[2])
			if err != nil {
				return errors.WithStack(err)
			}
		}
		inComment = inComment && !end()
		inComment = inComment || begin()
		return nil
	}
	for {
		c, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if err := advance(c); err != nil {
			return nil, err
		}
	}
	for i := 0; i < 2; i++ {
		if err := advance(0); err != nil {
			return nil, err
		}
	}
	return &result, nil
}
