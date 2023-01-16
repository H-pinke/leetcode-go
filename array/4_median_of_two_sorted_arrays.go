package main

import (
	"fmt"
	"sort"
)

//4. 寻找两个正序数组的中位数 困难
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2

func main() {
	//时间复杂度 O(m+n)
	res := findMedianArrays([]int{1, 3}, []int{2, 4})
	fmt.Println(res)
}

func findMedianArrays(num1 []int, num2 []int) float64 {
	res := merge(num1, num2)
	if res == nil {
		return -1
	}
	n := len(res)
	if n%2 == 0 {
		return float64(res[n/2-1]+res[n/2]) / 2
	}
	return float64(res[n/2])
}

func merge(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 && n2 == 0 {
		return nil
	}
	if n1 < n2 {
		for _, v := range nums1 {
			nums2 = append(nums2, v)
		}
		sort.Ints(nums2)
		return nums2
	} else {
		for _, v := range nums2 {
			nums1 = append(nums1, v)
		}
		sort.Ints(nums1)
		return nums1
	}
}
