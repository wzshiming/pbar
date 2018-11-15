package pbar

import (
	"encoding/json"
	"unsafe"
)

type MarkTab struct {
	json.RawMessage
}

func (m *MarkTab) String() string {
	v := m.RawMessage[:]
	if len(v) == 0 {
		return ""
	}
	if v[0] == '"' {
		v = v[1 : len(v)-1]
	}
	return *(*string)(unsafe.Pointer(&v))
}
