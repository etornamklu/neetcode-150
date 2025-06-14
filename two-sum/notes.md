# Two Sum

## Problem Type

* Arrays
* Hashing
* Two-pointer logic via complement lookup

---

## Trick

Use a **map** to store previously seen numbers and their indices. `[num: index]`
At each step, check if the *complement* (`target - current number`) is already in the map.

---

## Approach: HashMap-Based Solution (Go)

### Steps:

1. Create a map to store numbers we’ve seen and their corresponding **indices**.
2. Loop through the `nums` array:

   * Compute the **complement**: `target - num`.
   * If the complement exists in the map:

     * Return the **index of the complement** (from the map) and the current index.
   * Otherwise, store the current number and its index in the map.
3. If no pair is found, return `[-1, -1]`.

---

### Go Code:

```go
package twosum

func TwoSum(nums []int, target int) []int {
	count := make(map[int]int)

	for i, num := range nums {
		complememt := target - num

		if v, ok := count[complememt]; ok {
			return []int{v, i}
		}
		count[num] = i
	}
	return []int{-1, -1}
}
```

---

## Time Complexity: **O(n)**

* Single pass through the array: O(n)
* Map insertions and lookups are O(1) on average
* **Total time: O(n)**

---

## Space Complexity: **O(n)**

* In the worst case (no valid pair), map stores all `n` elements
* **Space used: O(n)**

---

## Summary

| Approach      | Time Complexity | Space Complexity | Notes                           |
| ------------- | --------------- | ---------------- | ------------------------------- |
| Hash Map (Go) | O(n)            | O(n)             | Fastest solution, returns early |

---

## Notes on Go `map`

* Declare a map using `make`:

  ```go
  m := make(map[KeyType]ValueType)
  ```

* In this case:

  ```go
  count := make(map[int]int)
  ```

  * Keys are `int` (the numbers in the array)
  * Values are `int` (the indices)
  * Used to quickly look up whether a complement has already been seen

