package arrays

// Problem Link: https://practice.geeksforgeeks.org/problems/reverse-an-array/0
func ReverseArray() {
	nums := []int{2, 7, 3, 5}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
