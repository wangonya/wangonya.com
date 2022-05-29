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

```go
// TODO: Go implementation
```

# Arrays & linked lists

## Arrays

> Random access. Fast reads, slow writes.

With arrays, data is stored contiguously (right next to each other) in memory. If more data needs to be added and there's not enough space for it to be stored contiguously in the current location, everything needs to be moved to a different location that will allow this to be done. Inserting data into the middle of an array also requires moving everything else to the right to make space for the new item. Adding data can therefore be quite slow. While it's possible to reserve slots as a workaround for having to change locations, the reserved slots end up being wasted in case they're not used. Deleting is also slow as everything needs to be moved up to fill the empty space (assuming the deletion was done in the middle).

## Linked lists

> Sequential access. Fast writes, slow reads.

Data can be saved anywhere in memory. Each item stores the address of the next item in the list. This solves the problem of needing to store everything together, but is slow when reading data. To read the last item in the list, all previous items will have to be accessed. Adding and deleting is fast as it only requires a change in the next address pointed to.

# Recursion

> Loops may achieve a performance gain for your program. Recursion may achieve a performance gain for your programmer. Choose which is more important in your situation! - [https://stackoverflow.com/a/72694/9312256](https://stackoverflow.com/a/72694/9312256)

Every recursive function has two parts: the base case, and the recursive case.

> **Tip** 💡
>
> When you’re writing a recursive function involving an array, the base case is often an empty array or an array with one element.

> **Exercise**
>
> Find the total of a list of numbers using recursion.

```python
def _sum(_list):
    if not _list:
        return 0

    if len(_list) == 1:
        return _list[0]
    else:
        return _list.pop() + _sum(_list)

print(_sum([1,2,3,4,5])) # => 15
print(_sum([2,4,6])) # => 12
```

## The stack

Items are added at the top (push) and removed from the top (pop). FILO.

### The call stack with recursion

Example using factorial calculation:

Code:

```python
def fact(x):
  if x == 1:
    return 1
  else:
    return x * fact(x-1)

fact(3)
```

Call stack:

| Code                  | Call stack                                                        |
| --------------------- | ----------------------------------------------------------------- |
| fact(3)               | **fact(x=3)**                                                     |
| if x == 1:            | **fact(x=3)**                                                     |
| else:                 | **fact(x=3)**                                                     |
| return x \* fact(x-1) | **fact(x=2)**<br/>fact(x=3)                                       |
| if x == 1:            | **fact(x=2)**<br/>fact(x=3)                                       |
| else:                 | **fact(x=2)**<br/>fact(x=3)                                       |
| return x \* fact(x-1) | **fact(x=1)**<br/>fact(x=2)<br/>fact(x=3)                         |
| if x == 1:            | **fact(x=1)**<br/>fact(x=2)<br/>fact(x=3)                         |
| return 1              | **fact(x=1)** <= popped off the stack<br/>fact(x=2)<br/>fact(x=3) |
| return x \* fact(x-1) | **fact(x=2)** <= popped off the stack<br/>fact(x=3)               |
| return x \* fact(x-1) | **fact(x=3)** <= popped off the stack                             |

Each call to `fact` has its own copy of `x`. You can’t access a different function’s copy of `x`.

# Divide & conquer

D&C algorithms are recursive algorithms.

To solve a problem using D&C, there are two steps:

1. Figure out the base case. This should be the simplest possible case.
2. Divide or decrease your problem until it becomes the base case.

> **Exercise**
>
> Suppose you’re a farmer with a plot of land - 1680 \* 640 meters. You want to divide this farm evenly into square plots. You want the plots to be as big as possible. How do you figure out the largest square size you can use for a plot of land?

```python
def divide_land(a, b):
    bigger_side = max(a,b)
    smaller_side = min(a,b)

    if smaller_side * 2 == bigger_side:
        return (smaller_side, smaller_side)
    else:
        rem = bigger_side % smaller_side
        return divide_land(rem, smaller_side)

print(divide_land(1680, 640)) # => (80, 80)
print(divide_land(400, 640)) # => (80, 80)
```

# Quicksort

Uses divide & conquer.

## Sorting an array of numbers with quick sort

Base case: empty array or array with one element.

For bigger arrays, choose an element from the array (pivot) and find elements smaller that the pivot and larger than it. All smaller elements go to the left of the pivot. Bigger elements go to the right of the pivot. Recurse until array is sorted.

```python
def quicksort(array):
    if len(array) < 2:
        return array
    pivot = array[0]
    less = [i for i in array[1:] if i < pivot]
    equal = [i for i in array[1:] if i == pivot]
    greater = [i for i in array[1:] if i > pivot]

    return quicksort(less) + equal + [pivot] + quicksort(greater)


print(quicksort([10, 2, 3, 7, 12, 2, 5, 1, 3])) # => [1, 2, 2, 3, 3, 5, 7, 10, 12]
```

The performance of quicksort heavily depends on the pivot choosen. If the first element is the pivot and quicksort is called with an array that is already sorted it doesn’t check to see whether the input array is already sorted. So it will still try to sort it.

Middle pivots work better. In the worst case (first item as pivot), the stack size is O(_n_). In the best case (middle pivot), the stack size is O(log _n_).

The overall complexity of quicksort is O(_n_ log _n_)

# Graphs

A graph models a set of connections / how different things are connected to one another. It's made up of nodes and edges. A node can be directly connected to many other nodes. Those nodes are called its neighbors.

## Breadth-first search

Used to find the shortest path to a target. It can help answer two types of questions:

- Is there a path from node A to node B?
- What is the shortest path from node A to node B?

> **Example**
>
> Suppose you’re the proud owner of a mango farm. You’re looking for a mango seller who can sell your mangoes. Are you connected to a mango seller on Facebook? Well, you can search through your friends.
>
> First, make a list of friends to search. Now, go to each person in the list and check whether that person sells mangoes. Suppose none of your friends are mango sellers. Now you have to search through your friends’ friends. Each time you search for someone from the list, add all of their friends to the list. This way, you not only search your friends, but you search their friends too.
>
> First should be searched first before friends of friends. To keep this order, you need a queue.

```python
from collections import deque

graph = dict()
graph["you"] = ["alice", "bob", "claire"]
graph["bob"] = ["anuj", "peggy"]
graph["alice"] = ["peggy"]
graph["claire"] = ["tom mango seller", "jonny"]
graph["anuj"] = []
graph["peggy"] = []
graph["tom mango seller"] = []
graph["jonny"] = []


def search(name):
    search_queue = deque()
    search_queue += graph[name]
    searched = []

    while search_queue:
        person = search_queue.popleft()
        if not person in searched:
            if "mango seller" in person:
                return person
            else:
                search_queue += graph[person]
                searched.append(person)
    return False

print(search("you")) # => tom mango seller
```

Running time: O(V+E) (V for number of vertices, E for number of edges).

> **Tip** 💡
>
> If you have a problem like “find the shortest X,” try modeling your problem as a graph, and use breadth-first search to solve.

## Dijkstra’s algorithm
