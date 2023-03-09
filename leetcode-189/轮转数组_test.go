package leetcode_189

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 4
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l

	// 双指针
	// 1 2 3 4 5 6 7 1 2 3 4 5 6 7 1 2 3 4 5 6 7
	// k = 3 : 03,14,25,36,   4 5 6 7 2 3 1

	// 1 2 3 4 5 6
	// k=2 : 1 2 5 6 3 4

	// 交换 错的
	//for i := 0; i < k; i++ {
	//	nums[i], nums[l-k+i] = nums[l-k+i], nums[i]
	//}

	// 会申请新的数组存储，不符合题意
	//nums = append(nums[k:], nums[:k]...)
	//fmt.Println(nums)

	// 遍历一遍，也是错的，因为切片共享对的是数组地址
	//tmp := nums[:k]
	//for i := 0; i < l; i++ {
	//	if i < k {
	//		nums[i] = nums[k+i-1]
	//	} else {
	//		nums[i] = tmp[i-k]
	//	}
	//}
}

/*
1 2 3 4 5 6 7

*/
