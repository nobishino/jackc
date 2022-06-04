package jackc_test

import (
	"strings"
	"testing"

	"github.com/nobishino/jackc"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	cases := []struct {
		src  string
		want int
	}{
		{`Class`, 1},
		{`method`, 1},
		{`}`, 1},
		{`{}`, 2},
		{`let temp = (xxx+12)*-63;`, 12},
		{`Class Bar {
	method Add3(int y) {
		return 3+y;
	}
}`, 17},
	}

	for _, tc := range cases {
		t.Run(tc.src, func(t *testing.T) {
			src := strings.NewReader(tc.src)
			tokenizer := jackc.NewTokenizer(src)
			var got int
			for {
				if !tokenizer.HasMoreTokens() {
					break
				}
				tokenizer.Advance()
				got++
			}
			assert.Equal(t, tc.want, got, "want %#v, got %#v", tc.want, got)
		})
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
		{
			src: `{`,
			want: `<tokens>
<symbol> { </symbol>
</tokens>
`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.src, func(t *testing.T) {
			src := strings.NewReader(tc.src)
			var buf strings.Builder
			err := jackc.TokenizeToXML(&buf, src)
			if err != nil {
				t.Fatal(err)
			}
			got := buf.String()
			assert.Equal(t, tc.want, got, "want %#v, got %#v", tc.want, got)
		})
	}
}
