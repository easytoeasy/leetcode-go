package leetcode_557

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

输入：s = "Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"

思路：
1. 双指针：头尾指针
2. 分割字符串后，再反转，在合并


*/

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	s2 := reverseWords(s)
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", s2)
}

func reverseWords(s string) string {
	var bs = []byte(s)
	head, tail := 0, len(s)-1
	for head <= tail {
		bs[head], bs[tail] = bs[tail], bs[head]
		head++
		tail--
	}
	return string(bs)
}

func TestReverseWords2(t *testing.T) {
	s := "Let's take LeetCode contest"
	s2 := reverseWords2(s)
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", s2)
}

func reverseWords2(s string) string {
	var bs = []byte(s)
	head, tail := 0, 1
	l := len(bs)
	for tail < l {
		for tail < l && bs[tail] != ' ' {
			tail++
		}
		tmp := tail - 1
		for head <= tmp {
			bs[head], bs[tmp] = bs[tmp], bs[head]
			head++
			tmp--
		}
		tail++
		head = tail
	}
	return string(bs)
}

func TestReverseWords3(t *testing.T) {
	s := "Let's take LeetCode contest"
	s2 := reverseWords3(s)
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", s2)
}

func reverseWords3(s string) string {
	var s2 []string
	strs := strings.Split(s, " ")
	for _, v := range strs {
		head, tail := 0, len(v)-1
		bs := []byte(v)
		for head <= tail {
			bs[head], bs[tail] = bs[tail], bs[head]
			head++
			tail--
		}
		s2 = append(s2, string(bs))
	}
	return strings.Join(s2, " ")
}
