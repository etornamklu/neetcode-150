# Longest Consecutive Sequence

## Problem Statement

Given an **unsorted array of integers**, return the **length of the longest consecutive sequence** of elements.

* You **must** write an algorithm that runs in **O(n)** time.

---

## 🔍 Problem Type

* Arrays
* HashSet (map in Go)
* Sequence Detection
* Greedy / Linear Scan

---

## Trick

Use a **set** (map in Go) to:

1. Store all numbers for **O(1) lookups**.
2. Check for sequences, a new sequence starts when `num - 1` isn't present
3. use a loop to count consecutive elements in the sequence 
4. update the max sequence count if current sequence count is higher

---

## Go Code

```go
func LongestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			length := 1
			for {
				if _, exists := numSet[num+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
```

---

## Step-by-Step Walkthrough

Input: `nums = [100, 4, 200, 1, 3, 2]`

### Step 1: Build Set

```go
numSet := map[int]bool{
    100: true, 4: true, 200: true,
    1: true, 3: true, 2: true,
}
```

### Step 2: Loop over each number

For each `num` in the set:

* Only start a sequence if `num - 1` **does not** exist.
* This ensures you only start **once per sequence**.

Example:

* `1` has no `0` → start new sequence → `1, 2, 3, 4` → length = 4
* `3` has `2` → skip (not the start)
* `4` has `3` → skip
* `100` has no `99` → start → length = 1
* `200` has no `199` → start → length = 1

### Final Result: `4`

---

## ⏱ Why is this O(n) Time?

| Step       | Operation                                  | Time Complexity |
| ---------- | ------------------------------------------ | --------------- |
| Build set  | `for _, num := range`                      | O(n)            |
| Outer loop | `for num := range set`                     | O(n)            |
| Inner loop | Each `num + i` scanned **once** across all | O(n) total      |

Even though there’s a **nested loop**, the inner loop runs **only for unique sequence starts**, and **each element is visited at most once**, so total work is `O(n)`.

This is **amortized linear time**.

---

## Space Complexity

| Structure           | Space |
| ------------------- | ----- |
| `map[int]bool`      | O(n)  |
| `longest`, `length` | O(1)  |

---

## Summary Table

| Step               | Description                                 |
| ------------------ | ------------------------------------------- |
| Hashing            | Store all nums in map for quick lookups     |
| Sequence Detection | Start count if previous number is missing   |
| Expand Sequence    | Keep checking `num+1`, `num+2`, ...         |
| Track Max          | Update `longest` if current sequence longer |

---

## Key Insights

Only start a sequence if `num - 1` doesn’t exist
Avoids repeated work by never starting from a middle
Each element is visited **at most once**
`map[int]bool` enables **O(1)** presence checks

---

## Go Feature Highlights

* `range map` iterates over **keys only** by default.
* Using `map[int]bool` is an idiomatic way to build a **set** in Go.

