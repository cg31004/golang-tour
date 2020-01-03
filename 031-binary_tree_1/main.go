package main

import "fmt"

//Student binary tree
type Student struct {
	Name  string
	Age   int
	score float64

	Left  *Student
	Right *Student
}

func traversal(Node *Student) {
	if Node == nil {
		return
	}

	fmt.Println(*Node)
	traversal(Node.Left)
	traversal(Node.Right)
}

func main() {
	var Root *Student = &Student{
		Name:  "Tom",
		Age:   20,
		score: 70.0,
	}

	var Left Student = Student{
		Name:  "Amy",
		Age:   18,
		score: 60.0,
	}

	Root.Left = &Left

	var Right *Student = new(Student)
	Right.Name = "Tommy"
	Right.Age = 19
	Right.score = 50

	Root.Right = Right

	left2 := Student{
		Name:  "Tony",
		Age:   50,
		score: 87.5,
	}
	Left.Left = &left2

	traversal(Root)

}
