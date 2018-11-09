package pbar

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

// Marks is a combination of multiple Marks
type Marks []Mark

// MarkFormat returns mark string
func (m *Marks) MarkFormat(info *Info) string {
	ss := []string{}
	for _, v := range *m {
		ss = append(ss, v.MarkFormat(info))
	}
	return strings.Join(ss, "")
}

var markMap = map[string]reflect.Type{}

func init() {
	for _, v := range []Mark{&MarkBar{}, &MarkAfter{}, &MarkPercent{}, &MarkRatio{}, &MarkRoll{}, &MarkText{}} {
		mapping(v)
	}
}
func mapping(v Mark) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	key := reflect.Indirect(val).Type().Name()
	typ := val.Type()
	key = strings.TrimPrefix(key, "Mark")
	markMap[key] = typ
}

func (m *Marks) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	rows := []json.RawMessage{}
	if data[0] == '[' {
		err := json.Unmarshal(data, &rows)
		if err != nil {
			return err
		}
	}
	for _, row := range rows {
		row = bytes.TrimSpace(row)
		if row[0] == '[' {
			r := &Marks{}
			err := r.UnmarshalJSON(row)
			if err != nil {
				return err
			}
			*m = append(*m, r)
		} else {
			var kind struct {
				Kind string `json:"_Kind"`
			}

			err := json.Unmarshal(row, &kind)
			if err != nil {
				return err
			}
			typ := markMap[kind.Kind]

			val := reflect.New(typ)
			inte := val.Interface()

			err = json.Unmarshal(row, inte)
			if err != nil {
				return err
			}

			v := val.Elem().Interface().(Mark)
			*m = append(*m, v)
		}
	}
	return nil
}
