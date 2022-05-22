/**
1. 两数之和 简单

给定义一个整数数组 nums 和一个整数目标值 target, 请你在该数组 中找出和为目标值的 两个整数，并返回它们的数组下标。
你可以假设每种输入 只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

package main

import "fmt"

func main() {
	fmt.Println(towSum1([]int{2, 7, 11, 15}, 9))
	fmt.Println(towSum2([]int{2, 7, 11, 15}, 9))
	fmt.Println(towSum3([]int{2, 7, 11, 15}, 9))
}

/**
方法1：
暴力双重遍历
*/
func towSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/**
方法2
哈希表索引
检查数组中是否存在目标元素， 若存在则找出索引
哈希表特别适合抽象类的配对结构，当要解决问题的数据单元是成对数据关系时，考虑哈希表map结构
*/
func towSum2(nums []int, target int) []int {
	num2Index := make(map[int]int, len(nums))
	for i, num := range nums {
		num2Index[num] = i
	}
	for i, num := range nums {
		pair := target - num
		if j, ok := num2Index[pair]; ok && i != j { //剔除自身相加的情况
			return []int{i, j}

		}
	}
	return nil
}

/**
方法3
优化后的哈希表 索引方式
看 towSum2 会发现进行了2 次 for 循环，可以进行合并优化，一边遍历，一边检查
*/
func towSum3(nums []int, target int) []int {
	num2Index := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num
		if j, ok := num2Index[pair]; ok && i != j {
			return []int{i, j}
		}
		num2Index[num] = i
	}
	return nil
}
