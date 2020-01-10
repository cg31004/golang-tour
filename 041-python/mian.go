package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

type Test struct {
	StrPy func(string)*python.PyObject
	Pystr func(*python.PyObject)string
}

func NewTest() *Test {
	err := python.Initialize()
	if err != nil{
		panic(err.Error())
	}
	test := &Test{
		StrPy: python.PyString_FromString,
		Pystr: python.PyString_AsString,
	}
	return test
}

func (test *Test)ImportModule(dir, name string) *python.PyObject{
	module := python.PyImport_ImportModule("sys")
	path := module.GetAttrString("path")
	python.PyList_Insert(path, 0, test.StrPy(dir))
	return python.PyImport_ImportModule(name)

}

func main() {
	test := NewTest()
	module := test.ImportModule("","test")
	f := module.GetAttrString("test")
	argv :=python.PyTuple_New(1)
	python.PyTuple_SetItem(argv, 0, test.StrPy("test"))
	res:= f.call(argv, python.Py_None)
	fmt.Println(test.Pystr(res))
}