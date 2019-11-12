---
title: "Documenting python code using docstrings"
date: 2019-11-12T09:56:54+03:00
description: "How to write simple documentation for your code using python docstrings."
tags: ["python"]
---

Until recently, I thought python docstrings were just an alternative way to write comments in the code. As I have since learnt, they can be quite useful in writing documentation right into the code. Here's what I mean:

```python3
# hello.py

class Hello:
    def __init__(self):
        """This is the Hello class init method. It doesn't really
        do anything in this code. I just included it here so I can write this
        long multi-line docstring."""
        pass

    def hello():
        'Simply prints hello world!'
        print("Hello World!")
```

Above is a simple python script - `hello.py`. To see the documentation for that class, run the script in the interactive shell:

```sh
$ python3 -i hello.py

>>>
```

Then, type `help(Hello)`:

```sh
>>> help(Hello)
```

A neatly formatted documentation for the class should be returned.

```
class Hello(builtins.object)
|  Methods defined here:
|
|  __init__(self)
|      This is the Hello class init method. It doesn't really
|      do anything in this code. I just included it here so I can write this
|      long multi-line docstring.
|
|  hello()
|      Simply prints hello world!
|
|  ----------------------------------------------------------------------
```
