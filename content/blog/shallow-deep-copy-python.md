---
author: Kelvin Wangonya
date: "2020-06-21T12:36:17+03:00"
title: Shallow & deep copying in python
tags:
  - python
---

What happens when a variable is assigned to another variable in python?
For example:

```{.bash org-language="sh"}
>>> x = 5
>>> y = x
```

Both `x` and `y` will have the value
`5`. But, when `x` was assigned to `y`,
`y` was not created as a completely new/separate object.
Instead, an **alias** for `x` was created. That is,
`y` points to the memory location of `x`. It does
not have it\'s own memory location - yet.

```{.bash org-language="sh"}
>>> id(x)
140428600776960
>>> id(y)
140428600776960
>>> x is y
True
```

You may never have any problems with this when working with [immutable
types](https://stackoverflow.com/a/23715872) because the alias is broken
as soon as either of the two variables change.

```{.bash org-language="sh"}
>>> x += 2
>>> x
7
>>> y
5
>>> id(x)
140539682924864
>>> id(y)
140428600776960
>>> x is y
False
```

But when working with mutable types, the alias is not broken when the
original is updated. This means changes in `x` would reflect
in `y`.

```{.bash org-language="sh"}
>>> x = [1,2,3]
>>> y = x
>>> x
[1, 2, 3]
>>> y
[1, 2, 3]
>>> x is y
True
>>>
>>> x.append(4)
>>> x.append(5)
>>> x
[1, 2, 3, 4, 5]
>>> y
[1, 2, 3, 4, 5]
>>> x is y
True
```

`y` was updated externally. This might be a cause of bugs if
for example a value is updated by an external library and other
variables are affected.

This can be prevented by creating shallow/deep copies of objects instead
of using assignment.

## Shallow copy

A shallow copy creates a new object, then populates it with references
of the _objects_ in the original object. Continuing with the previous
example, a shallow copy can be created using either the
`list()` or `copy()` command.

```{.bash org-language="sh"}
>>> z = list(x)
>>> z
[1, 2, 3, 4, 5]
```

Now if some more values are appended to `x`, `y`
will still be affected by `z` will not.

```{.bash org-language="sh"}
>>> x.append(6)
>>> x.append(7)
>>> x
[1, 2, 3, 4, 5, 6, 7]
>>> y
[1, 2, 3, 4, 5, 6, 7]
>>> z
[1, 2, 3, 4, 5]
```

However, a shallow copy doesn\'t fully solve the problem because even
though a new list was created, the objects in the list are still
references to the objects in `x`.

As it is currently, updating `x[0]` would not affect
`z[0]` because - immutable objects - the alias would be
broken. But, if we were dealing with a list of lists, an update in
`x[0]` would affect `z[0]`.

```{.bash org-language="sh"}
>>> x = [[1,2], [3,4]]
>>> y = x
>>> z = list(x)
>>> x
[[1, 2], [3, 4]]
>>> y
[[1, 2], [3, 4]]
>>> z
[[1, 2], [3, 4]]
>>> x.append([5,6])
>>> x
[[1, 2], [3, 4], [5, 6]]
>>> y
[[1, 2], [3, 4], [5, 6]]
>>> z
[[1, 2], [3, 4]]
>>>
>>> x[0][1] = 'edited'
>>> x
[[1, 'edited'], [3, 4], [5, 6]]
>>> y
[[1, 'edited'], [3, 4], [5, 6]]
>>> z
[[1, 'edited'], [3, 4]]
```

## Deep copy

A deep copy creates a new object, and completely new instances of the
objects in it. That is, a deep copied object is completely independent
of the original. Updating objects in the original would not affect the
deep copied object since there\'s no longer any connection.

```{.bash org-language="sh"}
>>> import copy
>>> x = [[1,2], [3,4]]
>>> z = copy.deepcopy(x)
>>> x
[[1, 2], [3, 4]]
>>> z
[[1, 2], [3, 4]]
>>> x is z
False
>>> x[0][1] = 'edited'
>>> x
[[1, 'edited'], [3, 4]]
>>> z
[[1, 2], [3, 4]]
```
