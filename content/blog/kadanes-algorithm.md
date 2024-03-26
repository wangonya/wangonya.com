---
author: Kelvin Wangonya
date: "2021-12-02T08:25:49+03:00"
title: Kadane's Algorithm
tags:
  - python
  - algorithms
---

Possible usecase: [Maximum subarray
problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem).

## Steps

1.  Initialize two variables: `max_sum = current_sum = 0`
2.  Loop through numbers in list
    - Set `current_sum = max(0, current_sum + list[i])`
    - Set `max_sum = max(max_sum, current_sum)`
3.  Return `max_sum`

## Caveats [^1], [^2]

- Kadane\'s Algorithm requires at least one positive number, so an
  input of all negative numbers would be invalid.
- If the array contains all non-negative numbers, then the problem is
  trivial; a maximum subarray is the entire array.
- If the array contains all non-positive numbers, then a solution is
  any subarray of size 1 containing the maximal value of the array (or
  the empty subarray (which has sum 0), if it is permitted).
- Several different sub-arrays may have the same maximum sum.

### Solving [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) with Kadane\'s Algorithm

Slightly tweak the algorithm to track max profit and min price.

```python
def maxProfit(self, prices: List[int]) -> int:
    max_profit, min_price = 0, float("inf")
    for price in prices:
        min_price = min(min_price, price)
        max_profit = max(max_profit, price - min_price)
    return max_profit
```

[^1]:
    [Wikipedia: Maximum subarray
    problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem)

[^2]:
    [SO: Kadane Algorithm Negative
    Numbers](https://stackoverflow.com/questions/9942228/kadane-algorithm-negative-numbers)
