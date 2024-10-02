package main

import (
	"fmt"
	"math"
)

/*
Leetcode 66: Plus One
nums := []int{1, 2, 9}
*/

func plusOne(digits []int) []int {
	number := 0
	var re_arr []int
	for d := 0; d < len(digits); d++ {
		number += int(digits[d]) * int(math.Pow(10, float64(len(digits)-1-d)))
	}
	number = number + 1
	fmt.Println("Number till now: ", number)

	re_arr = append(re_arr, number/10)
	re_arr = append(re_arr, number%10)

	return re_arr

}
