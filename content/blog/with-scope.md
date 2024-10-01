---
author: Kelvin Wangonya
date: 2024-10-01T16:27:39+03:00
title: with statements don't create scope
tags:
  - til
  - python
---

I saw some code today that initialized a variable inside `with` and continued to use it outside the block.
I've always assumed `with` statements had their own scope so this had me a little confused. Turns out this is perfectly fine.

> 

> A `with` statement does not create a scope (like `if`, `for` and `while` do not create a scope either).
> 
> As a result, Python will analyze the code and see that you made an assignment in the `with` statement, and thus that will make the variable local (to the real scope).
> 
> In Python variables do not need initialization in all code paths: as a programmer, you are responsible to make sure that a variable is assigned before it is used. This can result in shorter code: say for instance you know for sure that a list contains at least one element, then you can assign in a `for` loop. In Java assignment in a `for` loop is not considered safe (since it is possible that the body of the loop is never executed).
> 
> Initialization before the `with` scope can be safer in the sense that after the `with` statement we can safely assume that the variable exists. If on the other hand the variable should be assigned in the `with` statement, not initializing it before the `with` statement actually results in an additional check: Python will error if somehow the assignment was skipped in the `with` statement.
> 
> A `with` statement is only used for context management purposes. It forces (by syntax) that the context you open in the `with` is closed at the end of the indentation.
>
> [source](https://stackoverflow.com/a/45100308/9312256)

