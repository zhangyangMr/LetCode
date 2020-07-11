package main

import (
	"fmt"
)

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/

func main() {
	str := "abcda"
	//fmt.Printf("result : %v\n", longestPalindrome1(str))
	fmt.Printf("result : %v\n", longestPalindrome2(str))
}

//方法一：暴力破解法，循环把所有的子串全部枚举出来，然后判断所有子串是否是回文子串
func longestPalindrome1(s string) string {
	//先对字符串的长度为1和2进行特殊处理
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && checkPalindrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

//验证是否为回文子串
func checkPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			right--
			left++
		}
	}
	return true
}

//方法二：动态规划
func longestPalindrome2(s string) string {
	//当字符串的长度小于2时，一个字符和空字符肯定是回文子串
	if len(s) < 2 {
		return s
	}

	//定义回文子串的最大长度，回文子串的开始下表
	var maxLen int = 1
	var begin int = 0
	//定义二维数组
	dp := make([][]bool, len(s))

	//动态规划方程建立，实际上使填写一个二维表格的过程，需要满足的条件
	/*
		1. 生名一个二维数组，用来记录回文子串的起始下标和终止下标，d[i][j]=s[i:j], 是回文子串赋值true，不是回文子串赋值false
		2.表格上对角线的下标都是单字符回文子串，肯定是回文子串
		3.回文子串的判断条件，如果d[i][j]=s[i:j]的状态取决于首尾字符是否相当，以及d[i+1][j-1]是否为回文子串
		4.当有下标计算时，就要考虑下标越界的问题，当数组长度小于2时就没有意义了，只剩下一个字符，肯定是回文子串，即j-1-(i+1)+1<2 => j-i<3
	*/
	//给表的对角线进行赋值
	for i, _ := range s {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	//循环的条件是只需要填写对角线上方的表格即可，本文的填表方法是一列一列填表，外层循环为字符串的右边界，内层循环为字符串的左边界
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] != s[i] {
				dp[j][i] = false
			} else {
				if i-j < 3 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
			if maxLen < i-j+1 && dp[j][i] {
				maxLen = i - j + 1
				begin = j
			}
		}
	}
	return s[begin : begin+maxLen]
}
