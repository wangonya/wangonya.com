---
author: Kelvin Wangonya
date: 2019-08-20
title: Testing Click applications with Pytest
tags:
  - python
  - tutorial
  - python-click
  - pytest
---

It\'s good practice to, as much as possible, write tests for your code.
If you\'re working with Python, [pytest](https://pytest.org/en/latest/)
makes the process of writing and running tests much smoother. I wrote a
few posts some time back on getting started with testing with pytest, so
if you\'re completely new to it, you might want to take a look at them:

- [Getting started with
  pytest](https://wangonya.com/blog/getting-started-with-pytest/)
- [Asserting Exceptions with
  Pytest](https://wangonya.com/blog/pytest-asserting-exceptions/)
- [Capturing print statements while
  debugging](https://wangonya.com/blog/pytest-capture-print/)
- [Skipping tests](https://wangonya.com/blog/pytest-skip/)

For testing CLI apps, Click provides a convenient module:
`click.testing` which has some useful functions (notably
`CliRunner()`) to help us invoke commands and check their
behavior.

We\'ll go ahead and test each part of [our
app](https://wangonya.com/blog/cli-crud-with-firebase/) - creating,
reading, updating and deleting.

### Installing pytest and writing the first test

pytest can be installed via pip:

    (env) $ pip install pytest

After installing pytest, create a tests folder in the root directory and
add the first test file:

    (env) $ mkdir tests && cd tests

    (env) $ touch test_app.py

In the test~app~ file, add the following code for a start:

```python
def test_add():
    pass
```

To run the test, run `pytest` on the terminal:

```shell
(env) $ pytest
================== test session starts ====================
platform linux -- Python 3.7.3, pytest-5.1.0, py-1.8.0, pluggy-0.12.0
rootdir: /home/wangonya/code/contacts-cli
collected 1 item

tests/test_app.py .                              [100%]

================== 1 passed in 0.04s =======================
```

### Testing the `add` command

Let\'s edit the test~app~ file to add a test to see if the
`add` command adds a new contact:

```python
from click.testing import CliRunner

from app import add

runner = CliRunner()

def test_add():
    response = runner.invoke(add, ["test-user", "-m", "0"])
    assert response.exit_code == 0
    assert "Contact test-user added!" in response.output
    assert "{'mobile': '0'}" in response.output
```

First, we invoke the command as we would on the terminal, passing in the
required arguments and options:
`response = runner.invoke(add, ["test-user", "-m", "0"])`.

We then check that the command executes successfully:
`assert response.exit_code =` 0=.

If the command executes successfully, we expect a success message should
be returned in the response with the values we added:

```python
assert "Contact test-user added!" in response.output
assert "{'mobile': '0'}" in response.output
```

The rest of the tests will pretty much follow the same format.

### Testing the `list` command

```python
def test_list():
    response = runner.invoke(list)
    assert response.exit_code == 0
    assert "Here\'s a list of all your contacts:" in response.output
    assert "'test-user': {'mobile': '0'}" in response.output
```

The list command doesn\'t take any arguments or options so we just call
it directly: `response = runner.invoke(list)`.

### Testing the `view` command

```python
def test_view():
    response = runner.invoke(view, "test-user")
    assert response.exit_code == 0
    assert "{'mobile': '0'}" in response.output
```

### Testing the `update` command

```python
def test_update():
    response = runner.invoke(update, ["test-user", "-m", "12345"])
    assert response.exit_code == 0
    assert "Contact updated!" in response.output
    assert "{'mobile': '12345'}" in response.output
```

### Testing the `delete` command

```python
def test_delete():
    response = runner.invoke(delete, "test-user")
    assert response.exit_code == 0
    assert "Contact deleted!" in response.output

    # call view on test-user to confirm it doesn't exist
    response = runner.invoke(view, "test-user")
    assert response.exit_code == 0
    assert "The contact you searched for doesn't exist" in response.output
```

### Improvements

As your application grows, you may want to consider using
[fixtures](https://docs.pytest.org/en/latest/fixture.html) and set up
things like `runner` in a `conftest.py` file. We
got away with it here because our tests were simple and all in a single
file. Once multiple test files are introduced, following the approach we
used here would lead to a lot of unnecessarily duplicated code.

Also, we made direct calls to our API in the tests. This operation
should ideally be mocked.
