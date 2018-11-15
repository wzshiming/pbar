package pbar

import (
	"reflect"
	"strings"
)

var markMap = map[string]reflect.Type{}

func init() {
	for _, v := range []Marker{
		&MarkBar{}, &MarkAfter{},
		&MarkPer{}, &MarkPerCent{},
		&MarkRoll{}, &MarkScroll{},
		&MarkInput{},
	} {
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
