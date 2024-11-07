---
author: Kelvin Wangonya
date: "2024-11-07T09:00:30+03:00"
title: Be careful with join type typos
tags:
  - sql
  - til
---

I noticed a typo in one of my sql queries today, but the funny thing is the query still worked fine.
It looked something like this:

```sql {hl_lines=[5]}
SELECT
    *
FROM
    table_1
LEFFT JOIN
    table_2
USING (a,b);
```

At first I thought BigQuery was pretty smart and just assumed I meant `LEFT` instead.
Maybe it's a common typo and there's an alias for it. So I got curious and tried out this:

```sql {hl_lines=[5]}
SELECT
    *
FROM
    table_1
LEFFFFTFT JOIN
    table_2
USING (a,b);
```

No error. That can't be right.

That's when it hit me. This query isn't doing a `LEFT JOIN`. It's using the incorrectly spelled
name as an alias for `table_1` and doing an `INNER JOIN` instead. Basically this:

```sql
SELECT
    *
FROM
    table_1 AS LEFFT
JOIN
    table_2
ON LEFFT.a = table_2.a AND LEFFT.b = table_2.b
```

So the query works fine but the results would be incorrect.

This is one merit of syntax highlighting I've never really thought about until now. The only reason I noticed this quickly was by
seeing `LEFFT` wasn't highlighted as I expected. The results of the query looked fine at first glance so this would have definitely bit
me somewhere down the line if it got to production.
