package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	if str=="" {
		return map[string]int{}
	}
	m := make(map[string]int)
	for _, s := range strings.Fields(str) {
		v, ok := m[s]
		if ok {
		m[s] = v + 1
		} else {
		m[s] = 1
		}
	}
	return m
}

func main() {
	fmt.Print(WordCount("ala ma kota bo kota ala lubi"))
	wc.Test(WordCount)
}