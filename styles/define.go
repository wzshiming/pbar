package styles

import (
	"bytes"
	"text/template"
	"unsafe"

	"github.com/wzshiming/pbar"
)

var Normal, _ = NewConfig(MustAsset("normal.json"))

type Config struct {
	temp *template.Template
}

func NewConfig(conf []byte) (*Config, error) {
	config := *(*string)(unsafe.Pointer(&conf))
	temp, err := template.New("").Parse(config)
	if err != nil {
		return nil, err
	}
	return &Config{
		temp: temp,
	}, nil
}

func (c *Config) New(option map[string]interface{}) (pbar.Mark, error) {
	buf := bytes.NewBuffer(nil)
	err := c.temp.Execute(buf, option)
	if err != nil {
		return nil, err
	}
	marks := &pbar.Marks{}
	err = marks.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return marks, nil
}
