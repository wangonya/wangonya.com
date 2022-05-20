---
author: Kelvin Wangonya
date: "2021-11-22T06:58:33+03:00"
title: Running multiple programs in parallel from a bash script
tags:
  - reference
  - bash
---

This comes in handy when I want to, for example:

- run an MySql proxy to Cloudql or similar, then
- run a project server (fastapi, django etc) that depends on the
  database connection

From [this
answer](https://stackoverflow.com/questions/3004811/how-do-you-run-multiple-programs-in-parallel-from-a-bash-script):

```shell
set -m  # enable job control

prog1 & prog2 && fg
```

This will:

- Start prog1.
- Send it to background, but keep printing its output.
- Start prog2, and keep it in foreground, so you can close it with
  ctrl-c.
- When you close prog2, you\'ll return to prog1\'s foreground, so you
  can also close it with ctrl-c.
