package leetcode_977

import (
	"sort"
	"testing"
)

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
/*
示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

func TestSortedSquares(t *testing.T) {
	//nums := []int{-3, -2, -1, 0, 1, 2}
	nums := []int{-7, -3, 2, 3, 11}
	ret := sortedSquares(nums)
	t.Log(ret)
}

// 双指针：left right 都指向中间的数
// 如果都是正数，不需要排序
// 如果包含负数，则需要排序
func sortedSquares(nums []int) []int {
	square := func(n int) int {
		return nums[n] * nums[n]
	}
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}
	var ret = make([]int, len(nums))
	if nums[0] >= 0 {
		for k, v := range nums {
			ret[k] = v * v
		}
		return ret
	}
	var i, j = 0, len(nums) - 1
	var index = j
	for i < j {
		if square(i) > square(j) {
			ret[index] = square(i)
			i++
		} else {
			ret[index] = square(j)
			j--
		}
		index--
	}
	ret[index] = square(i)
	return ret
}

func TestSortedSquares2(t *testing.T) {
	//nums := []int{-3, -2, -1, 0, 1, 2}
	nums := []int{-7, -3, 2, 3, 11}
	ret := sortedSquares2(nums)
	t.Log(ret)
}

func sortedSquares2(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}
	var ret = make([]int, len(nums))
	if nums[0] >= 0 {
		for k, v := range nums {
			ret[k] = v * v
		}
		return ret
	}
	var i, j = 0, len(nums) - 1
	var index = j
	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	for i < j {
		if abs(nums[i]) > abs(nums[j]) {
			ret[index] = nums[i] * nums[i]
			i++
		} else {
			ret[index] = nums[j] * nums[j]
			j--
		}
		index--
	}
	ret[index] = nums[i] * nums[i]
	return ret
}

func TestSortedSquares3(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	ret := sortedSquares3(nums)
	t.Log(ret)
}

func sortedSquares3(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}
	if nums[0] >= 0 {
		for k, v := range nums {
			nums[k] = v * v
		}
		return nums
	}

	if nums[len(nums)-1] <= 0 {
		for k, v := range nums {
			nums[k] = v * v
		}
		sort.Reverse(sort.IntSlice(nums))
		return nums
	}

	var i, j = 0, len(nums) - 1
	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	for i < j {
		if abs(nums[i]) > abs(nums[j]) {
			nums[i], nums[j] = nums[j], nums[i]*nums[i]
			j--
		} else {
			nums[j] = nums[j] * nums[j]
			j--
		}
	}
	nums[j] = nums[i] * nums[i]
	return nums
}
