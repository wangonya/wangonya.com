#!/bin/sh
# scaffold blog post template
_=$1
DATE=$(date +"%Y-%m-%d")
TEMPLATE=$(echo "---
title: $1
date: $DATE
description:
---")
FOLDER_NAME=$(echo "$1" | iconv -t ascii//TRANSLIT | sed -E 's/[~\^]+//g' | sed -E 's/[^a-zA-Z0-9]+/-/g' | sed -E 's/^-+\|-+$//g' | sed -E 's/^-+//g' | sed -E 's/-+$//g' | tr A-Z a-z)
cd content/blog && mkdir "$FOLDER_NAME" && cd "$FOLDER_NAME" && echo "$TEMPLATE" > index.md
echo "created post: $FOLDER_NAME"