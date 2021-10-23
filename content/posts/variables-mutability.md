---
title: "Variables & mutability in rust"
author: Kelvin Wangonya
date: 2020-06-14T22:33:13+03:00
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

... will not compile. 

<!--more-->

The rust compiler gives a clear output of what went wrong:


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

Once a value is assigned to `var`, `var` is bound to that value until it goes out of scope.

## Opting out of immutability

Variables can be made mutable using the `mut` keyword:

```rust
fn main() {
    let mut var = 10;
    println!("var = {}", var);  // var = 10
    var = 5;  // reassign var
    println!("var is now {}", var);  // var is now 5
}
```

This makes it clear when reading the code that `var` expects to be changed somewhere else.

## Shadowing

Instead of making a variable mutable and reassigning it directly, rust provides another technique known as shadowing. Here's an example:

```rust
fn main() {
    let var = 10;  // not mutable
    println!("var = {}", var);  // var = 10
    let var = 5;  // shadow first var (10)
    println!("var is now {}", var);  // var is now 5
    let var = var + 2;  // shadow second var (5) and add 2 to it
    println!("var + 2 = {}", var);  // var + 2 = 7
}
```

`var` remains immutable all through so trying to attach any value to it without using `let` to tell Rust you're shadowing will not work.

Shadowing also allows changing the variable data type - something that making a variable mutable doesn't allow. An example of when to use this would be when writing a program that accepts a string and returns the number of characters in that string.

This works:

```rust
let hello = "hello";
let hello = hello.len();  // 5
```

This will not compile:

```rust
let mut hello = "hello";
hello = hello.len();
```

Output:

```sh
error[E0308]: mismatched types
  |
3 |     hello = hello.len();
  |             ^^^^^^^^^^^ expected `&str`, found `usize`
```

## Variables vs constants

- Constants are declared with `const` keyword, variables with `let`
- Constants cannot be made mutable
- Constants can only be assigned to constant expressions, not to values that are computed during runtime - like results of function calls

