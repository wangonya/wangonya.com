---
author: Kelvin Wangonya
date: "2024-03-06T05:15:40+03:00"
title: Pydantic validators don't raise validation errors immediately
tags:
  - python
  - til
---

This one really had me confused. I was facing an error with [Pydantic validators](https://docs.pydantic.dev/latest/concepts/validators/) which I thought I had handled, but it just didn't work as expected.

The schema looked something like this:

```py
class Schema(BaseModel):
    dates: List[str]
    start_date: Optional[str]
    end_date: Optional[str]
    ...

    @validator("dates")
    def update_date_format(cls, v):
        ...

    @validator("start_date")
    def set_start_date(cls, v, values):
        # some logic here that expects `dates` to exist
        v = values['dates'][0]
        return v

    @validator("end_date")
    def set_end_date(cls, v, values):
        # some logic here that expects `dates` to exist
        v = values['dates'][-1]
        return v
```

My assumption in the last two validators was that `dates` would always be available because it's a required field.
To my surprise, an `IndexError` was always raised on `set_start_date`.

So I thought, ok, I'll raise a `ValueError` myself on `update_date_format` since the required field check doesn't seem to be working.

```py
class Schema(BaseModel):
    dates: List[str]
    start_date: Optional[str]
    end_date: Optional[str]
    ...

    @validator("dates")
    def update_date_format(cls, v):
        if not v:
            raise ValueError("dates is a required field")
        ...

    @validator("start_date")
    def set_start_date(cls, v, values):
        # some logic here that expects `dates` to exist
        v = values['dates'][0]
        return v

    @validator("end_date")
    def set_end_date(cls, v, values):
        # some logic here that expects `dates` to exist
        v = values['dates'][-1]
        return v
```

Again, an `IndexError` was raised on `set_start_date`. It's like my check on `update_date_format` was completely ignored.
I even put a breakpoint on the line to make sure it was being hit. It was.

So, validation errrors aren't raised immediately. They're collected and returned all at once in the response.

Despite seeing this countless times in `422` responses, I had never thought about it. It made perfect sense.
You don't want to return one error at a time as an API response. It's much better to return a response detailing everything that's wrong with the payload.

From the [docs](https://docs.pydantic.dev/latest/errors/errors/):

> One exception will be raised regardless of the number of errors found, that `ValidationError` will contain information about all the errors and how they happened.

The important thing to remember is that this only happens for validation errors (i.e errors raised through `ValueError` or `AssertionError`).
This is why the `IndexError` was always raised immediately it happened.

## How to stop validation on the first error

According to [this answer](https://stackoverflow.com/a/69538066/9312256):

> If you have checks, the failure of which should interrupt the further validation, then put them in the `pre=True` root validator. Because field validation will not occur if `pre=True` root validators raise an error.
>
> For example:
>
> ```py
> class PayloadValidator(BaseModel):
>    emailId: List[str]
>    role: str
>
>    @root_validator(pre=True)
>    def root_validate(cls, values):
>        if not values['emailId']:
>            raise ValueError("Email list is empty.")
>        return values
>
>    @validator("emailId")
>    def valid_domains(cls, emailId):
>        return emailId
> ```
