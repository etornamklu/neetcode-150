# Encode and Decode Strings

Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings

## Problem Type

* Strings
* Delimiter encoding
* Custom serialization

---

## Trick

**Prefix each string with its length and a special delimiter (like `#`)**.
This helps in decoding because we can parse the length, then slice the string accordingly.

Example:
List `["hello", "world"]` becomes `"5#hello5#world"`.

---

## Approach: Length-Delimited Encoding

### `Encode` Method

1. Loop through each string.
2. Convert its length to a string using `strconv.Itoa`.
3. Append the format `length#string` to the result.
4. Return the final encoded string.

### `Decode` Method

1. Use two pointers to extract the length before each `#`.
2. Use `strconv.Atoi` to convert that length to an int.
3. Slice the encoded string using the decoded length.
4. Repeat until the entire string is parsed.

---

## Go Code

```go

import (
	"strconv"
	"strings"
)

type Solution struct{}

// Encode converts a slice of strings into a single string
func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
        // Prefix each string with its length followed by a '#'
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()
}

// Decode parses the encoded string back into the original slice of strings
func (s *Solution) Decode(str string) []string {
	var result []string
	i := 0

	for i < len(str) {
		j := i
		// Move j forward to find the position of the delimiter '#'
		for str[j] != '#' {
			j++
		}

		// Extract substring from i to j (length prefix), and convert to int
		length, _ := strconv.Atoi(str[i:j])

		// Extract the original string using the length
		word := str[j+1 : j+1+length]
		result = append(result, word)

		// Move i to the start of the next encoded part
		i = j + 1 + length
	}

	return result
}
```

---

## Explanation of `strconv` Methods

### `strconv.Itoa(int) string`

* Converts an `int` to its string representation.
* Example: `strconv.Itoa(5)` → `"5"`
* Used in `Encode` to convert string length to a string before concatenation.

### `strconv.Atoi(string) (int, error)`

* Converts a numeric string to an integer.
* Example: `strconv.Atoi("12")` → `12`
* Used in `Decode` to extract the length of the next word.

---

## Time and Space Complexity

| Operation | Time Complexity | Space Complexity | Notes                           |
| --------- | --------------- | ---------------- | ------------------------------- |
| `Encode`  | O(n)            | O(1)             | n = total length of all strings |
| `Decode`  | O(n)            | O(n)             | allocates result slice          |

---

## Summary

| Method   | Purpose                             | Key Step                         |
| -------- | ----------------------------------- | -------------------------------- |
| `Encode` | Convert strings to a single string  | Use `strconv.Itoa` and delimiter |
| `Decode` | Reconstruct strings from one string | Use `strconv.Atoi` and slicing   |


## What is `strings.Builder`?

`strings.Builder` is a type from Go's `strings` package that efficiently builds strings using a mutable buffer, avoiding repeated allocations.

### Why Use It?

* Concatenating strings with `+` or `+=` creates **new copies** each time.
* `strings.Builder` **accumulates bytes efficiently**, then converts them to a string **in one allocation**.

---

## Basic Usage Example

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteByte(',')
	builder.WriteString(" world!")
	
	result := builder.String()
	fmt.Println(result) // Output: Hello, world!
}
```

### Explanation:

* `WriteString("Hello")` adds `"Hello"` to the buffer.
* `WriteByte(',')` adds a single character.
* `WriteString(" world!")` appends more text.
* `builder.String()` gives the final string.

---

## 🔧 Common Methods

| Method                  | Description                                |
| ----------------------- | ------------------------------------------ |
| `WriteString(s string)` | Appends a string to the builder            |
| `WriteByte(b byte)`     | Appends a single byte (character)          |
| `WriteRune(r rune)`     | Appends a Unicode rune (e.g., emoji, char) |
| `String()`              | Returns the final string                   |
| `Reset()`               | Clears the builder to reuse it             |
| `Len()`                 | Returns the number of bytes written so far |
| `Cap()`                 | Returns the current capacity of the buffer |

---

## Builder vs `+` Example

```go
func withPlus(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str + " "
	}
	return result
}

func withBuilder(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(str)
		b.WriteByte(' ')
	}
	return b.String()
}
```

### ⚠ Performance Difference:

* `withPlus` causes **many intermediate strings** to be created and discarded.
* `withBuilder` grows an internal buffer and avoids unnecessary copies.

---

## Unicode-safe Example with `WriteRune`

```go
var b strings.Builder
b.WriteRune('😊')     // emoji rune
b.WriteString(" cool")
fmt.Println(b.String())  // Output: 😊 cool
```

---

## Summary

| Feature                    | `+` / `+=`           | `strings.Builder` |
| -------------------------- | -------------------- | ----------------- |
| Memory Efficient           | ❌ (many allocations) | ✅ (reuses buffer) |
| Easy to Use                | ✅                    | ✅                 |
| Best for Loops             | ❌                    | ✅                 |
| Unicode Safe (`WriteRune`) | ❌                    | ✅                 |


