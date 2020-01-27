package main

import "fmt"

func searchH(nums []int, start, end, target int) int {

	if start > end {
		return -1
	}

	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}

	mid := (start + end + 1) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < nums[start] || nums[mid] > nums[end] {
		goneLeft := searchH(nums, start, mid-1, target)
		if goneLeft != -1 {
			return goneLeft
		}
		return searchH(nums, mid+1, end, target)
	} else {
		if nums[mid] > target {
			return searchH(nums, start, mid-1, target)
		} else {
			return searchH(nums, mid+1, end, target)
		}
	}
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return searchH(nums, 0, len(nums)-1, target)
}

func main() {
	fmt.Printf("index of target %d\n", search([]int{6,7,0}, 7))
}
