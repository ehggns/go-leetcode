package main

/*
topKFrequent finds the k most frequent elements in an array of integers.

This function uses a bucket sort based approach to efficiently find the top k frequent elements:
1. Count the frequency of each number using a hash map
2. Create a bucket where the index represents frequency and the value is a list of numbers with that frequency
3. Traverse the bucket from highest frequency to lowest to collect the top k elements

Parameters:
  - nums: A slice of integers to analyze
  - k: The number of top frequent elements to return

Returns:
  - A slice containing the k most frequent integers in nums

Time Complexity: O(n) where n is the length of nums
Space Complexity: O(n) for storing the frequency map and bucket

Example:

	topKFrequent([]int{1,1,1,2,2,3}, 2) returns [1,2]
*/
func topKFrequent(nums []int, k int) []int {
	// Create a bucket for each frequency. len(nums)+1 to accommodate 0 frequency
	bucket := make([][]int, len(nums)+1)
	// Key is number, value is frequency
	frequencyMap := make(map[int]int)

	// Count the frequency of each number in nums [O(n)]
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Fill the bucket with numbers based on their frequency [O(n)]
	for num, freq := range frequencyMap {
		// Append the number to the corresponding frequency bucket
		bucket[freq] = append(bucket[freq], num)
	}

	// Iterate over the bucket in reverse order to get the top k frequent elements [O(n)]
	result := []int{}
	// len(bucket) - 1 is the max frequency position
	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		// If the bucket is empty, continue to the next frequency
		if len(bucket[i]) == 0 {
			continue
		}
		// Iterate over the numbers in the current frequency bucket
		for _, num := range bucket[i] {
			// Add the number to the result
			result = append(result, num)
			k--
			if k == 0 {
				break
			}
		}
	}
	return result
}

func main() {
	// Example usage
	nums := []int{1, 1, 1, 1, 5, 5, 6, 7}
	k := 2
	result := topKFrequent(nums, k)
	for _, num := range result {
		println(num)
	}
}
