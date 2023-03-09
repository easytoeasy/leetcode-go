package leetcode_283

import (
	"fmt"
	"testing"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func TestMoveZero2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 1}
	//nums := []int{0, 0}
	//nums := []int{1, 0}
	//nums := []int{1, 0, 1}
	moveZeroes2(nums)
	fmt.Println(nums)
}

func moveZeroes2(nums []int) {
	var a, b = 0, 0
	l := len(nums)
	for b = 0; b <= l-1; b++ {
		if nums[b] != 0 {
			nums[a] = nums[b]
			a++
		}
	}
	for a <= l-1 {
		nums[a] = 0
		a++
	}
}

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 1}
	//nums := []int{0, 0}
	//nums := []int{1, 0}
	//nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 8.92%
// 6.26%
// 退出条件开始没找到
func moveZeroes(nums []int) {
	var i, j = 0, 0
	var l = len(nums)
	if l == 1 {
		return
	}
	for j <= l-1 && i <= l-1 {
		for i <= l-1 && nums[i] != 0 {
			i++
			j = i
		}
		for j <= l-1 && nums[j] == 0 {
			j++
		}
		if j <= l-1 && i <= l-1 && i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]

// 0 1 0 3 12
// 1 3 12 0 0
