package utils

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseYaml(t *testing.T) {
	v := make(map[string]interface{})

	var b bytes.Buffer
	b.WriteString(`a: test1
b: test2
c: 3`)
	err := ParseYaml(&b, v)
	if err != nil {
		t.Errorf("ParseYaml failed.\n%v", err)
	}
	if v["a"] != "test1" {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", v["a"], "test1")
	}
	if v["b"] != "test2" {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", v["b"], "test2")
	}
	if v["c"] != 3 {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", v["c"], 3)
	}

	v = make(map[string]interface{})
	b.Reset()
	b.WriteString(`a: test1
b: test2
c:
  d: 3`)
	err = ParseYaml(&b, v)
	if err != nil {
		t.Errorf("ParseYaml failed.\n%v", err)
	}
	if v["a"] != "test1" {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", v["a"], "test1")
	}
	if v["b"] != "test2" {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", v["b"], "test2")
	}
	c := reflect.ValueOf(v["c"])
	if c.Kind() != reflect.Map {
		t.Errorf("ParseYaml failed.\n\"c\" is not a map")
	}
	d := c.MapIndex(c.MapKeys()[0])
	if d.Interface() != 3 {
		t.Errorf("ParseYaml failed.\nGot: %v\nWanted: %v", d.Interface(), 3)
	}
}
