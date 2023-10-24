---
title: "Trying out Ubuntu 23.04"
date: 2023-09-07T05:34:19+03:00
draft: false
tags:
  - linux
  - rant
---

I haven't used Ubuntu since I first discovered Linux 9 years ago in University, so I thought I'd give it a this past weekend.
I didn't like it. I only used it for a day and a half before I had to wipe everything and go back to Manjaro.

Before I get into the things that really bothered me, here's what I liked:

## 1. Online accounts

This is more of a Gnome thing than Ubuntu, but still, I love that I can add all my online accounts in the settings and have everything synced up automatically in the calendar and email apps.
I love KDE Plasma, but this is the one thing I really miss in it.

I actually don't care much for email, but calendar notifications is something I really need. I always find myself forgetting about scheduled meetings.
Ofcourse I could find another way to deal with this, but it would be really nice if I could have it OOTB with my distro.

## 2. Wallpapers

Not really a major thing but Ubuntu and ElementaryOS are the two distros I've seen that always have the best wallpapers shipped by default.

## 3. Somewhat consistent themeing

I think Ubuntu does this really well with the Yaru theme. I say somewhat because this is Gnome, and there are always [issues](#3-themeing-is-such-a-pain) with themeing Gnome.

That's about it for what I liked.

Here's what I didn't like:

## 1. Installer didn't work

This is a big problem. Ubuntu released a new installer written in flutter and it just didn't work for me.

It kept crashing in the middle of the setup process (before I actually got to the final install step), and when I finally got to the final step and clicked install, it froze at 5%.

I tried multiple times and even downloaded an update for the installer. Same thing.

Finally, I decided to go back to the website and download the legacy installer version - which worked fine.

This should have been a huge red flag for me. I should have given up on the whole thing the moment I had issues with the installer. But oh well, I apparently had time to waste.

This was a bad one.

Everything else after this is based on personal preference so not anything really wrong with Ubuntu itself.

## 2. Huge window title bars

I think the window title bars on Gnome are ridiculously huge. I can't stand them.
I tried getting around that with a theme that offers a "compact" version, but -

## 3. Themeing is such a pain

Why do I need to install an extra piece of software (Tweaks) to be able to change the default system font?
I don't understand the reasoning behind this. At the very least, this should be an option in the settings.

Fonts aside, applying a new shell theme doesn't really change the theme for all apps.
Something to do with "Libadwaita" and legacy apps. I didn't look into it enough to figure out what was going on.
But apps like the terminal and the file manager just didn't want to be themed differently.
I use these apps a lot so it would be too distracting for me to have them not looking the same as everything else.

## 4. Different ways to install software

This is totally me having been spoiled by the conveniences of the [AUR](https://aur.archlinux.org/).
When I install an arch based distro, I have a script that installs [yay](https://github.com/Jguer/yay), and uses it to install everything else from wherever it's available (official repos or AUR, I don't care).

`yay -S literally-anything`. It Just Works ™.

With Ubuntu, I have to figure out if the thing is in the repos. I may have to add some new repos, add GPG keys and run system updates.
Maybe the thing is a snap. In that case, I need to run some different commands to install snaps.

I wan't even able to install previous versions of Python [^1]. Something about it being a security issue for non-lts versions (can't remember the exact error).

There's probably something out there like yay but for Ubuntu - I don't know. But as it is, it's all just too much work.

[^1]: Ofcourse I should have used pyenv instead, but ain't nobody got time for that.
