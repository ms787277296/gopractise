package server

import (
	"reflect"

	"github.com/golang/glog"
)

func WhichType(i interface{}) {
	glog.Info("the type of i is ", reflect.TypeOf(i))

	value := reflect.ValueOf(i)
	glog.Info(value.Type(), value)

	return
}

type Student struct {
	ID        int
	FirstName string
}

func ReflectStruct(s interface{}) {
	s1 := reflect.ValueOf(s)
	value := s1.Field(1)

	glog.Info(s1.Field(1))
	glog.Info(value)

}

func ReflectPractise() {
	// var i int = 6
	// WhichType(i)

	// var a interface{} = 10
	// t := a.(int)

	// glog.Info(reflect.TypeOf(t), t)

	s := Student{1, "zrj"}
	ReflectStruct(s)

	return
}
