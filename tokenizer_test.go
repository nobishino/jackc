package jackc_test

import (
	"strings"
	"testing"

	"github.com/nobishino/jackc"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	type result = []jackc.Token
	cases := []struct {
		src  string
		want []jackc.Token
	}{
		{
			src: `class`,
			want: result{
				{jackc.KEYWORD, "CLASS"},
			},
		},
	}

	for _, tc := range cases {
		src := strings.NewReader(tc.src)
		tokenizer := jackc.NewTokenizer(src)
		tokenizer.Advance()
		var got []jackc.Token
		for tokenizer.HasMoreTokens() {
			var tok jackc.Token
			tok.TokenType = tokenizer.TokenType()
			tok.Value = string(tokenizer.KeyWord())
			got = append(got, tok)
			tokenizer.Advance()
		}

		assert.Equal(t, tc.want, got, "want %#v, got %#v", tc.want, got)
	}
}

func TestTokenizeToXML(t *testing.T) {
	cases := []struct {
		src  string
		want string
	}{
		{
			src: `class`,
			want: `<tokens>
<keyword> class </keyword>
</tokens>
`,
		},
	}
	for _, tc := range cases {
		src := strings.NewReader(tc.src)
		var buf strings.Builder
		err := jackc.TokenizeToXML(&buf, src)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		assert.Equal(t, tc.want, got, "want %#v, got %#v", tc.want, got)
	}
}
