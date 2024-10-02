package main

/*
Leetcode 1: Two Sum
*/

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for index, value := range nums {
		diff := target - value
		if _, ok := hashMap[diff]; ok {
			return []int{index, hashMap[diff]}
		}
		hashMap[value] = index
	}
	return nil
}
