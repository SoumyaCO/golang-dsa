package main

/*
	Leetcode [26] Remove Duplicates from sorted array
*/

func removeDuplicates(nums []int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return len(nums[:left+1])
}
