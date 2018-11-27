package deepclone

import (
	"reflect"
)

// Clone a complete clone
func Clone(v interface{}) interface{} {
	c := &cloner{}
	r := c.Clone(v)
	return r
}

// CloneValue a complete clone for reflect.Value
func CloneValue(v reflect.Value) reflect.Value {
	c := &cloner{}
	r := c.CloneValue(v)
	return r

}

type cloner struct {
	m map[uintptr]*reflect.Value
}

func (c *cloner) Clone(v interface{}) interface{} {
	return c.CloneValue(reflect.ValueOf(v)).Interface()
}

func (c *cloner) CloneValue(v reflect.Value) reflect.Value {
	c.m = map[uintptr]*reflect.Value{}
	r := c.cloneValue(&v)
	return *r
}

func (c *cloner) cloneValue(v *reflect.Value) (r *reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			return v
		}
		ve := v.Elem()

		return c.cloneValue(&ve)

	case reflect.Ptr:
		if v.IsNil() {
			return v
		}

		// 标记指针 处理循环引用
		point := v.Pointer()
		if nn, ok := c.m[point]; ok {
			return nn
		}

		nn := reflect.New(v.Type().Elem())
		c.m[point] = &nn

		ve := v.Elem()
		nt := c.cloneValue(&ve)
		nn.Elem().Set(*nt)

		return &nn

	case reflect.Struct:
		nf := v.NumField()
		rr := reflect.New(v.Type())
		nt := rr.Elem()

		for i := 0; i != nf; i++ {
			mi := v.Field(i)
			if !mi.CanSet() {
				continue
			}
			nf := nt.Field(i)
			mv := c.cloneValue(&mi)

			nf.Set(*mv)
		}

		return &nt

	case reflect.Map:
		nt := reflect.MakeMap(v.Type())
		for _, i := range v.MapKeys() {
			mi := v.MapIndex(i)
			mk := c.cloneValue(&i)
			mv := c.cloneValue(&mi)

			nt.SetMapIndex(*mk, *mv)
		}

		return &nt

	case reflect.Array:
		nt := reflect.New(v.Type()).Elem()
		l := nt.Len()

		for i := 0; i != l; i++ {
			mi := v.Index(i)
			ni := nt.Index(i)
			mv := c.cloneValue(&mi)

			ni.Set(*mv)
		}
		return &nt

	case reflect.Slice:
		l := v.Len()
		nt := reflect.MakeSlice(v.Type(), l, v.Cap())

		for i := 0; i != l; i++ {
			mi := v.Index(i)
			ni := nt.Index(i)
			mv := c.cloneValue(&mi)

			ni.Set(*mv)
		}
		return &nt

	default:
		return v
	}
}
