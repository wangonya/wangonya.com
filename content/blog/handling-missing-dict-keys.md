---
author: Kelvin Wangonya
date: "2022-02-10T06:10:15+03:00"
title: Handling missing dict keys
tags:
  - python
---

Trying to access a non-existent key using this notation:
`dict[key]` raises a `KeyError`. An easy
workaround for this is to use `get(key)` instead, which
returns `None` if `key` isn\'t found, or
`get(key, default)` which returns `default` if
`key` isn\'t found.

```shell
>>> d = {}
>>> d
{}
>>> d["x"]
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
KeyError: 'x'
>>> d.get("x")
>>> d
{}
>>> d.get("x","y")
'y'
>>> d
{}
>>>
```

But what if you want to not only return a default if the key is missing,
but also assign that default to the specified key in the dict? There are
a couple of ways to do that.

## 1. `setdefault`

> `setdefault(key[, default])`
>
> If key is in the dictionary, return its value. If not, insert key with
> a value of default and return default. `default` defaults
> to `None`.
>
> [docs](https://docs.python.org/3/library/stdtypes.html#dict.setdefault)

```shell
>>> d
{}
>>> d.setdefault("x","y")
'y'
>>> d
{'x': 'y'}
>>>
```

## 2. `defaultdict`

> `class collections.defaultdict(default_factory=None, /[, ...])`
>
> The first argument provides the initial value for the
> [default~factory~](https://docs.python.org/3/library/collections.html#collections.defaultdict.default_factory)
> attribute; it defaults to `None`. All remaining arguments
> are treated the same as if they were passed to the
> [dict](https://docs.python.org/3/library/stdtypes.html#dict)
> constructor, including keyword arguments.
>
> [docs](https://docs.python.org/3/library/collections.html#collections.defaultdict)

```shell
>>> from collections import defaultdict
>>> i = defaultdict(int)
>>> i
defaultdict(<class 'int'>, {})
>>> i["x"]
0
>>> i
defaultdict(<class 'int'>, {'x': 0})
>>> i["y"]
0
>>> i
defaultdict(<class 'int'>, {'x': 0, 'y': 0})
>>>
```

Note that no `KeyError` is raised despite the key not
existing at first. Instead, the key is created with the default value of
the type passed into `defaultdict`, in this case
`0` for `int`. If we used `list`
instead, the default value would be `[]`, an so on.

Further reading:

- [defaultdict
  Examples](https://docs.python.org/3/library/collections.html#defaultdict-examples)
- [Use cases for the \'setdefault\' dict
  method](https://stackoverflow.com/questions/3483520/use-cases-for-the-setdefault-dict-method)
  (vs defaultdict)

## 3. Implement `__missing__`

This is actually what `defaultdict` does behind the scenes.
Use this when `defaultdict` doesn\'t fit your usecase.

```shell
>>> class M(dict):
...     def __missing__(self, key):
...             value = "my default value"
...             self[key] = value
...             return value
...
>>>
>>> m = M()
>>> m
{}
>>> m["x"]
'my default value'
>>> m
{'x': 'my default value'}
>>> m["y"]
'my default value'
>>> m
{'x': 'my default value', 'y': 'my default value'}
>>>
```
