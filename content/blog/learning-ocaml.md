---
title: "Learning OCaml"
date: 2022-10-12T06:34:01+03:00
tags:
  - ocaml
draft: true
---

- [OCaml website](https://ocaml.org/)
- [Why OCaml?](https://ocaml.org/about)
- [OCaml for the Skeptical](https://www2.lib.uchicago.edu/keith/ocaml-class/why.html)
- [Why OCaml? , Why not OCaml?](https://www.imsc.res.in/~rsidd/ocaml.html)

Below are a few notes I jotted down while learning Ocaml.

## Setup (Arch linux)

```sh
yay -S opam

opam init
eval $(opam env)
opam install dune merlin ocaml-lsp-server odoc ocamlformat utop
```

`utop` is a better repl (has history navigation, auto-completion, etc) compared to the default `ocaml`.

```sh
~ > ocaml
OCaml version 4.14.0
Enter #help;; for help.

#
```

```sh
~ > utop                                                                                           
───────────────────┬──────────────────────────────────────────────────────────────┬────────────────────
                   │ Welcome to utop version 2.10.0 (using OCaml version 4.14.0)! │                    
                   └──────────────────────────────────────────────────────────────┘                    

Type #utop_help for help about using utop.

─( 06:46:02 )─< command 0 >─────────────────────────────────────────────────────────────{ counter: 0 }─
utop # 
┌───┬─────┬───────────┬──────────────┬──────┬────────┬────┬──────┬─────┬───────────┬────────┬─────────┐
│Arg│Array│ArrayLabels│Assert_failure│Atomic│Bigarray│Bool│Buffer│Bytes│BytesLabels│Callback│Camlinter│
└───┴─────┴───────────┴──────────────┴──────┴────────┴────┴──────┴─────┴───────────┴────────┴─────────┘
```

## Basics

Expressions can be evaluated directly in the repl. Each expression ends with a `;;`

```ocaml
# 2 * 2;;
- : int = 4
```

### Variables

"Variables" are declared using the `let` keyword. These are not variables in the traditional sense as their values can not be changed once initialized.

```ocaml
let x = 2;;
val x : int = 2

# x * x;;
- : int = 4
```

The above can also be written like this:

```ocaml
# let x = 2 in x * x;;
- : int = 4
```

### Functions

Functions are declared in a similar way:

```ocaml
# let square n = n * n;;
val square : int -> int = <fun>

# square 4;;
- : int = 16
```

Recursive functions require the `rec` keyword:

```ocaml
# let rec factorial n = if n = 1 then 1 else n * factorial (n-1);;
val factorial : int -> int = <fun>

# factorial 4;;
- : int = 24

# factorial 6;;
- : int = 720
```

### Pattern matching

The `if then else` format can be replaced by pattern matching:

```ocaml
# let rec factorial n = match n with 1 -> 1 | _ -> n * factorial (n-1);;
val factorial : int -> int = <fun>

# factorial 4;;
- : int = 24

# factorial 6;;
- : int = 720
```

## Lists
