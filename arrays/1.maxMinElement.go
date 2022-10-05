package arrays

import (
	"fmt"
	"math"
)

// Problem Link(s): https://practice.geeksforgeeks.org/problems/find-minimum-and-maximum-element-in-an-array4428/1
func maxMinElement() {
	nums := []int{3, 2, 1, 56, 10000, 167}
	max, min := math.MinInt, math.MaxInt

	for _, val := range nums {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println(max, min)
}
