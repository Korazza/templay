package flags

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type TemplayVars map[string]interface{}

// Format: a=1,b=2
func (tv *TemplayVars) Set(val string) error {
	var ss []string
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as key=value", val)
	case 1:
		ss = append(ss, strings.Trim(val, `"`))
	default:
		r := csv.NewReader(strings.NewReader(val))
		var err error
		ss, err = r.Read()
		if err != nil {
			return err
		}
	}

	out := make(map[string]interface{}, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}
		out[kv[0]] = kv[1]
	}

	if *tv == nil {
		*tv = out
	} else {
		for k, v := range out {
			(*tv)[k] = v
		}
	}

	return nil
}

func (tv *TemplayVars) Type() string {
	return "templayVar"
}

func (tv *TemplayVars) String() string {
	records := make([]string, 0, len(*tv)>>1)
	for k, v := range *tv {
		records = append(records, k+"="+fmt.Sprintf("%v", v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return "[" + strings.TrimSpace(buf.String()) + "]"
}

func (tv *TemplayVars) Load(file string) bool {
	templayvarsYAML, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}
	err = yaml.Unmarshal(templayvarsYAML, tv)
	return err == nil
}
