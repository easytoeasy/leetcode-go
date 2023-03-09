package leetcode_35

import (
	"sort"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 5, 5, 5, 5, 5, 6}
	target := 5
	ret := searchInsert(nums, target)
	t.Log(ret)
}

// 找到第一个小于target的数, 下面的只是找到相等的数
// 左边是小于，右边是等于或者大于
func searchInsert(nums []int, target int) int {
	var i, j = 0, len(nums) - 1
	for i <= j {
		mid := int(uint(i+j) >> 1)
		//mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { //右区间查找
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

func TestSearchFirstInsert(t *testing.T) {
	nums := []int{1, 3, 5, 5, 5, 5, 5, 5, 6}
	target := 5
	ret := searchFirstInsert(nums, target)
	t.Log(ret)
}

// 找到第一个插入的位置
func searchFirstInsert(nums []int, target int) int {
	var i, j = 0, len(nums) - 1
	for i <= j {
		mid := int(uint(i+j) >> 1)
		//mid := i + (j-i)/2
		if nums[mid] < target { //右区间查找
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

func TestSearchBySorted(t *testing.T) {
	nums := []int{1, 3, 5, 5, 5, 5, 5, 5, 6}
	target := 7
	ret := searchBySorted(nums, target)
	t.Log(ret)
}

func searchBySorted(nums []int, target int) int {
	return sort.Search(target, func(i int) bool {
		if nums[i] > target {
			return true
		}
		return false
	})
}
