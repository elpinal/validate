package emails

import (
	"strings"
	"testing"
)

var valids = []struct {
	input string
}{
	{"aaa@example.com"},
	{"aaa+bbb@example.com"},
	{"AAABBB@example.com"},
	{"aaa.bbb.ccc@example.com"},
	{strings.Repeat("a", 64) + "@" + strings.Repeat("b", 254-69) + ".com"}, // TODO: avoid using magic numbers
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
	{""},
	{"@"},
	{".@.com"},
	{"aaa.@example.com"},
	{"aaa bbb@example.com"},
	{"aaa..bbb..ccc@example.com"},
	{strings.Repeat("a", 65) + "@" + strings.Repeat("b", 254-70) + ".com"},
	{strings.Repeat("a", 255-len("@example.com")) + "@example.com"},
}

func TestValidateInvalids(t *testing.T) {
	for _, v := range invalids {
		if err := Validate(v.input); err == nil {
			t.Errorf("Validate(%s) succeeded unexpectedly", v.input)
		}
	}
}
