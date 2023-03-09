package leetcode_278

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;


 */

// 只要这个版本是错的，那么他的后面都是错的
func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}

func firstBadVersion2(n int) int {
	return sort.Search(n, isBadVersion)
}

func isBadVersion(version int) bool {
	return true
}

/*
输入：n = 5, bad = 4
输出：4
解释：
调用 isBadVersion(3) -> false
调用 isBadVersion(5)-> true
调用 isBadVersion(4)-> true
所以，4 是第一个错误的版本。
*/
