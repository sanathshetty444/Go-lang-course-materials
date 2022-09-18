package main

import (
	"fmt"
)

var e int

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}
}
