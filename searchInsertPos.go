package main

/*
Leetcode 35: Search Insert Position
*/

func binarySh(nums []int, left, right, val int) int {
	mid := left + (right-left)/2
	if left < right {
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			return binarySh(nums, left, mid, val)
		} else {
			return binarySh(nums, mid+1, right, val)
		}
	}
	return mid
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target < nums[0] {
		return 0
	} else {
		return binarySh(nums, left, right, target)
	}
}
