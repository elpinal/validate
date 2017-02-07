package passwords

import (
	"strings"
	"testing"
)

var valids = []struct {
	input string
}{
	{"123456abcdef"},
	{"!@#$%^&*("},
	{")_+{}|:/?>"},
	{"AAABBBCCC"},
	{strings.Repeat("a", 128)},
}

func TestValidateValids(t *testing.T) {
	for _, v := range valids {
		if err := Validate(v.input); err != nil {
			t.Errorf("Validate(%s): %v", v.input, err)
		}
	}
}

var invalids = []struct {
	input string
}{
	{"123456"},
	{"abcdef 123456"},
	{"hello世界"},
	{"🔑aaaaaaaa"},
	{string([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})},
	{strings.Repeat("a", 129)},
}

func TestValidateInvalids(t *testing.T) {
	for _, v := range invalids {
		if err := Validate(v.input); err == nil {
			t.Errorf("Validate(%s) succeeded unexpectedly", v.input)
		}
	}
}

var benchLong = strings.Repeat("a", 128)

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validate(benchLong)
	}
}
