---
title: Getting the running Go version
author: Kelvin Wangonya
date: 2021-07-24T05:02:01+03:00
tags:
  - go
  - learning-go
series: Learning Go
---

This would mostly be useful for debugging.

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
