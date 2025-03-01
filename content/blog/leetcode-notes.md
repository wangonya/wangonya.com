---
title: 'Leetcode Notes'
date: 2025-03-01T05:05:50+03:00
draft: true
---

## [Remove Element](https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/)
Remove `n` from array `nums` in-place.
```py
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

Naive:
- replace `n` with -1, keeping track of replacements in `vals`
- sort reverse true
- return `len(nums) - vals`

Better:
```py
k = 0
for i in range(len(nums)):
    if nums[i] != val:
        nums[k] = nums[i]
        k += 1
return k
```
