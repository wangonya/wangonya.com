---
title: "Grokking Algorithms Notes [WIP]"
date: 2022-05-21T12:03:23+03:00
tags:
  - algorithms
  - book notes
  - python
---

Book: [https://www.manning.com/books/grokking-algorithms](https://www.manning.com/books/grokking-algorithms)

# Binary search
Use if the input is sorted.

If the input consists of numbers, for example, half the numbers are eliminated in each step with binary search. When searching for a number between 1 and 100, a maximum of 7 steps are needed to get the result.

```
100 -> 50 -> 25 -> 13 -> 7 -> 4 -> 2 -> 1
```

In general, for any list of _n_, binary search will take log<sub>2</sub>_n_ steps to run in the worst case.

## Implementing binary search in python

```python
def binary_search(_list, item):
    high = len(_list) - 1
    low = 0

    while high >= low:
        mid = (high + low) // 2
        guess = _list[mid]
        if guess == item:
            return mid
        elif guess > item:
            high = mid - 1
        else:
            low = mid + 1

    return None


if __name__ == "__main__":
    position = binary_search([2, 3, 5, 7, 8, 9, 12, 24, 32], 12)
    print(position)
```

## Implementing binary search in go

```go
// TODO
```

# Arrays & linked lists

## Arrays

> Rendom access. Fast reads, slow writes.

With arrays, data is stored contiguously (right next to each other) in memory. If more data needs to be added and there's not enough space for it to be stored contiguosly in the current location, everything needs to be moved to a different location that will allow this to be done. Inserting data into the middle of an array also requires moving  everything else to the right to make space for the new item. Adding data can therefore be quite slow. While it's possible to reserve slots as a workaround for having to change locations, the reserved slots end up being wasted in case they're not used. Deleting is also slow as everything needs to be moved up to fill the empty space (assuming the deletion was done in the middle).

## Linked lists

> Sequential access. Fast writes, slow reads.

Data can be saved anywhere in memory. Each item stores the address of the next item in the list. This solves the problem of needing to store everything together, but is slow when reading data. To read the last item in the list, all previous items will have to be accessed. Adding and deleting is fast as it only requires a change in the next address pointed to.

# Selection sort
