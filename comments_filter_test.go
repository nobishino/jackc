package jackc_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobishino/jackc"
	"github.com/stretchr/testify/assert"
)

func TestFilterComments(t *testing.T) {
	const testDir = "testdata/comments_filter"
	testcases := []string{
		"Main.jack",
	}
	for _, tc := range testcases {
		t.Run(tc, func(t *testing.T) {
			src := openFile(t, filepath.Join(testDir, tc))
			want := readFile(t, filepath.Join(testDir, tc+".filtered"))
			got, err := jackc.FilterComments(src)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, want, got.String())
		})
	}
}

func openFile(t *testing.T, path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	return f
}
func readFile(t *testing.T, path string) string {
	f := openFile(t, path)
	var sb strings.Builder
	_, err := io.Copy(&sb, f)
	if err != nil {
		t.Fatal(err)
	}
	return sb.String()
}
