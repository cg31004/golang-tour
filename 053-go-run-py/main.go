package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("C:\\Users\\simon_su\\AppData\\Local\\Programs\\Python\\Python38-32\\python.exe", "./write_py.py")
	if err := cmd1.Run(); err != nil {
		fmt.Println("Error:  ", err)
	}

	cmd2 := exec.Command("C:\\Users\\simon_su\\AppData\\Local\\Programs\\Python\\Python38-32\\python.exe", "./OutPrint.py")
	out, err := cmd2.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))

}
