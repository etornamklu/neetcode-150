# Binary Search – Koko Eating Bananas 🍌

## Problem Statement

Koko loves to eat bananas. Given an array `piles` where `piles[i]` represents the number of bananas in the `i-th` pile, and an integer `h` representing the total number of hours Koko has to eat them, return the **minimum integer eating speed `k`** (bananas/hour) such that she can eat all bananas within `h` hours.

---

## Problem Type

* Arrays
* Binary Search on Answer
* Greedy
* Optimization
* Monotonic Predicate

---

## Trick

This is a **search in value space** (not index-based). The idea is:

### 1. Use Binary Search to find the **minimum feasible speed** `k`.

### 2. Key observations:

* Lower bound (`l`) = `1` banana/hour (slowest speed)
* Upper bound (`r`) = `max(piles)` (fastest speed needed)

### 3. Check mid-speed `m`:

* For each pile, Koko needs `ceil(pile / m)` hours.
* If total hours needed `> h`, she’s too slow → increase speed
* Else, try a slower speed → reduce right bound

---

## Go Code

```go
package koko_eating_bananas

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0

	for _, p := range piles {
		if p > r {
			r = p
		}
	}

	for r > l {
		m := l + (r-l)/2
		totalHours := 0

		for _, pile := range piles {
			totalHours += (pile + m - 1) / m // Equivalent to ceil(pile / m)
		}

		if totalHours > h {
			l = m + 1 // Need to eat faster
		} else {
			r = m // Can try slower speeds
		}
	}
	return l // The minimum speed that allows Koko to eat all bananas in h hours
}
```

---

## Example

```go
piles := []int{3, 6, 7, 11}
h := 8
```

**Search space**: `1` to `11`
Binary search midpoints try out speeds like `6`, `4`, `3`, etc.
At each speed, calculate total hours:

| Speed (k) | Hours Needed                    |
| --------- | ------------------------------- |
| 6         | 1 + 1 + 2 + 2 = 6 ✅             |
| 3         | 1 + 2 + 3 + 4 = 10 ❌ (too slow) |
| 4         | 1 + 2 + 2 + 3 = 8 ✅ (exact)     |

Final Answer: `4`

---

## Time and Space Complexity

| Operation         | Time           | Space |
| ----------------- | -------------- | ----- |
| Binary Search     | O(log M)       | O(1)  |
| Feasibility Check | O(N) per check |       |
| Total             | O(N log M)     | O(1)  |

Where:

* `N` = number of piles
* `M` = max pile size

---

## Go Feature Highlights

✅ Smart midpoint calc with `l + (r-l)/2`
✅ Integer ceiling without math lib: `(pile + m - 1) / m`
✅ Efficient, readable control flow
✅ Constant space, linear checks per iteration

---

## Key Insights

✅ Search is not over indices, but **possible speeds**
✅ The total hours needed is **monotonically decreasing** as speed increases
✅ So binary search can zero in on the minimum feasible speed
✅ Greedy check per speed ensures optimality

---

## Summary Table

| Step                | Description                             |
| ------------------- | --------------------------------------- |
| Init Bounds         | `l = 1`, `r = max(piles)`               |
| Midpoint            | Try `m := l + (r-l)/2` as eating speed  |
| Feasibility Check   | Compute hours with ceil(pile / m)       |
| Adjust Search Space | If too slow → `l = m + 1`, else `r = m` |
| Final Answer        | When `l == r`, return `l`               |


