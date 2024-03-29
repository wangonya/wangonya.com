---
author: Kelvin Wangonya
date: "2021-10-15T15:56:36+03:00"
title: Setting up MariaDB on Arch Linux
tags:
  - reference
  - linux
  - db
  - mysql
---

[Arch Linux MariaDB docs](https://wiki.archlinux.org/title/MariaDB)

> MariaDB is a reliable, high performance and full-featured database
> server which aims to be an \'always Free, backward compatible,
> drop-in\' replacement of MySQL. Since 2013 MariaDB is Arch Linux\'s
> default implementation of MySQL.

## Install

```shell
$ sudo pacman -S mariadb
```

## Setup the data directory

```shell
$ sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

## Start mariadb.service

```shell
$ systemctl start mariadb.service
$ systemctl enable mariadb.service
```

## Add user

> In the below example, the user monty with some~pass~ as password is
> being created, then granted full permissions to the database mydb:

```shell
$ sudo mysql -u root -p

MariaDB> CREATE USER 'monty'@'localhost' IDENTIFIED BY 'some_pass';
MariaDB> GRANT ALL PRIVILEGES ON mydb.* TO 'monty'@'localhost';
MariaDB> FLUSH PRIVILEGES;
MariaDB> quit
```

## Logging into MariaDB

```shell
mysql -u //user_name// -p -h //ip_address// //db_name//
```

Eg:

```shell
mysql -u username -ppassword -h localhost database_name
```
