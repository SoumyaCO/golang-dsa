package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	res := search(nums, target)
	fmt.Println(res)
}
