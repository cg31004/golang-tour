package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("040-file\\test.txt")
	fmt.Println(f)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	// l, err := f.WriteString("Hello World")
	// fmt.Println(l)
	// fmt.Println(err)
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// fmt.Println(l, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("-------------------------------")
	fi, er := os.Open("040-file\\test.txt")
	defer fi.Close()
	if er != nil {
		log.Fatal(er)
	}
	b, _ := ioutil.ReadAll(fi)
	ns := string(b)
	fmt.Print(ns)
}
