---
title: "Setting up PostgreSQL on Arch Linux"
author: Kinyanjui Wangonya
date: 2021-10-15T16:09:08+03:00
description: Setting up MariaDB on Arch Linux
tags:
  - reference
  - linux
  - psql
  - db
---

[Arch Linux PostgreSQL docs](https://wiki.archlinux.org/title/PostgreSQL)

## Install
```sh
$ sudo pacman -S postgresql
```

## Setup the data directory
```sh
$ sudo -iu postgres
[postgres]$ initdb -D /var/lib/postgres/data
```

## Start postgresql.service
```sh
$ systemctl start postgresql.service 
$ systemctl enable postgresql.service 
```

## Add user
```sh
[postgres]$ createuser --interactive
```

## Create db
```sh
$ createdb myDatabaseName
```
