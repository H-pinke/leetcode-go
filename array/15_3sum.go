package main

import (
	"fmt"
	"sort"
)

/**
15. 三数之和 中等
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	sum := 0
	var threeSums [][]int
	for i, num := range nums {
		if num > 0 {
			break
		}
		//向前去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum = num + nums[l] + nums[r]
			switch {
			case sum < 0:
				l++
			case sum > 0:
				r--
			case sum == 0:
				threeSums = append(threeSums, []int{num, nums[l], nums[r]})
				//向后去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				//向后去重
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return threeSums
}
