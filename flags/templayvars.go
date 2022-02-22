package flags

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/Korazza/templay/utils"
)

type TemplayVars map[string]interface{}

// Format: a=1,b=2
func (tv *TemplayVars) Set(v string) error {
	var ss []string
	n := strings.Count(v, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as key=value", v)
	case 1:
		ss = append(ss, strings.Trim(v, `"`))
	default:
		r := csv.NewReader(strings.NewReader(v))
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

func (tv *TemplayVars) Load(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return fmt.Errorf("failed to open file %s", fname)
	}
	defer f.Close()

	err = utils.ParseYaml(f, tv)
	if err != nil {
		return fmt.Errorf("failed to parse yaml file %s", fname)
	}

	return nil
}
