package jackc_test

import (
	"strings"
	"testing"

	"github.com/nobishino/jackc"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	type result = []jackc.TokenType
	cases := []struct {
		src  string
		want []jackc.TokenType
	}{
		{
			src:  `class`,
			want: result{jackc.KEYWORD},
		},
	}

	for _, tc := range cases {
		src := strings.NewReader(tc.src)
		tk := jackc.NewTokenizer(src)
		var got []jackc.TokenType
		for tk.HasMoreTokens() {
			got = append(got, tk.TokenType())
			tk.Advance()
		}

		assert.Equal(t, tc.want, got)
	}

}
