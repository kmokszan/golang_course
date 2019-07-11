package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	returnY := make([][]uint8,dy)
	for i := range returnY {
		returnY[i] = make([]uint8,dx)
		for j := range returnY[i] {
			returnY[i][j] = uint8(i*j)
		}
	}
	fmt.Printf("%d\n",len(returnY))
	return returnY
}

func main() {
	pic.Show(Pic)
}
