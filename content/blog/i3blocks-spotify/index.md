---
title: "Adding a Spotify blocklet to i3blocks"
date: 2023-06-14T21:34:04+03:00
draft: false
tags:
  - linux
  - i3wm
---

I like seeing what's playing without having to bring up the Spotify app every time. If I can also control playback without bring up the app, even better.
[i3blocks](https://github.com/vivien/i3blocks) makes doing this on i3 trivial.

By default, i3 comes with [i3bar](https://i3wm.org/i3bar/). I'm sure I could achieve the same thing on i3bar [with enough time and effort](https://i3wm.org/docs/i3status.html#_external_scripts_programs_with_i3status), but its just much easier to do it with i3blocks.

## Getting song info

[spotify-cli-linux](https://github.com/pwittchen/spotify-cli-linux) works well for this. It also provides a good [list of options](https://github.com/pwittchen/spotify-cli-linux#usage) to control playback and do a lot more.

## Adding the blocklet

Assuming i3blocks is already running,
