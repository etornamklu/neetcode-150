package threesum

import "sort"

func ThreeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := range nums {

		a := nums[i]

		// once a > 0, any future triplet will never sum to 0 because array is sorted
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for r > l {
			if a+nums[l]+nums[r] > 0 {
				r--
			} else if a+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}

// explanation on the else block
// When you find a triplet that sums to zero (a + nums[l] + nums[r] == 0), you:
// Append the triplet to the result.
// Move both pointers inwards:
// Because you're done with nums[l] and nums[r]
// You want to look for the next possible unique pair
// Skip duplicates:
// nums[l] == nums[l-1] means you're seeing the same number again
// You don't want to include duplicate triplets like [a, 2, 3] and [a, 2, 3]
// So you move l forward to skip over same values
