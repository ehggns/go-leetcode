/*
productExceptSelf calculates the product of all elements in the array except self
for each position in the input array.

For an input array nums = [a,b,c,d], it returns an array where:
result[i] = product of all elements in nums except nums[i]
Example: for nums[0], result[0] = b*c*d

The function uses two arrays to store:
  - prefix products (product of all elements to the left);
  - suffix products (product of all elements to the right);

then combines them to get the final result.
This approach avoids using division and achieves O(n) time complexity with O(n) space complexity.

Parameters:
  - nums: input array of integers

Returns:
  - array where each element is the product of all elements in nums except self
*/
package main

// productExceptSelf calculates the product of all elements in the array except self
// for each position in the input array.
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// calculate prefix products (i.e. product of all elements to the left of each index)
	// start from the second element and calculate the product of all elements to the left
	// of the current index
	left := make([]int, n)
	left[0] = 1 // first element has no elements to the left, so product is 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// calculate suffix products (i.e. product of all elements to the right of each index)
	// start from the last element and calculate the product of all elements to the right
	// of the current index
	// last element has no elements to the right, so product is 1
	right := make([]int, n)
	right[n-1] = 1
	penult := n - 2 // last index is n-1, penultimate is n-2
	for i := penult; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1] // using i+1 to pass to the element to the right
	}

	// calculate the result by multiplying the prefix and suffix products
	for i := range n {
		result[i] = left[i] * right[i]
	}

	return result
}

/*
The optimized function should give the same result as the original function
for the same input
The optimized function uses less space by not creating separate prefix and suffix arrays
and directly modifies the result array to store the final product
The time complexity remains O(n) for both functions
The space complexity is reduced to O(1) for the optimized function
as it uses the result array to store the prefix and suffix products
while the original function uses O(n) space for the prefix and suffix arrays
The optimized function is more space-efficient and should be preferred
when space is a concern
The original function is still valid and can be used for clarity
or if the prefix and suffix products need to be accessed separately
*/
func productExceptSelfOptimized(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range n {
		result[i] = 1
	}

	// Calculate prefix products
	prefix := 1
	for i := range n {
		result[i] *= prefix // storing the prefix product in the result array
		prefix *= nums[i]   // preparing the prefix product for the next index
	}
	// Calculate suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix // storing the suffix product in the result array
		suffix *= nums[i]   // preparing the suffix product for the next index
	}
	// The result array now contains the product of all elements except self
	// for each index
	// result[i] = prefix[i] * suffix[i]
	// where prefix[i] is the product of all elements to the left of i and
	// suffix[i] is the product of all elements to the right of i
	// The prefix and suffix products are calculated in a single pass each
	// and combined to get the final result
	return result
}

func main() {
	// Example usage
	nums := []int{5, 2, 3, 1, 4}
	result := productExceptSelf(nums)
	println("Original result:")
	for _, v := range result {
		println(v)
	}

	// Example usage of optimized function
	numsOptimized := []int{5, 2, 3, 1, 4}
	resultOptimized := productExceptSelfOptimized(numsOptimized)
	println("Optimized result:")
	for _, v := range resultOptimized {
		println(v)
	}
}
