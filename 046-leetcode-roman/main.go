package main

import "fmt"

var ans int

func romanToInt(s string) int {
	StringMap := make(map[string]int)
	var result int
	var close int
	SLen := len(s)
	for i, v := range s {
		if _, ok := StringMap[string(v)]; ok {
			StringMap[string(v)] = StringMap[string(v)] + 1
		} else {
			StringMap[string(v)] = 1
		}
		if (SLen - 1) != i {
			switch string(s[i]) {
			case "I":
				if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
					close = close - 2
				}
			case "X":
				if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
					close = close - 20
				}
			case "C":
				if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
					close = close - 20
				}
			}
		}
	}
	for i, v := range StringMap {
		result += StrToNum(i, v)
	}
	result = result + close
	return result
}

func StrToNum(Symbol string, times int) int {
	SymbolMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	ans := 0
	base_num := SymbolMap[Symbol]
	ans = times * base_num
	return ans
}

func main() {

	// r := romanToInt("XIII")
	rChannel := chanRomanToInt("XIII")
	fmt.Println(rChannel)

}

func chanRomanToInt(s string) int {
	StringMap := make(map[string]int)
	var result int
	var close int
	SLen := len(s)
	for i, v := range s {
		if _, ok := StringMap[string(v)]; ok {
			StringMap[string(v)] = StringMap[string(v)] + 1
		} else {
			StringMap[string(v)] = 1
		}
		if (SLen - 1) != i {
			switch string(s[i]) {
			case "I":
				if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
					close = close - 2
				}
			case "X":
				if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
					close = close - 20
				}
			case "C":
				if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
					close = close - 20
				}
			}
		}
	}
	for i, v := range StringMap {
		result += StrToNum(i, v)
	}
	result = result + close
	return result
}
