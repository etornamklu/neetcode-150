# Binary Search in 2D Matrix – Search Target in Sorted Matrix

## Problem Statement

Given a **2D matrix** where each row is sorted from **left to right** and each column is sorted from **top to bottom**, implement a search function that returns `true` if a given `target` exists in the matrix, otherwise returns `false`.

---

## Problem Type

* Matrices
* Binary Search (2D variant)
* Greedy
* Search Space Reduction
* Logarithmic-to-Linear Hybrid Time Complexity

---

## Trick

### 1. Start from **top-right** corner:

```go
r, c := 0, n-1  // where n = number of columns
```

### 2. While inside matrix bounds (`r < m && c >= 0`):

* Compare `matrix[r][c]` with `target`:

    * If **greater**, move **left**: `c--`
    * If **less**, move **down**: `r++`
    * If **equal**, return `true`

This exploits the sorted nature of the matrix:

* Moving **left** gives smaller values
* Moving **down** gives larger values

---

## Go Code

```go
package search_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1

	for r < m && c >= 0 {
		if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		} else {
			return true
		}
	}
	return false
}
```

---

## Example

```go
matrix := [][]int{
    {1,  4,  7, 11, 15},
    {2,  5,  8, 12, 19},
    {3,  6,  9, 16, 22},
    {10,13, 14,17, 24},
    {18,21, 23,26, 30},
}
target := 5
```

**Steps**:

| r | c | matrix\[r]\[c] | Action        |
| - | - | -------------- | ------------- |
| 0 | 4 | 15             | move left     |
| 0 | 3 | 11             | move left     |
| 0 | 2 | 7              | move left     |
| 0 | 1 | 4              | move down     |
| 1 | 1 | 5              | return true ✅ |

---

## Time and Space Complexity

| Operation      | Time     | Space |
| -------------- | -------- | ----- |
| Search         | O(m + n) | O(1)  |
| Result Storage | O(1)     |       |

*Worst-case: traverse one row and one column.*

---

## Go Feature Highlights

✅ Clean matrix traversal logic
✅ No extra memory allocation
✅ Index bounds strictly checked: `r < m` and `c >= 0`
✅ Smart greedy navigation using matrix properties

---

## Key Insights

✅ Matrix is sorted in both **rows and columns**
✅ Starting at **top-right** simplifies navigation
✅ Eliminates one row or one column each step
✅ Efficient for sparse matrices and avoids full scan

---

## Why Top-Right Works

* The top-right element is:

    * **Largest in its row**
    * **Smallest in its column**

Thus, it gives you **maximum information** about which direction to go.

---

## Summary Table

| Step                | Description                                           |
| ------------------- | ----------------------------------------------------- |
| Initialize Pointers | `r = 0`, `c = n - 1`                                  |
| While Loop          | Run while `r < m && c >= 0`                           |
| Compare Cell Value  | Navigate left or down based on comparison             |
| Return Result       | Return `true` if match, else loop until out of bounds |

