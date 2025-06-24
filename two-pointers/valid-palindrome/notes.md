# Valid Palindrome (Ignoring Non-Alphanumerics)

## Problem Statement

Given a string `s`, return `true` **if it is a palindrome**, considering **only alphanumeric characters** and **ignoring case**.

---

## Problem Type

* Two-Pointer Technique
* String Sanitization
* Case Normalization
* Palindrome Check

---

## Trick

Use two pointers (`l` and `r`) to scan from both ends of the string:

* Skip characters that are **not letters or digits**
* Convert letters to lowercase using `unicode.ToLower`
* Compare corresponding characters

---

## Intuition

Instead of creating a new cleaned string:

* Use **in-place comparison** with two pointers.
* Skip over non-alphanumeric characters using helper function.
* Normalize both characters before comparing.

This keeps time and space efficient.

---

## Go Code

```go
func ValidPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for r > l {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
```

---

## Step-by-Step Example

Input: `"A man, a plan, a canal: Panama"`

### Cleaned String (conceptually): `"amanaplanacanalpanama"`

Which **is a palindrome**.

### Trace Execution

| l   | r   | s\[l] | s\[r] | Normalized Equal? |
| --- | --- | ----- | ----- | ----------------- |
| 0   | 29  | `'A'` | `'a'` | ✅ `'a' == 'a'`    |
| 1   | 28  | `' '` |       | skip              |
| ... | ... | ...   | ...   | ...               |
| 10  | 19  | `'n'` | `'n'` | ✅                 |

Eventually all characters match → returns `true`.

---

## Time and Space Complexity

| Operation | Complexity | Notes                               |
| --------- | ---------- | ----------------------------------- |
| Time      | ✅ O(n)     | Each character visited at most once |
| Space     | ✅ O(1)     | No extra space (in-place check)     |

---

## Go Feature Highlights

* `rune(s[i])`: Converts byte to Unicode rune (for `unicode` package).
* `unicode.ToLower`: Properly lowers Unicode letters.
* `unicode.IsLetter`, `unicode.IsDigit`: Efficient alphanumeric checks.
* `for r > l`: Efficient in-place scanning with two pointers.

---

## Key Insights

✅ Skips punctuation, whitespace, and symbols
✅ Normalizes case before comparison
✅ No string reconstruction—uses **constant space**
✅ Works for full Unicode input due to `rune` and `unicode` usage

---

## Summary Table

| Step           | Description                              |
| -------------- | ---------------------------------------- |
| Pointer Setup  | `l = 0`, `r = len(s)-1`                  |
| Skip Non-Chars | Use `isAlphaNum` with `unicode` package  |
| Normalize Case | Use `unicode.ToLower()`                  |
| Compare        | Return `false` on mismatch, else move on |
| Return `true`  | If loop completes without mismatch       |

.
