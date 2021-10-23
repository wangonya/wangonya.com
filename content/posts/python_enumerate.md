---
title: When to use python's enumerate() instead of range() in loops
author: Kelvin Wangonya
date: 2020-01-26
tags: 
  - python
---

The `range()` function is often useful when iterating over a set of integers:

```python
for n in range(50):
    ...

#

for n in range(10, 30):
    ...
```

or a list of strings:

```
for fruit in ["apple", "mango", "banana"]:
    ...
```

<!--more-->

Now, say you want to iterate over the list of fruits and also show the index of the current item in the list. Using `range()`, this might be done like this:

```python
fruits = ["apple", "mango", "banana"]

for i in range(len(fruits)):
    fruit = fruits[i]
    print(f"{i}: {fruit}")

# Output:
# 0: apple
# 1: mango
# 2: banana
```

It gets the job done, but not very pythonic. You have to get the length of the list to keep track of the index, then index into the array to get the current fruit - which makes the code more verbose and harder to read.

## A better way: `enumerate()`

```python
fruits = ["apple", "mango", "banana"]

for i, fruit in enumerate(fruits):
    print(f"{i}: {fruit}")

# Output:
# 0: apple
# 1: mango
# 2: banana
```

Numbering can also be set to begin at any desired number.

```python
fruits = ["apple", "mango", "banana"]

for i, fruit in enumerate(fruits, 7):
    print(f"{i}: {fruit}")

# Output
# 7: apple
# 8: mango
# 9: banana
```
