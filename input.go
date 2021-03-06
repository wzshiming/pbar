package pbar

import (
	"bytes"
	"encoding/json"
	"unsafe"
)

type MarkInput struct {
	Kind string          `json:"_Kind"`
	Name string          `json:"_Name"`
	Val  json.RawMessage `json:",omitempty"`
}

func (m *MarkInput) Input(val string) {
	m.Val = *(*json.RawMessage)(unsafe.Pointer(&val))
}

func (m *MarkInput) Float64() (float64, error) {
	return json.Number(m.String()).Float64()
}

func (m *MarkInput) Int64() (int64, error) {
	return json.Number(m.String()).Int64()
}

func (m *MarkInput) String() string {
	v := m.Val[:]
	if len(v) == 0 {
		return ""
	}
	if v[0] == '"' {
		v = m.Val[1 : len(m.Val)-1]
	}
	return *(*string)(unsafe.Pointer(&v))
}

type _markInput MarkInput

func (m *MarkInput) MarshalJSON() ([]byte, error) {
	if m.Name == "" {
		return []byte(m.Val), nil
	}
	return json.Marshal(_markInput(*m))
}

func (m *MarkInput) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)

	if data[0] == '{' {
		var markInput _markInput
		err := json.Unmarshal(data, &markInput)
		if err != nil {
			return err
		}
		*m = MarkInput(markInput)
	} else {
		m.Val = json.RawMessage(data)
	}
	return nil
}
