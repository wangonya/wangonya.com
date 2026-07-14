---
author: Kelvin Wangonya
date: '2026-07-14T10:32:51+03:00'
title: 'BigQuery time travel is pretty neat'
draft: false
tags:
  - db
  - bigquery
---

```sql
SELECT *
FROM `mydataset.mytable`
  FOR SYSTEM_TIME AS OF TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 HOUR);
```

That query allows you to get data from the table as it was an hour ago.

In a setup where things like scheduled pipeline runs are constantly updating the table, it helps a lot to be able to see what a table looked like at a specific point in time.

The timestamp has to be within your time travel window. That's seven days by default, and can be configured from two to seven days.

For a full restore, copy the historical data into a new table. If you only need to repair the current table, copy the time travel data to a temp table, then update the current table using the temp table as source.

The temp table step is necessary for that second case because you cannot query from and write to the same table at different snapshot times. So doing:

```sql
INSERT INTO `mydataset.mytable`
SELECT *
FROM
  `mydataset.mytable`
    FOR SYSTEM_TIME AS OF TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 HOUR);
```

gives this error:
```
Table 'mydataset.mytable' is referenced with multiple snapshot read timestamps.
For example, if 'FOR SYSTEM_TIME AS OF' expressions are used on the table, they should use the same TIMESTAMP value.
Using the table as DML destination table references the table at query start time.
```
