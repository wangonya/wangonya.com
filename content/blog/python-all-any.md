---
title: "Analyzing python iterables with all() and any()"
date: 2019-11-01T12:32:12+03:00
tags: ["python"]
---

`all()` and `any()` are built-in functions that help analyze python iterables. 

## `all()`

`all()` returns `True` if **all** elements of the iterable are true (or if the iterable is empty).

```sh
Python 3.7.4

>>> x = [2, 3, 5, 1]
>>> all(x)
True

>>> x = [2, 3, 5, 0]
>>> all(x)
False

>>> x = []
>>> all(x)
True
```

In the second instance, `False` is returned because of the `0` in the list. Note that this would not be the case if the `0` was a string.

```sh
>>> x = [2, 3, 5, '0']
>>> all(x)
True
```

For checking dictionary values,

```sh
>>> x = {'item1': 'pen', 'item2': 'paper', 'item3': 'book'}
>>> all(x.values())
True

>>> x = {'item1': 'pen', 'item2': 'paper', 'item3': False}
>>> all(x.values())
False

>>> x = {}
>>> all(x)
True
```


## `any()`

`any()` returns `True` if **any** element of the iterable is true. If the iterable is empty, it returns `False`.

```sh
>>> x = [2, 3, 5, 1]
>>> any(x)
True

>>> x = [2, 3, 5, 0]
>>> any(x)
True

>>> x = [0, 0, 0, '0']
>>> any(x)
True

>>> x = [0, 0, 0, 0]
>>> any(x)
False

>>> x = []
>>> any(x)
False
```

It also works the same for dictionaries:

```sh
>>> x = {'item1': 'pen', 'item2': 'paper', 'item3': 'book'}
>>> any(x)
True

>>> x = {'item1': 'pen', 'item2': 'paper', 'item3': False}
>>> any(x)
True

>>> x = {}
>>> any(x)
False
```
