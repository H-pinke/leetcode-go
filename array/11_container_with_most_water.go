package main

import "fmt"

/**
盛最多水的容器
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/
func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func maxArea(area []int) int {
	i, j := 0, len(area)-1
	maxNum := 0
	for i < j {
		w := j - i
		h := min(area[i], area[j])
		maxNum = max(maxNum, w*h)
		if area[i] <= area[j] {
			i++
		} else {
			j--
		}
	}
	return maxNum
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
