package main

/*
Two Sum solved for an array/slice of ordered integers. It uses two pointer pattern to
find the indexes of two numbers.
*/
func twoSumTwoPointers(nums []int, target int) []int {
	l, r := 0, len(nums) -1
	sum := 0

	for l < r {
		sum = nums[l] + nums[r]

		if sum == target {
			return []int{l, r}
		}
		if sum > target {
			r--
		}
		if sum < target {
			l++
		}
	}
	return []int{}
}

func main() {
	nums := []int{1, 3, 4, 7, 9, 15, 17}
	target := 11
	result := twoSumTwoPointers(nums, target)

	for _, v := range result {
		println("index:", v)
	}
}
