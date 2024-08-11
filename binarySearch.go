package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	return binarySearch(nums, target, left, right)
}

func binarySearch(nums []int, target, left, right int) int {
	// this condition is needed for the edge case (if there is no match)
	if left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return binarySearch(nums, target, mid+1, right)
		} else {
			return binarySearch(nums, target, left, mid)
		}
	}
	return -1 // if there is no match it'll return -1
}
