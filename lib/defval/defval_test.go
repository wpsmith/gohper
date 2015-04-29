package defval

import (
	"testing"

	"github.com/cosiner/gohper/lib/test"
)

func TestInt(t *testing.T) {
	tt := test.Wrap(t)
	val := 0
	Int(&val, 10)
	tt.Eq(val, 10)
}

func TestString(t *testing.T) {
	tt := test.Wrap(t)
	val := ""
	String(&val, "abc")
	tt.Eq(val, "abc")
}
