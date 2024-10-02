package main

/*
Leetcode 27: Remove Element
*/

func removeElement(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == val {
		} else {
			nums[left] = nums[right]
			left++
		}
	}
	return len(nums[:left])

}
