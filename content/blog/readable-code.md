---
title: "Writing readable code"
date: 2019-10-10T18:53:42+03:00
---

> Readability, I have been told, is not everything. Neither is breathing, but it does come before whatever comes next. <footer>William Sloane, The Craft of Writing, (W.W. Norton, 1979), page 11.</footer>

<span></span>

> Readability counts. <footer>The Zen of Python</footer>

It's not uncommon to write code only you can understand, assuming no one else would have to read it, then come back a few months later to find even you can't understand what the code is doing.

![crappy-code](../images/who-wrote-this-crap-code.png)

Be kind to the person who's going to maintain the code you write in future. It's likely going to be you.

![maintaining](../images/maintain.jpg)

# A few tips on writing readable code

## Do:

### Indent properly

Some languages (like Python) are strict on indentation - which is good. Others may not be. Either way, make it a practice to indent your code uniformly such that it's readable enough.

> Beautiful is better than ugly. <footer>The Zen of Python</footer>

### Read other people's code

This will help put you in the next person's shoes. If the code is readable, you'll, 1) get to see what makes readable code, and 2) appreciate the feeling of not having to struggle so much to make sense of things. If the code is not so readable, you'll learn what mistakes to avoid.

### Limit functions to a single task, or to single and highly related tasks

As much as possible.

### Have meaningful names for those functions

This actually builds on the last point. If functions are limited to a single task, it's easier to name them.

### Write short descriptive comments

When necessary.

## Don't:

### Write comments reiterating code

It's much better to get in the habit of writing self-documenting code, and reserving comments for describing the plain-English intent of the code.

### Be too clever

I don't know much about other languages, but I think Python is notorious for this kind of thing with its one-liners. I find these especially unreadable when lambda functions are involved. As elegant as a solution may seem at first, always consider the readability aspect.

> Simple is better than complex. <footer>The Zen of Python</footer>

### Use magic numbers

Magic numbers are unnamed numerical constants that appear in the code. If you were writing a piece of code to shuffle the values in a list representing a standard pack of playing cards for example, it would be more readable to write:

```py
deck_size = 52

for card in range(0, deck_size):
...
```

It's very clear from the onset here that the number `52` represents the number of cards in a deck.

```py
for x in range(0, 52):
...
```

Here, we are left wondering where `52` came from. We'd have to dig more into the code to figure it out. `52` here may be considered a magic number. Not good for readability.

That being said, there are cases where the use of magic numbers is accepted. You can read more about that <a href="https://en.wikipedia.org/wiki/Magic_number_(programming)#Accepted_limited_use_of_magic_numbers" target="_blank">here</a>.
