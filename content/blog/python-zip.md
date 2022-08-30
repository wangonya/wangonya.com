---
author: Kelvin Wangonya
date: "2021-11-03T06:13:46+03:00"
title: Processing iterators in parallel
tags:
  - python
---

Use [zip](https://docs.python.org/3/library/functions.html#zip)

> `zip(*iterables, strict=False)`
>
> Iterate over several iterables in parallel, producing tuples with an
> item from each one.

## Examples

### Without zip

```python
countries = ["Kenya", "Tanzania", "Uganda"]
capitals = ["Nairobi", "Dodoma", "Kampala"]

print("=== using range ===")
for i in range(len(countries)):
    country, capital = countries[i], capitals[i]
    print(f"Country: {country}, Capital: {capital}")

print("=== using enumerate ===")
for i, country in enumerate(countries):
    country, capital = countries[i], capitals[i]
    print(f"Country: {country}, Capital: {capital}")
```

Output:

```shell
=== using range ===
Country: Kenya, Capital: Nairobi
Country: Tanzania, Capital: Dodoma
Country: Uganda, Capital: Kampala

=== using enumerate ===
Country: Kenya, Capital: Nairobi
Country: Tanzania, Capital: Dodoma
Country: Uganda, Capital: Kampala
```

### With zip

```python
print("=== using zip ===")
for country, capital in zip(countries, capitals):
    print(f"Country: {country}, Capital: {capital}")
```

Output:

```shell
=== using zip ===
Country: Kenya, Capital: Nairobi
Country: Tanzania, Capital: Dodoma
Country: Uganda, Capital: Kampala
```

Same result, but more elegant.
