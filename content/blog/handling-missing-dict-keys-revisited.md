---
author: Kelvin Wangonya
date: '2026-04-08T07:00:59+03:00'
title: 'Handling missing dict keys, revisited'
tags:
  - Python
---

A while back, I wrote a [post](/blog/handling-missing-dict-keys/) showing how to handle missing dict keys. In summary:
- Using `setdefault`
- Using `defaultdict`
- Implementing `__missing__`

While the advice there still stands, I've been reading through Fluent Python and came across two things worth being aware of when subclassing `dict`:

1. `__contains__` doesn't call `__missing__`, so `k in d` returns `False` for keys not yet set, even if `__missing__` would handle them  
2. `dict.get` doesn't call `__missing__`, so `.get(k)` won't use your custom default  

Both are by design and often what you want. `__missing__` is only used by `d[key]`, not other dict methods. Even `defaultdict` follows the same rule: only `d[key]` triggers the default factory and methods like `.get()` and `in` do not.

Overriding `get` is tempting but tricky. A naive implementation might look like this:

```py
class M(dict):
    def __missing__(self, key):
        value = "my default value"
        self[key] = value
        return value

    def get(self, key, default=None):
        try:
            return self[key]  # triggers __missing__ if key is absent
        except KeyError:
            return default    # dead code, never reached
```

This doesn't work. `self[key]` calls `__missing__` instead of raising `KeyError`, so the `except` block is unreachable and `default` is effectively useless:

```shell
>>> m = M()
>>> m.get("x", "fallback")
'my default value'          # expected "fallback", got __missing__ value instead
>>> m
{'x': 'my default value'}  # side effect: key was set in the dict
```

There's no clean way to honour both the `__missing__` default and the `default` parameter in `get` - they pull in opposite directions. It's better to accept that `get` and `__missing__` serve different purposes and leave `get` alone. Note that this applies to `dict` subclasses specifically. If you subclass `UserDict` instead, `get` calls `__getitem__` internally and `__missing__` is triggered. With a `dict` subclass, `get` is implemented in C and bypasses `__getitem__` altogether. The behaviour is inconsistent across the standard library depending on which base class you use.

As for `__contains__`, overriding it is a design decision. If your subclass provides a value for every possible key, returning `True` always can be reasonable:

```py
def __contains__(self, key):
    return True
```

Just be aware that `"x" in m` returns `True` even when `m` is `{}`, which looks confusing.
