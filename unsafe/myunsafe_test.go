package unsafe

import (
	"testing"
)

var (
	strs = []string{
		"a",
		"abc",
		"some words",
		"loooooooooooooooger",
		"Characters with 1234567890 +-*/ and !@#$%^&*()=",
		`a multi-line long text here is line one.
line two.
line three.
some other texts:
1234567890-=!@#$%^&*()_+
abcdefghijklmnopqrstuvwxyz
ABCDEFGHIJKLMNOPQRSTUVWXYZ  `,
	}
)

func BenchmarkCast(b *testing.B) {
	n := len(strs)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			b := []byte(strs[j])
			_ = string(b)
		}
	}
}

func BenchmarkUnsafeCast(b *testing.B) {
	n := len(strs)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			b := StringToBytes(strs[j])
			_ = BytesToString(b)
		}
	}
}
