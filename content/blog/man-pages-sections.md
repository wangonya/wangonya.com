---
author: Kelvin Wangonya
date: 2024-09-25T05:47:44+03:00
title: man pages have sections!
tags:
  - til
---

I noticed this when reading through [Beej's Guide to C Programming](https://beej.us/guide/bgc/html/split/index.html). While explaining `printf`, reference is made to `man 3 printf` - which looked different from how I usually run `man`. Why the `3`?

So I ran `man printf` instead. Output:

```
PRINTF(1)                                       User Commands                                       PRINTF(1)

NAME
       printf - format and print data
 ...
```

Then `man 3 printf`:

```
printf(3)                                  Library Functions Manual                                 printf(3)

NAME
       printf,  fprintf,  dprintf, sprintf, snprintf, vprintf, vfprintf, vdprintf, vsprintf, vsnprintf - for‐
       matted output conversion

LIBRARY
       Standard C library (libc, -lc)

SYNOPSIS
       #include <stdio.h>
...
```

And I noticed what I'd been missing. The first output has `User Commands` at the top, while the other one has `Library Functions Manual`.

As it turns out, man pages are divided into sections. I've never thought to check `man man`, but it explains this.

```
MAN(1)                                        Manual pager utils                                       MAN(1)

NAME
       man - an interface to the system reference manuals

SYNOPSIS
       man [man options] [[section] page ...] ...

...

A manual page consists of several sections.

A section, if provided, will direct man to look only in that section of the manual.
The default action is to search in all of the available sections following a
pre-defined  order (see DEFAULTS), and to show only the first page found,
even if page exists in several sections.

    The table below shows the section numbers of the manual followed by the types of pages they contain.

    1   Executable programs or shell commands
    2   System calls (functions provided by the kernel)
    3   Library calls (functions within program libraries)
    4   Special files (usually found in /dev)
    5   File formats and conventions, e.g. /etc/passwd
    6   Games
    7   Miscellaneous (including macro packages and conventions), e.g. man(7), groff(7), man-pages(7)
    8   System administration commands (usually only for root)
    9   Kernel routines [Non standard]

```

To check what sections are available for a command:

```
man -f <command>

# for example: 
man -f printf
printf (1)           - format and print data
printf (1p)          - write formatted output
printf (3)           - formatted output conversion
printf (3p)          - print formatted output
```

You may need to run `sudo mandb` if you get the error `nothing appropriate` with the above command.