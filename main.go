package main

import (
	"fmt"
)

func main() {
	fmt.Println("----DSA with Golang-----")
	// nums := []int{1, 2, 9}
	// fmt.Println(plusOne(nums))

	number := 130
	arr := []int{}
	for number > 10 {
		arr = append(arr, number/100)
		arr = append(arr, number/10)
	}

}
