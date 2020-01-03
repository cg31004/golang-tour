package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	map1()
	fmt.Println("=========================")
	map2()
	fmt.Println("=========================")
	map3()
	fmt.Println("=========================")
	map4()
	fmt.Println("=========================")
}

func map1() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -7439967,
	}
	fmt.Println(m["Bell Labs"])
}

func map2() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.11112, -122.08230,
		},
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])

}

func map3() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.11112, -122.08230},
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
}

func map4() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("the value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value: ", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present", ok)
}
