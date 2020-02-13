package main

func intToRoman(num int) string {
	RomanDic := map[string]int{
		"M": 1000,
		"D": 500,
		"X": 10,
		"C": 100,
		"L": 50,
		"V": 5,
		"I": 1,
	}
	for i, v := range RomanDic {
		while num-i < 0 {
			
		}
	}
	return "123"
}

func main() {
	intToRoman(25)
}
