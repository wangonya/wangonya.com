---
title: "mysql-dump specific tables"
date: 2022-08-13T15:19:50+03:00
tags:
  - mysql
  - database
  - reference
---

If you only need to dump one table:

```sh
mysqldump -u username -ppassword db_name table_name > dump.sql
```

For multiple tables, separate the table names with a whitespace:

```sh
mysqldump -u username -ppassword db_name table_name_1 table_name_2 > dump.sql
```

To only get tables with a certain prefix, e.g `pref_`:

```sh
mysqldump db_name $(mysql -D db_name -Bse "show tables like 'pref\_%'") > dump.sql
# flags: -B batch, -s silent mode, -e execute
```

To only dump rows that meet a specific criteria:

```sh
mysqldump -u username -ppassword db_name table_name --where="where clause here" > dump.sql
```

Importing the dump file:

```sh
mysql -u username -ppassword db_name < dump.sql
```
