---
author: Kelvin Wangonya
date: "2022-04-29T11:02:01+03:00"
title: Checking string contains substring
tags:
  - go
---

From [this answer](https://stackoverflow.com/a/61842548/9312256):

```go
import (
    "fmt"
    "regexp"
    "strings"
)

const (
    str    = "something"
    substr = "some"
)

// 1. Contains
res := strings.Contains(str, substr)
fmt.Println(res) // true

// 2. Index: check the index of the first instance of substr in str, or -1 if substr is not present
i := strings.Index(str, substr)
fmt.Println(i) // 0

// 3. Split by substr and check len of the slice, or length is 1 if substr is not present
ss := strings.Split(str, substr)
fmt.Println(len(ss)) // 2

// 4. Check number of non-overlapping instances of substr in str
c := strings.Count(str, substr)
fmt.Println(c) // 1

// 5. RegExp
matched, _ := regexp.MatchString(substr, str)
fmt.Println(matched) // true

// 6. Compiled RegExp
re = regexp.MustCompile(substr)
res = re.MatchString(str)
fmt.Println(res) // true
```

Option two easily solves [this
problem](https://leetcode.com/problems/implement-strstr/).
