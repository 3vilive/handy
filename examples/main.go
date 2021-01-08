package main

import (
	"fmt"

	"github.com/3vilive/handy/fn"
)

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// filter odd numbers
	intArr = fn.FilterInt(intArr, func(i int) bool { return i%2 != 0 })
	fmt.Printf("%#v\n", intArr)

	// double it
	intArr = fn.MapIntToInt(intArr, func(i int) int { return i * 2 })
	fmt.Printf("%#v\n", intArr)

	// reverse it
	fn.ReverseInt(intArr)
	fmt.Printf("%#v\n", intArr)

	// sum
	sumOfArr := fn.ReduceInt(intArr, func(accm int, curr int) int { return accm + curr })
	fmt.Println(sumOfArr)
}
