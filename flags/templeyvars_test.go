package flags

import (
	"testing"
)

func TestSet(t *testing.T) {
	tv := make(TemplayVars)

	tv.Set("a=test")
	if tv["a"] != "test" {
		t.Errorf("Set a=test failed.\nGot: %v\nWanted: %v", tv["a"], 1)
	}

	tv.Set("b=1,c=2")
	if tv["b"] != "1" || tv["c"] != "2" {
		t.Errorf("Set b=1,c=2 failed.\nGot: %v,%v\nWanted: %v,%v", tv["a"], tv["b"], 1, 2)
	}
}

func TestType(t *testing.T) {
	tv := make(TemplayVars)
	if tv.Type() != "templayVar" {
		t.Errorf("Type failed.\nGot: %v\nWanted: %v", tv.Type(), "templayVar")
	}
}

func TestString(t *testing.T) {
	tv := make(TemplayVars)

	tv.Set("a=test1,b=test2,c=3")
	if tv.String() != "[a=test1,b=test2,c=3]" {
		t.Errorf("String failed.\nGot: %v\nWanted: %v", tv.String(), "[a=test1,b=test2,c=3]")
	}
}
