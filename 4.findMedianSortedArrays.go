package main

import "fmt"

/*
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
*/

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//将最短的数组作为nums1
	if len(nums2) < len(nums1) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	length1, length2 := len(nums1), len(nums2)
	//算出分割线左边元素的个数
	k := (length1 + length2 + 1) / 2
	left := 0
	right := length1
	//在nums1中找到恰当的分割线后，nums2中的分割线=k-nums1找到的左边元素个数
	// 取到分割线时，需要满足的条件，分割线交叉比较
	// 第一个数组分割线左边的元素大小小于第二个分割线右边的元素大小，同时第一个数组分割线右边的元素要大于第二个数组分割线左边的元素
	// 即nums1[i-1]<=nums2[j] && nums2[j-1]<=nums1[i]

	for left < right {
		//找到第一个数组的中间元素
		i := left + (right-left+1)/2
		//找到第二个数组的中间元素,保证左右两边的元素个数一致(第一个数组左边取了i个元素，第二个数组左边需要取k-i个元素)
		j := k - i
		//如果左边大就向左边找，如果左边小就向右边找，知道刚好满足条件
		if nums1[i-1] > nums2[j] {
			//下一轮搜索区间[left,i-1]
			right = i - 1
		} else {
			//下一个搜索区间[i,right]
			left = i
		}
	}

	//退出循环最终确定分割线的位置
	i := left
	j := k - i
	var nums1LeftMax int
	var nums1RightMin int
	var nums2LeftMax int
	var nums2RightMin int
	//数组越界问题处理：
	// ^int(^uint(0) >> 1)整型数的最小值
	//int(^uint(0) >> 1)整型数的最大值
	if i == 0 {
		nums1LeftMax = ^int(^uint(0) >> 1)
	} else {
		nums1LeftMax = nums1[i-1]
	}
	if i == length1 {
		nums1RightMin = int(^uint(0) >> 1)
	} else {
		nums1RightMin = nums1[i]
	}
	if j == 0 {
		nums2LeftMax = ^int(^uint(0) >> 1)
	} else {
		nums2LeftMax = nums2[j-1]
	}
	if j == length2 {
		nums2RightMin = int(^uint(0) >> 1)
	} else {
		nums2RightMin = nums2[j]
	}
	if (length1+length2)%2 == 1 {
		return float64(Max(nums1LeftMax, nums2LeftMax))
	} else {
		return float64(Max(nums2LeftMax, nums1LeftMax)+Min(nums2RightMin, nums1RightMin)) / 2
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
