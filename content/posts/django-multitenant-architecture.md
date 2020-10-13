---
title: "Setting up a multi-tenant architecture with Django and Postgres"
author: Kinyanjui Wangonya
date: 2020-10-13T13:34:16+03:00
description: Setting up a multi-tenant architecture with Django and Postgres
tags:
  - python
  - django
  - architecture
  - postgres
---

I was recently tasked to help set up a multi-tenant architecture for an existing (in development) backend API made with Django/DRF. The whole concept of multi-tenancy was new to me so this was a grest learning experience.

<!--more-->

The API in question is meant to serve a SaaS application so the main purpose of this endevour was to achieve data segregation for users.
