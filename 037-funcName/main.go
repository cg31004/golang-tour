package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// test1
	fmt.Println("============   test1   ==============")
	fmt.Printf("%v : %d\n", GetFunctionName(foo), foo())
	fmt.Println("============   test2   ==============")
	funcName, funcValue := foo2()
	fmt.Printf("%v : %d\n", funcName, funcValue)
}

//============   test1   ==============
func foo() int {
	return 1
}

//GetFunctionName  GetFunctionName
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//============   test2   ==============
func foo2() (string, int) {
	return GetFunctionNam2(), 2
}

//GetFunctionName2  getFunctionName2
func GetFunctionNam2() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
