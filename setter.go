package main

import (
	"fmt"
)

/*
	这种方法的好处在于工厂函数参数是可变长度的
*/

type Option struct {
	Field1 int
	Field2 string
	Field3 float32
}

func (this *Option) Set(sets ...SetOption) {
	for _, set := range sets {
		set(this)
	}
}

type SetOption func(*Option)

func SetField1(f int) SetOption {
	return func(o *Option) {
		o.Field1 = f
	}
}

func SetField2(f string) SetOption {
	return func(o *Option) {
		o.Field2 = f
	}
}

func SetField3(f float32) SetOption {
	return func(o *Option) {
		o.Field3 = f
	}
}

func NewOption(sets ...SetOption) *Option {
	op := &Option{}
	for _, set := range sets {
		set(op)
	}
	return op
}

func main() {
	opt := NewOption(SetField1(1), SetField2("ssss"), SetField3(0.1))
	opt.Set(SetField1(2))
	fmt.Println(opt)
}
