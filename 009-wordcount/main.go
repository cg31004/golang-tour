package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	sf := strings.Fields(s)
	smap := make(map[string]int)
	for i := 0; i < len(sf); i++ {
		if smap[sf[i]] == 0 {
			smap[sf[i]] = 1
		} else {
			smap[sf[i]] = smap[sf[i]] + 1
		}
	}
	return smap
}

func main() {
	wc.Test(WordCount)
}
