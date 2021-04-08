---
title: "Writing DRYer tests using Pytest parametrize"
author: Kinyanjui Wangonya
date: 2021-04-08T17:08:41+03:00
description: "Taking advantage of Pytest's parametrize fixture to avoid unnecessary repetition in tests."
tags:
  - python
  - pytest
---

Tests tend to not always be so DRY, which isn't necessarily a bad thing.

[This](https://stackoverflow.com/a/129722/9312256) SO answer sums it up nicely:

> Readability is more important for tests. If a test fails, you want the problem to be obvious. The developer shouldn't have to wade through a lot of heavily factored test code to determine exactly what failed. You don't want your test code to become so complex that you need to write unit-test-tests.
>
> However, eliminating duplication is usually a good thing, as long as it doesn't obscure anything, and eliminating the duplication in your tests may lead to a better API. Just make sure you don't go past the point of diminishing returns.

Pytest gives a couple of ways to reduce duplication with fixtures.

Say you had a couple of endpoints that return data to be used in a report.
The report data is supposed to be displayed in an excel sheet with different sheets.

Sometimes, only data for one sheet is required. Other times, data for
all the sheets is fetched. So the endpoints end up being broken down like this:

```
/report/sheet-a
/report/sheet-b
/report/sheet-c
/report/sheet-d
```

Let's take a simple test case: checking that the endpoints return `200` when called.

```python
import pytest

def test_report_sheet_a_returns_200():
    response = client.get('/report/sheet-a')
    assert response.status == 200

def test_report_sheet_b_returns_200():
    response = client.get('/report/sheet-b')
    assert response.status == 200

def test_report_sheet_c_returns_200():
    response = client.get('/report/sheet-c')
    assert response.status == 200

def test_report_sheet_d_returns_200():
    response = client.get('/report/sheet-d')
    assert response.status == 200
```

This might not look too bad, but if we wanted to test for another operation on the endpoints -
checking if the endpoints require authentication for example - you get the feeling that this can
be done a little bit better.

```python
import pytest

# test GET requests

def test_report_sheet_a_returns_200():
    response = client.get('/report/sheet-a')
    assert response.status == 200

def test_report_sheet_b_returns_200():
    response = client.get('/report/sheet-b')
    assert response.status == 200

def test_report_sheet_c_returns_200():
    response = client.get('/report/sheet-c')
    assert response.status == 200

def test_report_sheet_d_returns_200():
    response = client.get('/report/sheet-d')
    assert response.status == 200

# test auth

def test_report_sheet_a_requires_auth():
    response = unauthorized_client.get('/report/sheet-a')
    assert response.status == 401

def test_report_sheet_b_requires_auth():
    response = unauthorized_client.get('/report/sheet-b')
    assert response.status == 401

def test_report_sheet_c_requires_auth():
    response = unauthorized_client.get('/report/sheet-c')
    assert response.status == 401

def test_report_sheet_d_requires_auth():
    response = unauthorized_client.get('/report/sheet-d')
    assert response.status == 401
```

Notice that the only thing changing in the tests is the endpoints. Everything else remains
the same.

Given that all the sheets belong to one report, we can refactor the tests to reduce duplication
without sacrificing readability.

Here's how the tests can be rewritten using [Pytest parametrize](https://docs.pytest.org/en/stable/parametrize.html#parametrize-basics):

```python
import pytest

report_sheet_endpoints = (
  '/sheet-a',
  '/sheet-b',
  '/sheet-c',
  '/sheet-d',
)


@pytest.mark.parametrize('endpoint', report_sheet_endpoints)
def test_report_sheets_return_200(endpoint)
    response = client.get(f'/report{endpoint}')
    assert response.status == 200


@pytest.mark.parametrize('endpoint', report_sheet_endpoints)
def test_report_sheets_require_auth(endpoint)
    response = unauthorized_client.get(f'/report{endpoint}')
    assert response.status == 401
```

Now we have two tests instead on eight. But when you run the tests, 8 tests will run,
not 2. Pytest takes each value in `report_sheet_endpoints` and feeds it into the test.
This reduces duplication while maintaining readability.

Running the tests gives this output:

```sh
test_reports.py::test_report_sheets_return_200[sheet-a] PASSED
test_reports.py::test_report_sheets_return_200[sheet-b] PASSED
test_reports.py::test_report_sheets_return_200[sheet-c] PASSED
test_reports.py::test_report_sheets_return_200[sheet-d] PASSED

test_reports.py::test_report_sheets_require_auth[sheet-a] PASSED
test_reports.py::test_report_sheets_require_auth[sheet-b] PASSED
test_reports.py::test_report_sheets_require_auth[sheet-c] PASSED
test_reports.py::test_report_sheets_require_auth[sheet-d] PASSED
```
