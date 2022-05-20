---
author: Kelvin Wangonya
date: "2021-07-24T05:02:01+03:00"
title: Getting the running Go version
tags:
  - go
---

```go
package main

import (
    "log"
    "runtime"
)

func main() {
    log.Printf("Running Go: %s", runtime.Version())
}
```

Output: `Running Go: go1.16.6`
