package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	
	ret := make([][]uint8, dy)
	
	for i, row := range ret {
		row = make([]uint8, dx)
		for j := range row {
			row[j] = 2;
		}
		ret[i] = row
	}
	
	fmt.Println(ret)
		
	return ret
}

func main() {
	pic.Show(Pic)
}
