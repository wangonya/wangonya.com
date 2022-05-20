---
author: Kelvin Wangonya
date: "2021-10-15T16:09:08+03:00"
title: Setting up PostgreSQL on Arch Linux
tags:
  - reference
  - linux
  - db
  - mysql
---

[Arch Linux PostgreSQL
docs](https://wiki.archlinux.org/title/PostgreSQL)

## Install

```{.bash org-language="sh"}
$ sudo pacman -S postgresql
```

## Setup the data directory

```{.bash org-language="sh"}
$ sudo -iu postgres
[postgres]$ initdb -D /var/lib/postgres/data
```

## Start postgresql.service

```{.bash org-language="sh"}
$ systemctl start postgresql.service
$ systemctl enable postgresql.service
```

## Add user

    [postgres]$ createuser --interactive

## Create db

```{.bash org-language="sh"}
$ createdb myDatabaseName
```
