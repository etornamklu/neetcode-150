package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {

	result := make([]int, len(nums))

	for i := range result {
		result[i] = 1
	}

	prefix := 1
	for i := range nums {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix // what is at result[i] * postfix
		postfix *= nums[i]   // for each iteraction update postfix, postfix = postfix * nums[i]

	}

	return result

}
