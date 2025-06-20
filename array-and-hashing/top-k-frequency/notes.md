
# Top K Frequent Elements
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

## Problem Type

* Arrays
* Hashing
* Bucket Sort
* Heap (alternative)

---

## Trick

Use a **bucket sort strategy** instead of a heap to achieve linear time.
The trick is to reverse-map frequencies to values:
Build an array `freq[i]` that holds all numbers that appear exactly `i` times.
The return the top `k` elements.

---

## Approach: Bucket Sort with Frequency Map (Go)

### Steps:

1. **Count Frequencies**
   Use a map `count[num] = frequency` to tally how often each number appears.

2. **Group by Frequency**
   Use a slice of slices `freq[i]` to group numbers by how many times they occur.
   `freq[3]` holds all numbers that appear exactly 3 times.

3. **Collect Top K**
   Starting from the end of `freq` (i.e. highest frequencies), collect numbers until you have `k` of them.

---

### Go Code:

```go
func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)              // num -> frequency
    freq := make([][]int, len(nums)+1)      // freq[i] = list of nums that appear i times

    for _, num := range nums {
        count[num]++
    }

    for num, c := range count {
        freq[c] = append(freq[c], num)
    }

    res := make([]int, 0, k) // pre-allocate with capacity k
    for i := len(freq) - 1; i >= 0 && len(res) < k; i-- {
        res = append(res, freq[i]...)
    }

    return res[:k] // ensure only k elements in case we appended more
}
```

---

## Time Complexity: **O(n)**

* Counting frequencies: O(n)
* Bucketing into `freq`: O(n)
* Collecting top K: O(n) in worst case
* **Total: O(n)**

> Note: This beats heap-based solutions which are O(n log k)

---

## Space Complexity: **O(n)**

* Map to store frequencies: O(n)
* Bucket array of size `n+1`: O(n)
* Result slice: O(k)
* **Total space: O(n)**

---

## Summary

| Approach    | Time Complexity | Space Complexity | Notes                                |
| ----------- | --------------- | ---------------- | ------------------------------------ |
| Bucket Sort | O(n)            | O(n)             | Fastest for unsorted input, no heap  |
| Heap (alt)  | O(n log k)      | O(n)             | Slower, but supports streaming input |

---

## Notes on Go `map` and slices

* Frequency map:

  ```go
  count := make(map[int]int)
  ```

  * Keys are `int` (the numbers)
  * Values are `int` (how many times they appear)

* Bucket array of slices:

  ```go
  freq := make([][]int, len(nums)+1)
  ```

  * `freq[i]` contains all numbers that appeared `i` times
  * Max frequency is at most `len(nums)`, hence size `n + 1`

