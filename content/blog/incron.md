---
title: "Incron"
date: 2023-07-12T08:58:30+03:00
draft: false
tags:
  - linux
  - reference
---

> [incron](http://inotify.aiken.cz/?section=incron&page=about&lang=en) is an "inotify cron" system. It consists of a daemon and a table manipulator. You can use it a similar way as the regular cron. The difference is that the inotify cron handles filesystem events rather than time periods.

I use it to auto-commit changes made to specific files to version control.

## Example usage

View incrontabs:

```
incrontab -l
```

Add/edit incrontabs:

```
incrontab -e
```

```
# commit and push when file is modified and saved
/home/wangonya/dir_to_track	IN_MODIFY	cd /home/wangonya/dir_to_track; git commit -a -m 'autocommit on change'; git push
```

[Archwiki](https://wiki.archlinux.org/title/Incron) for more info.
