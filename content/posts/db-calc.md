---
title: "You can use your database as a simple calculator"
author: Kinyanjui Wangonya
date: 2021-10-23T17:49:32+03:00
description:
tags:
  - db
---

Saw this in the [MySQL docs](https://dev.mysql.com/doc/refman/8.0/en/entering-queries.html). I'm assuming this only works for relational databases.

I don't know if it's very useful but definitely interesting.

```sql
MariaDB[] > SELECT SIN(PI()/4), (4+1)*5;
+------------------+---------+
| SIN(PI()/4)      | (4+1)*5 |
+------------------+---------+
| 0.70710678118655 |      25 |
+------------------+---------+
1 row in set (0.02 sec)
```

```sql
postgres=# SELECT SIN(PI()/4), (4+1)*5;
        sin         | ?column?
--------------------+----------
 0.7071067811865475 |       25
(1 row)
```