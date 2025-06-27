# Binary Search – Find Target Index in Sorted Array

## Problem Statement

Given a **sorted integer array** `nums`, implement binary search to return the **index** of a given `target`. If `target` is not found, return `-1`.

---

## Problem Type

* Arrays
* Binary Search
* Divide and Conquer
* Logarithmic Time Complexity

---

## Trick

### 1. Use **two pointers**: `l` (left) and `r` (right)

### 2. Compute **midpoint** safely:

```go
m := l + (r - l) / 2
```

(Avoids overflow)

### 3. Compare `nums[m]` with `target`:

* If equal → return `m`
* If less  → search right half (`l = m + 1`)
* If more  → search left half (`r = m - 1`)

---

## Go Code

```go
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1

    for l <= r {
        m := l + (r - l) / 2

        if target > nums[m] {
            l = m + 1
        } else if target < nums[m] {
            r = m - 1
        } else {
            return m
        }
    }

    return -1
}
```

---

## Example

```go
nums := []int{1, 3, 5, 7, 9}
target := 5
```

Steps:

| l | r | m | nums\[m] | Action     |
| - | - | - | -------- | ---------- |
| 0 | 4 | 2 | 5        | return 2 ✅ |

---

## Time and Space Complexity

| Operation      | Time     | Space |
| -------------- | -------- | ----- |
| Search         | O(log n) | O(1)  |
| Result Storage | O(1)     |       |

---

## Go Feature Highlights

✅ Efficient integer midpoint with `l + (r - l)/2`
✅ Clean, readable binary search logic
✅ No extra memory allocation
✅ Guarded while-loop condition with `l <= r`

---

## Key Insights

✅ Works on sorted arrays
✅ Reduces search space by half each step
✅ Stops early if match is found
✅ Constant space, logarithmic time

---

## Why `l < r` Fails for `nums = [5], target = 5`

Let’s analyze:

### Input:

```go
nums := []int{5}
target := 5
```

### With condition: `for l < r`

| l | r | Condition | Enters loop? |
| - | - | --------- | ------------ |
| 0 | 0 | `0 < 0`   | ❌ No         |

**So the loop never runs**, and the function returns `-1`, even though the element exists.

---

### Why `l <= r` is Correct

| l | r | `l <= r` | Mid | nums\[m] | Action     |
| - | - | -------- | --- | -------- | ---------- |
| 0 | 0 | ✅        | 0   | 5        | return 0 ✅ |

✅ The loop runs even when `l == r`, allowing single-element arrays to be checked.

---

## Summary Table

| Step                 | Description                                      |
| -------------------- | ------------------------------------------------ |
| Initialize Pointers  | `l = 0`, `r = len(nums) - 1`                     |
| Midpoint Calculation | Safe formula: `l + (r - l)/2`                    |
| While Loop           | Use `l <= r` to include single-element subarrays |
| Compare Midpoint     | Binary chop space using comparisons              |
| Return Result        | Return index if match, else `-1`                 |

