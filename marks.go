package pbar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Marks is a combination of multiple Marks
type Marks struct {
	Slice  []Marker
	Inputs []*MarkInput
}

func (m *Marks) InputNames() []string {
	r := make([]string, 0, len(m.Inputs))
	for _, input := range m.Inputs {
		r = append(r, input.Name)
	}
	return r
}

func (m *Marks) Input(name string, val string) {
	for _, input := range m.Inputs {
		if input.Name == name {
			input.Val = json.RawMessage(val)
		}
	}
}

// MarkFormat returns mark string
func (m *Marks) MarkFormat(info *Info) string {
	ss := []string{}
	for _, v := range m.Slice {
		ss = append(ss, v.MarkFormat(info))
	}
	return strings.Join(ss, "")
}

var markMap = map[string]reflect.Type{}

func init() {
	for _, v := range []Marker{&MarkBar{}, &MarkInput{}, &MarkAfter{}, &MarkPer{}, &MarkPerCent{}, &MarkRoll{}, &MarkText{}} {
		mapping(v)
	}
}

func mapping(v Marker) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	key := reflect.Indirect(val).Type().Name()
	typ := val.Type()
	key = strings.TrimPrefix(key, "Mark")
	markMap[key] = typ
}

func (m *Marks) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Slice)
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
	for i := 0; i != len(rows); i++ {
		row := rows[i]
		row = bytes.TrimSpace(row)
		if row[0] == '[' {
			v := []json.RawMessage{}
			err := json.Unmarshal(row, &v)
			if err != nil {
				return err
			}
			rows = append(rows[:i], append(v, rows[1+1:]...)...)
		} else if row[0] == '"' {
			m.Slice = append(m.Slice, &MarkInput{
				Val: json.RawMessage(row),
			})
		} else {
			var kind struct {
				Kind string `json:"_Kind"`
			}

			err := json.Unmarshal(row, &kind)
			if err != nil {
				return err
			}
			typ, ok := markMap[kind.Kind]
			if !ok {
				return fmt.Errorf("Error: kind `%s` not found", kind.Kind)
			}
			val := reflect.New(typ)

			vp := val.Interface()
			err = json.Unmarshal(row, vp)
			if err != nil {
				return err
			}

			v := val.Elem().Interface()
			if mark, ok := v.(Marker); ok {
				m.Slice = append(m.Slice, mark)
				m.Inputs = mergeMarkInput(m.Inputs, mark)
			}
		}
	}
	return nil
}

var typeMarkInput1 = reflect.TypeOf((*MarkInput)(nil))
var typeMarkInput2 = reflect.TypeOf((*Marks)(nil))

func mergeMarkInput(result []*MarkInput, v interface{}) []*MarkInput {
	val := reflect.ValueOf(v)
	val = reflect.Indirect(val)

	if val.Kind() != reflect.Struct {
		return nil
	}
	length := val.NumField()
	for i := 0; i != length; i++ {
		field := val.Field(i)
		if field.Type() == typeMarkInput1 {
			v := field.Interface().(*MarkInput)
			if v != nil && v.Name != "" {

				index := sort.Search(len(result), func(i int) bool {
					return result[i].Name >= v.Name
				})

				if index != len(result) && result[index].Name == v.Name {
					val := reflect.ValueOf(result[index])
					field.Set(val)
				} else {
					result = append(result, nil)
					copy(result[i+1:], result[i:])
					result[i] = v
				}
			}
		} else if field.Type() == typeMarkInput2 {
			v := field.Interface().(*Marks)
			for _, input := range v.Inputs {
				result = mergeMarkInput(result, input)
			}
		}
	}

	return result
}
