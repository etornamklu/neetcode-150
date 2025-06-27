# 3Sum – Find Unique Triplets That Sum to Zero

## Problem Statement

Given an **integer array** `nums`, return all **unique triplets** `[nums[i], nums[j], nums[k]]` such that:

* `i ≠ j ≠ k`
* `nums[i] + nums[j] + nums[k] == 0`

### Notes:

* The solution set **must not contain duplicate triplets**.
* Output can be in **any order**.

---

## Problem Type

* Arrays
* Two Pointers
* Sorting
* Duplicate Handling
* Triple Sum Pattern

---

## Trick

### 1. **Sort the array**

This helps with:

* Skipping duplicates
* Using the two-pointer strategy

### 2. **Fix one number** `a = nums[i]`, and find two numbers `b` and `c` such that:

```
a + b + c == 0  ⟶  b + c == -a
```

Use two pointers (`l`, `r`) to search for `b + c` in the remaining subarray.

---

## Intuition

Once the array is sorted:

* Fix each number `a = nums[i]` and use the two-pointer method on the subarray `nums[i+1:]` to find pairs that sum to `-a`.
* Skip duplicate elements for both `a` and `b` to avoid repeating triplets.
* Since we skip duplicates and use sorted array, the output triplets are naturally unique.

---

## Go Code

```go
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := range nums {
		a := nums[i]

		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for r > l {
			sum := a + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++ // skip duplicate b's
				}
			}
		}
	}

	return res
}
```

---

## Walkthrough Example

Input:
`nums = [-1, 0, 1, 2, -1, -4]`

### Sorted:

`[-4, -1, -1, 0, 1, 2]`

### Triplets Found:

| a  | b  | c   | Sum | Included? | Reason             |
|----|----|-----|-----|-----------|--------------------|
| -1 | -1 | ... |     | skip      | duplicate `a`      |
| -1 | 0  | 1   | 0   | ✅         | valid & unique     |
| -1 | -1 | 2   | 0   | ✅         | valid & unique     |
| 0  | 1  | 2   | 3   | skip      | sum > 0 → no match |

Output: `[[-1, -1, 2], [-1, 0, 1]]`

---

## Time and Space Complexity

| Operation      | Time       | Space                    |
| -------------- | ---------- | ------------------------ |
| Sorting        | O(n log n) | O(1)                     |
| Main loop      | O(n²)      | O(1) extra               |
| Result storage | —          | O(k) for output triplets |

**Overall Time:** ✅ `O(n²)`
**Overall Space:** ✅ Constant extra (excluding output)

---

## Go Feature Highlights

* `sort.Ints(nums)` – Sorts in-place
* `res = append(...)` – Appends new triplet slice
* Skips duplicates using comparisons like `nums[i] == nums[i-1]` and `nums[l] == nums[l-1]`

---

## Key Insights

✅ Leverages sorting to simplify logic
✅ Two-pointer scan reduces from `O(n³)` to `O(n²)`
✅ Skips duplicate values at both the outer and inner loop levels
✅ Clean, efficient, and commonly asked interview question

---

## Summary Table

| Step             | Description                                     |
| ---------------- | ----------------------------------------------- |
| Sort Array       | Required for two-pointer & duplicate skipping   |
| Outer Loop       | Fix `a = nums[i]`, skip if duplicate or `a > 0` |
| Inner Loop       | Use two pointers to find `b + c == -a`          |
| Avoid Duplicates | Skip duplicates for `a` and `b`                 |
| Append Result    | Save unique triplets to `res`                   |

