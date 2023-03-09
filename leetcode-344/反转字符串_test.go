package leetcode_344

import "testing"

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]

思路：
1. 头尾指针
2. 递归
3. 一次循环
*/

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	t.Log(s)
}

func reverseString(s []byte) {
	head, tail := 0, len(s)-1
	for head <= tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}

func reverseString2(s []byte) {
	var helper func(i, j int)
	helper = func(i, j int) {
		if i >= j {
			return
		}
		s[i], s[j] = s[j], s[i]
		helper(i+1, j-1)
	}
	helper(0, len(s)-1)
}

func reverseString3(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
