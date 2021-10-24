---
title: "Getting the Python version at runtime"
author: Kelvin Wangonya
date: 2021-10-24T17:42:19+03:00
description:
tags:
  - python
---

[Again](https://wangonya.com/blog/getting-the-running-go-version/), might come in handy while debugging.

```py
import sys

print(sys.version)
```

Output:
```sh
3.9.7 (default, Aug 31 2021, 13:28:12) 
[GCC 11.1.0]
```