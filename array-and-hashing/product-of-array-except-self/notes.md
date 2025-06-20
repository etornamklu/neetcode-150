# Product of Array Except Self

## Problem Statement

Given an integer array `nums`, return an array `result` such that `result[i]` is the **product of all the elements of `nums` except `nums[i]`**.

**Do not use division**
**Solve it in O(n) time**
**Use constant space (excluding the output array)**

---

## 🔍 Problem Type

* Arrays
* Prefix and Postfix Products
* Space Optimization
* In-place Computation

---

## Trick

Instead of dividing the product of the whole array by `nums[i]`, we compute two passes:

1. **Prefix Pass** – For each index `i`, store the product of elements **before** `i`.
2. **Postfix Pass** – Multiply by the product of elements **after** `i`.

---

## Approach: Prefix and Postfix Products

### Intuition

* For an index `i`, the product of all elements **except** `nums[i]` is:
  `product = product_of_elements_before_i * product_of_elements_after_i`

### Prefix Pass

1. Create an output array `result` and initialize with 1.
2. Traverse from left to right.
3. At each index `i`, store the product of all elements before it.

```go
prefix := 1
for i := range nums {
	result[i] = prefix
	prefix *= nums[i]
}
```

### Postfix Pass

1. Initialize `postfix` as 1.
2. Traverse from right to left.
3. Multiply `result[i]` (which holds prefix product) by `postfix` product.

```go
postfix := 1
for i := len(nums) - 1; i >= 0; i-- {
	result[i] *= postfix
	postfix *= nums[i]
}
```

---

## Go Code

```go
package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = 1
	}

	prefix := 1
	for i := range nums {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
```

---

## Walkthrough Example

Given: `nums = [1, 2, 3, 4]`

### Prefix Pass

| i | prefix | result\[i] | new prefix (`*= nums[i]`) |
| - | ------ | ---------- | ------------------------- |
| 0 | 1      | 1          | 1 × 1 = 1                 |
| 1 | 1      | 1          | 1 × 2 = 2                 |
| 2 | 2      | 2          | 2 × 3 = 6                 |
| 3 | 6      | 6          | 6 × 4 = 24                |

Result after prefix: `[1, 1, 2, 6]`

### Postfix Pass

| i | postfix | result\[i] | result\[i] × postfix | new postfix (`*= nums[i]`) |
| - | ------- | ---------- | -------------------- | -------------------------- |
| 3 | 1       | 6          | 6                    | 1 × 4 = 4                  |
| 2 | 4       | 2          | 8                    | 4 × 3 = 12                 |
| 1 | 12      | 1          | 12                   | 12 × 2 = 24                |
| 0 | 24      | 1          | 24                   | 24 × 1 = 24                |

Final Result: `[24, 12, 8, 6]`

---

## Why No Division?

Division is **not allowed** due to:

* Potential **division by zero**
* Precision issues (in float/int mismatch)
* Interview constraint

Instead, prefix/postfix multiplication gives a **mathematically clean and safe** way to get the result.

---

## Time and Space Complexity

| Operation    | Time Complexity | Space Complexity | Notes                              |
| ------------ | --------------- | ---------------- | ---------------------------------- |
| Prefix Loop  | O(n)            | O(1)             | O(1) extra, excluding result array |
| Postfix Loop | O(n)            | O(1)             | No extra array needed              |
| Overall      | ✅ **O(n)**      | ✅ **O(1)**       | Efficient and optimal              |

---

## Summary

| Step         | Description                                       |
| ------------ | ------------------------------------------------- |
| Prefix Pass  | Fill result with product of elements before index |
| Postfix Pass | Multiply result with product of elements after    |
| Final Output | Each index `i` holds product of all `nums[j≠i]`   |

---

## Key Insights

* The **result array reuses space** by holding prefix first, then being multiplied by postfix.
* This is a classic example of how **multiple passes and state reuse** can solve problems without extra memory.
* Very common interview problem—**master this pattern**!


## `for range` in Go

The `for range` loop is Go’s **idiomatic way to iterate** over:

* **Arrays**
* **Slices**
* **Strings**
* **Maps**
* **Channels**

---

## Basic Syntax

```go
for index, value := range collection {
    // use index and value
}
```

* `index`: current position in the collection (0-based for slices/arrays).
* `value`: item at that position.

---

## Slice Example

```go
nums := []int{10, 20, 30}

for i, val := range nums {
    fmt.Println(i, val)
}
// Output:
// 0 10
// 1 20
// 2 30
```

---

## Ignoring Index or Value

Use `_` (blank identifier) to ignore what you don’t need:

```go
for i := range nums {       // only index
    fmt.Println(i)
}

for _, val := range nums {  // only value
    fmt.Println(val)
}
```

---

## Behind the Scenes

Under the hood:

```go
for index := 0; index < len(nums); index++ {
    value := nums[index]
    ...
}
```

`for range` is just syntactic sugar for this.

---

## Important Notes

| Behavior                                  | Example                                |
| ----------------------------------------- | -------------------------------------- |
| Loop variable **copies** the value        | Modifying `val` won’t change `nums[i]` |
| Value is read-only unless using pointer   | Use `&nums[i]` if you want to mutate   |
| For **maps**, order is **not guaranteed** | Random access order in `for range map` |

---

## Value vs Pointer

```go
for _, val := range nums {
    val = 100 // does NOT change original slice
}
```

But:

```go
for i := range nums {
    nums[i] = 100 // changes original slice
}
```

---

## Use Cases

| Type    | Loop Pattern                  | Notes                          |
| ------- | ----------------------------- | ------------------------------ |
| Slice   | `for i, val := range slice`   | Common for indexing or value   |
| Array   | Same as slice                 |                                |
| String  | Iterates runes (Unicode-safe) | Not just bytes                 |
| Map     | `for key, val := range map`   | Order not guaranteed           |
| Channel | `for val := range ch`         | Blocks until channel is closed |

---

## Summary Table

| Feature                | `for range`                            |
| ---------------------- | -------------------------------------- |
| Read value and index   | `for i, v := range list`               |
| Ignore index or value  | Use `_` in place of unused variable    |
| Value is copied        | Changes to `val` don’t affect original |
| Efficient & idiomatic  | Preferred loop style in Go             |
| Works on maps/channels | With key/value or values respectively  |
