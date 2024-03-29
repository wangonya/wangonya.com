---
author: Kelvin Wangonya
date: 2019-08-07
title: Adding options to CLI commands
tags:
  - python
  - tutorial
  - python-click
---

In the previous post, we used `setuptools` to package our app
and enable us to run it using a single executable command
`hello`. In this post, we look at how to make CLI apps more
functional and interactive using options.

## Options

Options are used to alter the default behavior of commands. This is
often accomplished by passing in parameters in the form of
`--option` or `-o` for short.

Let\'s assume that instead of just printing out \"Hello World!\", we
wanted to have the option of printing it in all lowercase characters or
all uppercase characters. We can add this functionality by adding an
`option()` decorator.

```python
# helloworld.py

import click

@click.command()
@click.option('-c', '--case', type=click.Choice(['upper', 'lower']))
def hello(case):
    if case:
        click.echo('You selected {} case.'.format(case))
    # click.echo('Hello World!')
```

Then run the app, passing in the created option:

```shell
(venv) $ hello -c upper  # --case would also work
You selected upper case.
```

```shell
(venv) $ hello -c lower
You selected lower case.
```

Now try passing in some other option, like \"normal\":

```shell
(venv) $ hello -c normal
Usage: hello [OPTIONS]
Try "hello --help" for help.

Error: Invalid value for "-c" / "--case": invalid choice: normal. (choose from upper, lower)
```

It gives us an error: `invalid choice: normal` and even
guides us on the options to choose from
`(choose from upper, lower)`, since that\'s what we specified
`type=click.Choice(['upper', 'lower'])`. Error handling on
that part is handled by default.

Ok, now that the option is passed in successfully, we can add in the
logic to see our desired output.

```python
import click

@click.command()
@click.option('-c', '--case', type=click.Choice(['upper', 'lower']))
def hello(case):
    response = "Hello World!"
    if case == 'upper':
        click.echo(response.upper())
    elif case == 'lower':
        click.echo(response.lower())
    else:
        click.echo(response)
```

All we\'ve done is add an `if-else` to check the option
passed in, and return the text in either upper case, lower case, or just
normal. Now when we install our app again and run it, we get these
results:

```shell
(venv) $ hello
Hello World!

(venv) $ hello -c upper
HELLO WORLD!

(venv) $ hello -c lower
hello world!
```

In the next post, we\'ll enhance our command even more by adding
arguments.
