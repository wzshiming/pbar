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
			input.Input(val)
		}
	}
}

// MarkFormat returns mark string
func (m *Marks) String() string {
	buf := bytes.NewBuffer(nil)
	for _, v := range m.Slice {
		b := v.String()
		b = strings.Replace(b, "\\u001b", "\x1b", -1)
		buf.WriteString(b)
	}
	return buf.String()
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
	} else {
		rows = append(rows, data)
	}
	for i := 0; i != len(rows); i++ {
		row := rows[i]
		err := m.unmarshalJSON(row)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Marks) unmarshalJSON(row []byte) error {
	row = bytes.TrimSpace(row)

	switch row[0] {
	case '[':
		v := []json.RawMessage{}
		err := json.Unmarshal(row, &v)
		if err != nil {
			return err
		}
		for _, v := range v {
			err = m.unmarshalJSON(v)
			if err != nil {
				return err
			}
		}
	case '{':
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

			if input, ok := mark.(*MarkInput); ok {
				m.Inputs, input = setMarkInput(m.Inputs, input)
				m.Slice = append(m.Slice, input)
			} else {
				m.Slice = append(m.Slice, mark)
				m.Inputs = mergeMarkInput(m.Inputs, mark)
			}
		}
	default:
		m.Slice = append(m.Slice, &MarkTab{json.RawMessage(row)})
	}
	return nil
}

func setMarkInput(set []*MarkInput, inset *MarkInput) ([]*MarkInput, *MarkInput) {
	index := sort.Search(len(set), func(i int) bool {
		return set[i].Name >= inset.Name
	})

	if index == len(set) {
		return append(set, inset), inset
	} else if set[index].Name == inset.Name {
		return set, set[index]
	} else {
		set = append(set, nil)
		copy(set[index+1:], set[index:])
		set[index] = inset
		return set, inset
	}
}

func mergeMarkInput(result []*MarkInput, v Marker) []*MarkInput {
	val := reflect.ValueOf(v)

	val = reflect.Indirect(val)

	if val.Kind() != reflect.Struct {
		return result
	}
	length := val.NumField()
	for i := 0; i != length; i++ {
		field := val.Field(i)
		if !field.CanInterface() {
			continue
		}

		val := field.Interface()

		switch v := val.(type) {
		case *MarkInput:
			if v != nil && v.Name != "" {
				result, v = setMarkInput(result, v)
				field.Set(reflect.ValueOf(v))
			}
		case *Marks:
			if v != nil {
				for _, v := range v.Slice {
					result = mergeMarkInput(result, v)
				}
			}
		case []Marker:
			for _, v := range v {
				result = mergeMarkInput(result, v)
			}
		case Marker:
			if v != nil {
				result = mergeMarkInput(result, v)
			}
		default:
			// No action
		}

	}
	return result
}
