---
title: "Variables & Mutability"
date: 2020-06-14T22:33:13+03:00
description: Understanding variables and mutability in rust
tags:
  - rust
---

By default, variables are immutable in Rust. Coming from a Python background, I have to keep in mind that this:

```rust
fn main() {
    let var = 10;
    println!("var = {}", var);  // var = 10
    var = 5;  // try and reassign var
    println!("var is now {}", var);
}
```

... will not compile. The compiler output:


```sh
error[E0384]: cannot assign twice to immutable variable `var`
  |
2 |     let var = 10;
  |         ---
  |         |
  |         first assignment to `var`
  |         help: make this binding mutable: `mut var`
3 |     println!("var = {}", var);  // var = 10
4 |     var = 5;  // try and reassign var
  |     ^^^^^^^ cannot assign twice to immutable variable
```
