package styles

import (
	"github.com/wzshiming/pbar"
)

func Normal() *pbar.Marks {
	bar, _ := OpenBuiltinStyle("normal.json")
	return bar
}

type Config struct {
	conf []byte
}

func OpenBuiltinStyle(name string) (*pbar.Marks, error) {
	conf, err := Asset(name + ".json")
	if err != nil {
		return nil, err
	}
	return NewConfig(conf)
}

func NewConfig(conf []byte) (*pbar.Marks, error) {
	marks := &pbar.Marks{}
	err := marks.UnmarshalJSON(conf)
	if err != nil {
		return nil, err
	}
	return marks, nil
}
