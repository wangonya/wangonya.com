---
title: "Working with celery signals"
author: Kelvin Wangonya
date: 2020-11-03T06:17:15+03:00
description:
taxonomies:
  tags:
    - python
    - celery
---

If you're new to celery, [start here](https://docs.celeryproject.org/en/stable/getting-started/introduction.html).

Sometimes when using celery, you may want to get notified when a task running in the background executes successfully or when it fails. You may also want to run a function each time before the celery task runs or after it completes. These and many others along the same line are all situations where [signals](https://docs.celeryproject.org/en/stable/userguide/signals.html) would come in handy.

<!--more-->

I'll create a simple file `tasks.py` and set up celery to demonstrate how to use celery signals.

```py
from celery import Celery

app = Celery('tasks', broker='redis://localhost:6379/0')

@app.task
def add(x, y):
    return x + y
```

Make sure your redis server is running and start your celery worker:

```sh
(env) $ celery -A tasks worker --loglevel=INFO
```

Then run your tasks.py file and execute the `add` task:

```sh
(env) $ python -i tasks.py

>>> add.delay(4, 4)
<AsyncResult: ce1ee079-6434-4f54-ace2-360ff316546b>
>>>
```

By default, what is returned is an [AsyncResult](https://docs.celeryproject.org/en/stable/reference/celery.result.html#celery.result.AsyncResult) instance but that's not what we're interested in. On the terminal with your Celery worker running, you should see something similar to this:

```sh
[2020-11-03 07:01:02,024: INFO/ForkPoolWorker-2] Task tasks.add[ce1ee079-6434-4f54-ace2-360ff316546b] succeeded in 0.0005510429999979749s: 8
```

The task executes successfully, and 8 is the result as expected.

## Signals

There are [a lot](https://docs.celeryproject.org/en/stable/userguide/signals.html) of signals that celery offers but I'll focus on 4 simple ones to demonstrate how signals work in general.

1. task_prerun
2. task_postrun
3. task_success
4. task_failure

### task_prerun

This signal is dispatched **before** a task is executed.

```py
from celery import Celery
from celery.signals import task_prerun

app = Celery('tasks', broker='redis://localhost:6379/0')

@app.task
def add(x, y):
    return x + y

@task_prerun.connect(sender=add)
def task_prerun_notifier(sender=None, **kwargs):
    print("From task_prerun_notifier ==> Running just before add() executes")
```

Sender is the task object being executed (the `add` function in this case).

Running `add.delay(4, 4)` like before now gives the following output on the celery terminal:

```sh
[2020-11-03 07:23:19,183: WARNING/ForkPoolWorker-2] From task_prerun_notifier ==> Running just before add() executes
[2020-11-03 07:23:19,184: INFO/ForkPoolWorker-2] Task tasks.add[1ef11c46-f461-4eb8-84ca-5c5cdab62a74] succeeded in 0.0016491969999998801s: 8
```

Just before the task runs, the signal dispatches and prints as expected.

### task_postrun

Dispatched **after** a task has been executed.

```py
from celery.signals import task_prerun, task_postrun

app = Celery('tasks', broker='redis://localhost:6379/0')

@app.task
def add(x, y):
    return x + y

@task_prerun.connect(sender=add)
def task_prerun_notifier(sender=None, **kwargs):
    print("From task_prerun_notifier ==> Running just before add() executes")

@task_postrun.connect(sender=add)
def task_postrun_notifier(sender=None, **kwargs):
    print("From task_postrun_notifier ==> Ok, done!")
```

Running this should give the following result:

```sh
[2020-11-03 17:03:51,655: WARNING/ForkPoolWorker-2] From task_prerun_notifier ==> Running just before add() executes
[2020-11-03 17:03:51,656: INFO/ForkPoolWorker-2] Task tasks.add[7da6ee71-1941-4a87-b993-8136d94ac067] succeeded in 0.0017917519999999243s: 8
[2020-11-03 17:03:51,657: WARNING/ForkPoolWorker-2] From task_postrun_notifier ==> Ok, done!
```

### task_success

Dispatched when a task succeeds.

```py
from celery import Celery
from celery.signals import task_prerun, task_postrun, task_success

app = Celery('tasks', broker='redis://localhost:6379/0')

@app.task
def add(x, y):
    return x + y

@task_prerun.connect(sender=add)
def task_prerun_notifier(sender=None, **kwargs):
    print("From task_prerun_notifier ==> Running just before add() executes")

@task_postrun.connect(sender=add)
def task_postrun_notifier(sender=None, **kwargs):
    print("From task_postrun_notifier ==> Ok, done!")

@task_success.connect(sender=add)
def task_success_notifier(sender=None, **kwargs):
    print("From task_success_notifier ==> Task run successfully!")
```

Result:

```sh
[2020-11-03 17:40:47,276: INFO/MainProcess] Received task: tasks.add[6603eb49-75ab-4653-b32f-ebe760a52de0]
[2020-11-03 17:40:47,279: WARNING/ForkPoolWorker-2] From task_prerun_notifier ==> Running just before add() executes
[2020-11-03 17:40:47,281: WARNING/ForkPoolWorker-2] From task_success_notifier ==> Task run successfully!
[2020-11-03 17:40:47,281: INFO/ForkPoolWorker-2] Task tasks.add[6603eb49-75ab-4653-b32f-ebe760a52de0] succeeded in 0.00201471799999986s: 8
[2020-11-03 17:40:47,282: WARNING/ForkPoolWorker-2] From task_postrun_notifier ==> Ok, done!
```

### task_failure

Dispatched when a task fails.

```py
from celery import Celery
from celery.signals import task_prerun, task_postrun, task_failure

app = Celery('tasks', broker='redis://localhost:6379/0')

@app.task
def add(x, y):
    raise Exception

@task_prerun.connect(sender=add)
def task_prerun_notifier(sender=None, **kwargs):
    print("From task_prerun_notifier ==> Running just before add() executes")

@task_postrun.connect(sender=add)
def task_postrun_notifier(sender=None, **kwargs):
    print("From task_postrun_notifier ==> Ok, done!")

@task_failure.connect(sender=add)
def task_failure_notifier(sender=None, **kwargs):
    print("From task_failure_notifier ==> Task failed successfully! 😅")
```

Result:

```sh
[2020-11-03 17:44:36,082: INFO/MainProcess] Received task: tasks.add[da4a03e8-5530-4c9e-afeb-75f8e0b1be5d]
[2020-11-03 17:44:36,085: WARNING/ForkPoolWorker-2] From task_prerun_notifier ==> Running just before add() executes
[2020-11-03 17:44:36,096: WARNING/ForkPoolWorker-2] From task_failure_notifier ==> Task failed successfully! 😅
[2020-11-03 17:44:36,096: ERROR/ForkPoolWorker-2] Task tasks.add[da4a03e8-5530-4c9e-afeb-75f8e0b1be5d] raised unexpected: Exception()
Traceback (most recent call last):
  ...
   in add
    raise Exception
Exception
[2020-11-03 17:44:36,097: WARNING/ForkPoolWorker-2] From task_postrun_notifier ==> Ok, done!
```
