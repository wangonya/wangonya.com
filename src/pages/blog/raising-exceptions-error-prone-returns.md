---
title: Raising exceptions to avoid error prone returns
date: "2020-02-11"
---

Consider a simple example function meant to divide two numbers. Here's one way to do it:

```python
def divide_numbers(a, b):
    try:
        return a / b
    except ZeroDivisionError:
        return None
```

At first glance, the code may look ok. The function can be used at another part of the code like so:

```python
x, y = 2, 0
result = divide_numbers(x, y)

# python programmers will usually do this instead of 
# explicitly checking for None
if not result: 
    print('Cannot divide by zero.') # this will run
else:
    print(f'result = {result}')
```

The problem is, what if we're dividing 0 by 2?

```python
x, y = 0, 2
result = divide_numbers(x, y)

if not result:
    print('Cannot divide by zero.') # this will still run
else:
    print(f'result = {result}')
```

Because `result = 0`, `if not result` evaluates to `True`, and `'Cannot divide by zero.'` is printed. Clearly wrong.

There's a couple of ways to deal with this. One way would be to explicitly check for `None`:

```python
x, y = 0, 2
result = divide_numbers(x, y)

if result is None:
    print('Cannot divide by zero.')
else:
    print(f'result = {result}') # this will run
```

Another way would be to split the return value into a two-tuple, and maintain the `if not` syntax:

```python
def divide_numbers(a, b):
    try:
        return True, a / b
    except ZeroDivisionError:
        return False, None

x, y = 0, 2
success, result = divide_numbers(x, y)

if not success:
    print('Cannot divide by zero.')
else:
    print(f'result = {result}') # this will run
```

This method introduces some unnecessary complexity. The tuple needs to be unpacked, and we now have to consider both `success` and `result` instead of just focusing on `result`.

The best way to get around all these issues would be to avoid returning `None` and raise an exception. Then, when calling the function, another `try, except` block is used to handle the exception if it occurs, otherwise, the return value is displayed.

```python
def divide_numbers(a, b):
    try:
        return a / b
    except ZeroDivisionError as e:
        raise ValueError('Invalid inputs') from e

x, y = 0, 2

try:
  result = divide_numbers(x, y)
except ValueError:
    print('Cannot divide by zero.')
else:
    print(f'result = {result}') # this will run
```

