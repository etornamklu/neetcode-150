# Container With Most Water – Maximize Water Area Between Lines

## Problem Statement

Given an **integer array** `height` where each element represents the height of a vertical line at that index on the x-axis, find two lines such that:

* The container formed between them holds the **most water**.
* Return the **maximum amount of water** a container can store.

---

## Problem Type

* Arrays
* Two Pointers
* Greedy Strategy
* Optimization Pattern

---

## Trick

### 1. **Use two pointers**

Start with:

```
l = 0            (left end)
r = len(height)-1 (right end)
```

### 2. **Move the shorter line inward**

Why?

* The area is limited by the **shorter** height.
* Moving the taller line **doesn't help** — only a **potentially taller** shorter line can improve the area.

---

## Intuition

To **maximize area**, we need:

```
Area = min(height[l], height[r]) * (r - l)
```

The **width** shrinks as we move pointers, so we want to **maximize height** while decreasing width.

Strategy:

* Start with the widest container.
* At each step, shrink the width by moving the pointer pointing to the shorter line.
* Keep track of the **maximum area** encountered.

---

## Go Code

```go
package countainer_w_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for r > l {
		area := min(height[l], height[r]) * (r - l)

		if area > res {
			res = area
		}
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

---

## Walkthrough Example

Input:

```go
height := []int{1,8,6,2,5,4,8,3,7}
```

### Step-by-step:

| l | r | height\[l] | height\[r] | width | Area | Max Area |
| - | - | ---------- | ---------- | ----- | ---- | -------- |
| 0 | 8 | 1          | 7          | 8     | 8    | 8        |
| 1 | 8 | 8          | 7          | 7     | 49   | 49 ✅     |
| 1 | 7 | 8          | 3          | 6     | 18   | 49       |
| 1 | 6 | 8          | 8          | 5     | 40   | 49       |
| 1 | 5 | 8          | 4          | 4     | 16   | 49       |
| 1 | 4 | 8          | 5          | 3     | 15   | 49       |
| 1 | 3 | 8          | 2          | 2     | 4    | 49       |
| 1 | 2 | 8          | 6          | 1     | 6    | 49       |

Final Answer: `49`

---

## Time and Space Complexity

| Operation      | Time | Space |
| -------------- | ---- | ----- |
| Loop           | O(n) | O(1)  |
| Result storage | —    | O(1)  |

**Overall Time:** ✅ `O(n)`
**Overall Space:** ✅ Constant

---

## Go Feature Highlights

* Two-pointer strategy for optimization
* `min(a, b)` helper avoids repetitive logic
* No need to sort the input array
* Pure greedy scanning — no nested loops

---

## Key Insights

✅ Two-pointer method efficiently searches entire space
✅ Always move pointer at shorter height
✅ Linear scan avoids brute-force O(n²)
✅ Classic optimization pattern — **greedy shrink-and-compare**

---

## Summary Table

| Step              | Description                            |
| ----------------- | -------------------------------------- |
| Two Pointers Init | l = 0, r = len(height)-1               |
| Area Calculation  | `min(height[l], height[r]) * (r - l)`  |
| Move Pointers     | Move pointer at **shorter height**     |
| Track Max Area    | Update `res` if current area is larger |
| End Condition     | Loop until `l >= r`                    |

