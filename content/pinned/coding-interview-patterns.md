---
title: Coding interview question patterns
author: Kelvin Wangonya
pinned: true
last_updated: 2021-11-18
tags:
  - cheatsheet
  - reference
---

| Problem                           | Category              | Notes                                                                           |
| --------------------------------- | --------------------- | ------------------------------------------------------------------------------- |
| [Two Sum]                         | lists, dicts          | track `num:nums[i]` in dict, check if `target - nums[i]` exists                 |
| [Best Time to Buy and Sell Stock] | lists, dp             | [Kadane\'s Algorithm]                                                           |
| [Contains Duplicate]              | lists, dicts, sorting | list -> set and compare lengths / loop while adding to dict and check existence |
| [CyclicRotation]                  | lists                 | loop while `insert(0, pop())` / use [`colections.deque.rotate`]                 |
| [OddOccurrencesInArray]           | lists                 | use [`Counter`] + generator expression to get the key with an odd value         |
| [FrogJmp]                         | -                     | -                                                                               |

[two sum]: https://leetcode.com/problems/two-sum/
[best time to buy and sell stock]: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
[kadane\'s algorithm]: https://wangonya.com/blog/kadanes-algorithm/
[contains duplicate]: https://leetcode.com/problems/contains-duplicate/
[cyclicrotation]: https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/
[`colections.deque.rotate`]: https://docs.python.org/3/library/collections.html#collections.deque.rotate
[oddoccurrencesinarray]: https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/
[`counter`]: https://docs.python.org/3/library/collections.html#collections.Counter
[frogjmp]: https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
