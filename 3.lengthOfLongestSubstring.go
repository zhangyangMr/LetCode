package main

import (
	"fmt"
)

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func main() {
	str := "pwwkew"
	fmt.Printf("lengthOfLongestSubstring : %v\n", lengthOfLongestSubstring(str))
}

//解题思路，标记两个重复字段的下标值，取两个下标值之差
func lengthOfLongestSubstring(s string) int {
	var result int
	flag := make(map[byte]int, 0)
	for i, j := 0, 0; i < len(s); i++ {
		flag[s[i]]++
		for flag[s[i]] > 1 {
			flag[s[j]]--
			j++
		}
		result = max(result, i-j+1)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
