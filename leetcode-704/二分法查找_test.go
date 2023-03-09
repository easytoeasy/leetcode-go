package leetcode_704

import (
	"sort"
	"testing"
)

func Test_Search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	ret := search(nums, 0, 9)
	t.Log(ret)
}

// 要点：记录每次二分之前的索引位置
// 效率不高
func search(nums []int, index, target int) int {
	l := len(nums)
	if l == 1 && target != nums[l-1] {
		return -1
	}
	mid := nums[l/2]
	if mid == target {
		return index + l/2
	} else if mid > target {
		return search(nums[0:l/2], index, target)
	} else {
		return search(nums[l/2:], index+l/2, target)
	}
}

func Test_Search2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	ret := search2(nums, 9)
	t.Log(ret)
}

func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func Test_Search3(t *testing.T) {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	nums := []int{2, 5}
	ret := search3(nums, 5)
	t.Log(ret)
}

// 这个有问题，并不适合这个地方
func search3(nums []int, target int) int {
	return sort.Search(target, func(i int) bool {
		// if false and i = h + 1
		// if true and j = h
		if nums[i] < target { // 前半区间找
			return true
		}
		return false // 后半区间找
	})
}

func Test_Search4(t *testing.T) {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	nums := []int{2, 5}
	ret := search4(nums, 5)
	t.Log(ret)
}

func search4(nums []int, target int) int {
	i, found := sort.Find(target, func(i int) int {
		if nums[i] == target {
			return 0
		} else if nums[i] > target {
			return -1
		} else {
			return 1
		}
	})
	if found {
		return i
	}
	return -1
}

/*
（1）暴力破解法
（2）二分查找法
思路：有序数组，从中比较大小。如果比target大，则往后半区间查找，否则往前半区间查找。

*/
