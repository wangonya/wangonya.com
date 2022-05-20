---
author: Kelvin Wangonya
date: 2019-09-22
title: difflib - Finding close matches of strings from a list
tags:
  - python
---

Say we have a list of strings: `_list = [...,]` and user
input `_input = '...'`, how do we find the items in
`_list` that most closely resemble `_input`?

Python has a built-in package called `difflib` with the
function `get_close_matches()` that can help us.

`get_close_matches(word, possibilities, n, cutoff)` accepts
four parameters:

- `word` - the word to find close matches for in our list
- `possibilities` - the list in which to search for close
  matches of `word`
- `n` (optional) - the maximum number of close matches to
  return. Must be `> 0`. Default is `3`.
- `cutoff` (optional) - a float in the range \[0, 1\] that
  a `possibility` must score in order to be considered
  similar to `word`. `0` is very lenient,
  `1` is very strict. Default is `0.6`.

An example from the docs:

```shell
Python 3.7.3

>>> from difflib import get_close_matches
>>> get_close_matches('appel', ['ape', 'apple', 'peach', 'puppy'])
['apple', 'ape']
```

The above example can easily be modified to use a custom list
`_list` for `possibilities` and user input
`_input` for `word`.
