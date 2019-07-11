package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	ret := 1
	prev1 := 1
	prev2 := 1
	iter := 1
	return func ()int {
		if iter <= 2 {
			prev1 = 1
			prev2 = 1
			ret = 1
		}else{
			ret = prev1 + prev2
			prev2 = prev1
			prev1 = ret
		}
		iter += 1
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
