# Group Anagrams

## Problem Type

* Strings
* Hashing
* Frequency Counting

---

## Trick

Use a **fixed-size array** `[26]int` to count character frequencies for each word.
Two words are anagrams if their frequency counts match.
Use this array as a key in a map:
`[char_freq] → list of anagrams`

---

## Approach: Frequency Map as Key (Go)

### Steps:

1. Create a map with key type `[26]int` and value type `[]string`.
2. For each string in the input list:

   * Create a `[26]int` array representing character counts.
   * For each character in the string:

     * Increment the corresponding position in the array.
   * Use the array as a key in the map, and append the string to the corresponding group.
3. Convert the map values into a 2D string slice and return it.

---

### Go Code:

```go
package groupanagrams

func GroupAnagrams(strs []string) [][]string {
	track := make(map[[26]int][]string)

	for _, str := range strs {
		key := [26]int{}

		for _, char := range str {
			key[char-'a']++
		}

		track[key] = append(track[key], str)
	}

	results := [][]string{}
	for _, v := range track {
		results = append(results, v)
	}

	return results
}
```

---

## Time Complexity: **O(n \* k)**

* `n` = number of strings
* `k` = average length of each string
* Creating and comparing the key (frequency array) takes O(k)
* Total time: **O(n \* k)**

---

## Space Complexity: **O(n \* k)**

* Map stores up to `n` keys, each associated with one or more strings
* Each key is a `[26]int` array
* Each string is stored in a group, total storage is O(n \* k)

---

## Summary

| Approach              | Time Complexity | Space Complexity | Notes                             |
| --------------------- | --------------- | ---------------- | --------------------------------- |
| Char Frequency as Key | O(n \* k)       | O(n \* k)        | Faster than sorting; Go map trick |

---

## Notes on Go `map`

* Declare a map using `make`:

  ```go
  m := make(map[KeyType]ValueType)
  ```

* In this case:

  ```go
  track := make(map[[26]int][]string)
  ```

  * Keys are `[26]int` — fixed-size arrays (representing letter counts)
  * Values are `[]string` — groups of anagrams
  * Go supports using arrays as map keys (by value comparison)


### Time Complexity of the **Sorting-Based** Approach for Group Anagrams:

---

## 🔍 Sorting-Based Strategy

For each string:

1. Sort the characters → `O(k log k)` where `k = len(str)`
2. Use the sorted string as a key in a map to group anagrams → O(1) average time

---

## 🔢 Time Complexity: **O(n \* k log k)**

| Term | Meaning                               |
| ---- | ------------------------------------- |
| `n`  | Number of strings in the input `strs` |
| `k`  | Maximum length of a single string     |

We sort each of the `n` strings of up to `k` characters →
**`O(n * k log k)` total time**

---

## 📦 Space Complexity: **O(n \* k)**

* In the worst case:

  * You store all strings in groups → O(n \* k)
  * Plus the sorted versions as map keys → O(n \* k)

---

## 🧠 Summary Table

| Approach             | Time Complexity | Space Complexity | Notes                        |
| -------------------- | --------------- | ---------------- | ---------------------------- |
| Char Freq (Go array) | O(n \* k)       | O(n \* k)        | Fastest; avoids sorting      |
| Sorting-based        | O(n \* k log k) | O(n \* k)        | Simpler; uses sorted strings |

---

Let me know if you want the code + notes in your preferred format for the sorting-based version!
