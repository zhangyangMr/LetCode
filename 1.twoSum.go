package main

import (
	"fmt"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	nums := []int{3, 2, 4, 5}
	target := 6
	fmt.Printf("twoSum1 %v\n", twoSum1(nums, target))
	fmt.Printf("twoSum2 %v\n", twoSum2(nums, target))
}

//twoSum1和twoSum2是两种方法，方法1消耗内存会比较小，但是比较麻烦；方法二消耗内存大，但是代码简单，不易懂
func twoSum1(nums []int, target int) []int {
	result := make([]int, 0)
	value := make([]int, 0)

	var flag bool
	for key1, num1 := range nums {
		for key2, num2 := range nums {
			if (num1+num2 == target && num1 == num2 && key1 != key2) || (num1+num2 == target && num1 != num2) {
				result = append(result, key1, key2)
				value = append(value, num1, num2)
				flag = true
				break
			}
		}
		if flag == true {
			break
		}
	}
	return result
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
