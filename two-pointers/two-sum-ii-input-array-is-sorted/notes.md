# Two Sum II – Input Array Is Sorted

## Problem Statement

Given a **1-indexed** array of integers `numbers` that is **sorted in non-decreasing order**, find two numbers such that they **add up to a specific target number**.

Return the **indices (1-based)** of the two numbers.

**Constraints:**

* The input is **sorted**
* There is **exactly one solution**
* Use **constant extra space**

---

## 🔍 Problem Type

* Two Pointers
* Sorted Array Optimization
* Binary Search Alternative
* Array Traversal

---

## Trick

Use a **left** and **right** pointer to scan from both ends of the array:

* If `sum > target`, move **right** pointer left (decrease pointer).
* If `sum < target`, move **left** pointer right (increase pointer).
* If `sum == target`, return the indices (1-based).

---

## Intuition

Since the array is **sorted**, we can exploit this order:

* A **larger sum** means the **right number is too big** → move right
* A **smaller sum** means the **left number is too small** → move left
* This avoids unnecessary checks and gives **O(n)** time.

---

## Go Code

```go
package twosumiiinputarrayissorted

func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for r > l {
		if numbers[r]+numbers[l] > target {
			r--
		} else if numbers[r]+numbers[l] < target {
			l++
		} else {
			return []int{l + 1, r + 1} // 1-indexed return
		}
	}

	return []int{-1, -1} // failsafe, not expected to hit
}
```

---

## Walkthrough Example

Input:
`numbers = [2, 7, 11, 15]`
`target = 9`

### Steps

| l | r | numbers\[l] | numbers\[r] | sum | Action           |
| - | - | ----------- | ----------- | --- | ---------------- |
| 0 | 3 | 2           | 15          | 17  | r--              |
| 0 | 2 | 2           | 11          | 13  | r--              |
| 0 | 1 | 2           | 7           | 9   | ✅ return \[1, 2] |

---

## Time and Space Complexity

| Operation        | Time   | Space  |
| ---------------- | ------ | ------ |
| Two-pointer loop | ✅ O(n) | ✅ O(1) |

Efficient use of **sorted property** avoids nested loops or hash maps.

---

## Go Feature Highlights

* Slices: `numbers` is a slice passed by reference.
* No need for extra space like hash maps.
* Returns a **slice** with 1-indexed positions.

---

## Key Insights

✅ Takes advantage of the sorted array
✅ Constant space, single pass
✅ Avoids brute force `O(n^2)`
✅ Ideal example of **greedy + two pointer** strategy

---

## Summary Table

| Step            | Description                                   |
| --------------- | --------------------------------------------- |
| Initialize      | `l = 0`, `r = len(numbers)-1`                 |
| Compare Sum     | Check `numbers[l] + numbers[r]`               |
| Adjust Pointers | Move inward depending on how sum compares     |
| Return Indices  | If match found, return `l+1, r+1` (1-indexed) |


