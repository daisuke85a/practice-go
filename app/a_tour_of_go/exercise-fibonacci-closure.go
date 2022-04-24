package main

import "fmt"

func fibonacci() func() int {
	var i int
	recentRet := [2]int{0, 1}
	return func() int {
		if i == 0 || i == 1 {
			ret := recentRet[i]
			i++
			return ret
		}
		ret := recentRet[0] + recentRet[1]
		recentRet[0] = recentRet[1]
		recentRet[1] = ret
		i++
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
