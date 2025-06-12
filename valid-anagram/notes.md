
# Valid Anagram

## Problem Type
- Strings
- Hashing / Frequency Counting


## Trick
Use a **fixed-size array** of length 26 to count character frequencies for lowercase letters ('a'–'z').  
- Increment for string `s`, decrement for string `t`.
- If both are anagrams, all counts should return to zero.

## Approach: Frequency Counter Using Array (Go)

### Steps:
1. If `s` and `t` are not the same length → return `false`.
2. Create an array of 26 integers (`tracker`) to count character frequencies.
3. Loop through string `s`, increment count at position `v - 'a'`.
4. Loop through string `t`, decrement count at position `v - 'a'`.
5. Loop through the `tracker` array:
   - If any value ≠ 0 → return `false` (not an anagram).
6. If all values are zero → return `true`.

### Go Code:
```go
package validanagram

func ValidAnagram(s string, t string) bool {
    // Strings must be same length to be anagrams
    if len(s) != len(t) {
        return false
    }

    // Create a tracker of 26 ints to count 'a' to 'z'
    tracker := [26]int{}
    //tracker := make([]int, 26)

    // Count characters in s
    for _, v := range s {
        tracker[v-'a']++
    }

    // Subtract characters in t
    for _, v := range t {
        tracker[v-'a']--
    }

    // Check if all counts are zero
    for _, v := range tracker {
        if v != 0 {
            return false
        }
    }

    return true
}
```

## Time Complexity: **O(n)**

* Each loop runs in **O(n)** where `n` is the length of the strings.
* Final check through the 26-length array is **O(1)**.
* **Total: O(n)**


## Space Complexity: **O(1)**

* Only one fixed-size array of 26 integers is used → **constant space**.
* Even if input size increases, space usage does not.


## Assumptions

* Input strings contain only lowercase English letters (`a` to `z`).
* If Unicode characters or uppercase letters are allowed, approach must be adjusted (e.g. using `map[rune]int`).


## 💡 Alternate Solution: HashMap

### Idea:

Use a `map[rune]int` to count character frequency for general Unicode strings.

### Time: **O(n)**

### Space: **O(k)** where `k` is number of unique characters

---

## Summary

| Approach         | Time Complexity | Space Complexity | Notes                              |
| ---------------- | --------------- | ---------------- | ---------------------------------- |
| Array (26 slots) | O(n)            | O(1)             | Best for lowercase English letters |
| HashMap          | O(n)            | O(k)             | Works with any Unicode characters  |


## Notes on Go Arrays

* Fixed-size arrays are declared like this:

  ```go
  var arr [26]int
  ```

  or:

  ```go
  arr := [26]int{}
  ```

* This creates an array of 26 `int`s (indexed from 0 to 25), initialized to zero.

* Arrays in Go have fixed length. For dynamic resizing, you'd use slices instead.

---

## Notes on Go Runes

* In Go, a `rune` is an alias for `int32` and represents a Unicode code point.
* When you iterate over a string using `for _, v := range s`, each `v` is a `rune`.

---

## Why `v - 'a'` Works

* Characters in Go are represented as numeric Unicode values.

* For lowercase English letters, `'a'` has a value of 97, `'b'` is 98, ..., `'z'` is 122.

* Subtracting `'a'` from any lowercase letter gives you an index between `0` and `25`:

  * `'b' - 'a'` = 98 - 97 = 1
  * `'z' - 'a'` = 122 - 97 = 25

* This allows us to directly map characters to array indices.


## Example:

```go
r := 'c'          // rune literal
i := r - 'a'      // i = 2
fmt.Println(i)    // Output: 2
```

