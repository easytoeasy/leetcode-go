package leetcode_213

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	ret1 := rob(nums[0 : len(nums)-1])
	ret2 := rob(nums[1:])
	t.Log(getMax(ret1, ret2))
}

func rob(nums []int) int {
	fmt.Println(nums)
	var dp = make([][2]int, len(nums))
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + v
	}

	return getMax(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*

定义状态：
（1）进: dp[i][1] = dp[i-1][0] + nums[i]
（2）出: dp[i][0] = max(dp[i-1][0], dp[i-1][1])

*/
