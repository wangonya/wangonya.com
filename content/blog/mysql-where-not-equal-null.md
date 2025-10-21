---
title: 'Handling NULL values for "not equal" checks in SQL'
date: 2025-05-19T11:08:03+03:00
draft: false
tags:
  - mysql
  - til
---

This one bit me today. Apparently `NULL` values give unexpected results when checked using `!=` or `<>`.

For example:
```sql
id | model     | competitor_type
--------------------------------
1  | Ariya     | NULL
2  | Frontier  | NULL
3  | Rogue     | Competitor
4  | Sentra    | Non-competitor

SELECT * FROM table WHERE LOWER(competitor_type) != "non-competitor";
```

I expected to see ids 1, 2 and 3 as a result, but only got 3.

This makes sense when you think about it because comparisons can't be carried out against unavailable values. So if `NULL` values are expected in the result, they should explicitly be included in the query.

```sql
SELECT * FROM table WHERE (LOWER(competitor_type) != "non-competitor" OR competitor_type IS NULL);
```

This can also be handled by using `COALESCE` to get values from `NULL`, or by using the NULL-safe equal `<=>`, but I think it's a lot more clear to just check for `NULL` explicitly.