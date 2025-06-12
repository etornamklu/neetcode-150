# Contains Duplicates

## Problem Type
- Arrays
- Hashing


## Trick
Use a **map** to keep track of numbers we've seen before.  
If a number appears again, it’s a duplicate.


## Approach: HashMap-Based Solution (Go)

### Steps:
1. Create a map to store seen numbers.
2. Loop through each number in the array:
   - If the number is already in the map → return `true` (duplicate found).
   - Otherwise, add the number to the map.
3. If loop finishes with no duplicates → return `false`.

### Go Code:

```go
func ContainsDuplicate(nums []int) bool {
    // Create a map to track seen numbers.
    // In Go, maps are created using the make function.
    // `map[int]bool` means keys are integers, and values are booleans (used here as flags).
    tracker := make(map[int]bool)

    for _, num := range nums {
        // Check if the number is already in the map.
        if _, ok := tracker[num]; ok {
            return true
        }
        // Mark the number as seen.
        tracker[num] = true
    }

    return false
}
```

## Time Complexity: **O(n)**

* Looping through the slice takes **O(n)**.
* Each map lookup and insertion is **O(1)** on average.
* **Total time: O(n)**.

## Space Complexity: **O(n)**

* In the worst case (no duplicates), the map stores all `n` elements.
* **Space used: O(n)**.

## Summary

| Approach         | Time Complexity | Space Complexity | Notes                    |
| ---------------- | --------------- | ---------------- | ------------------------ |
| Hash Map (Go)    | O(n)            | O(n)             | Fastest for large inputs |

## Notes on Go `map`

* Declare a map using `make`:

  ```go
  m := make(map[KeyType]ValueType)
  ```

* Example:

  ```go
  tracker := make(map[int]bool)
  ```

* This creates a map where:

  * Keys are `int`
  * Values are `bool`
  * Useful for presence checks like a set.
