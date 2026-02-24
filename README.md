`goforth` is a toy Go implementation of a REPL-driven Forth. 

As of this writing, `goforth` is little more than a calculator. It supports only integers. It doesn't have conditionals, loops, or strings. It lacks good error messages. And it probably has bugs that may never be fixed. Turning this into a beautiful, complete Forth implementation isn't out of the question, but finding the time might be.

`goforth` runs from the command line. Follow these steps to get `goforth` going:

1. [Install Go.](https://go.dev/doc/install)
2. Clone this repository.
3. Run `go run .` in the project root.

You should now see the prompt
```
goforth>
```
Type `^C` to exit the program.

For an introduction to Forth, see [Starting Forth](https://www.forth.com/starting-forth/).

There are two very important things to understand about Forth when getting started:
- Forth uses postfix notation, where an operator appears _after_ its arguments.
- Execution is performed on a last-in, first-out stack.

To add the numbers 1 and 2, and then see the result, enter

```
goforth> 1 2 + .
3
```

In words, this line corresponds to 4 steps:
1. push 1 onto the stack.
2. push 2 onto the stack.
3. consume the two previous numbers, add the result, and push the result onto the stack.
4. pop the value at the head of the stack and print the result.

Without the `.` at the end, you wouldn't see the result.

Here's a more complex arithmetic expression):

```
goforth> 1 2 + 4 * 6 / 1 - .
1
```
It might be helpful to work through this by hand to convince yourself you know what's going on.

`goforth` also supports conditionals, but the results might look a little surprising:

```
goforth> 2 1 > .
-1
```

This isn't a typo. In Forth, the _true_ value is -1 and _false_ is 0 (yes, there is a [reason](https://en.wikipedia.org/wiki/Two%27s_complement) for this). Here are two more examples:

```
goforth> 2 1 > 3 2 > and .
-1
goforth> 2 1 < 0 or .
0
```

In addition to arithmetic, boolean operations, and "dot" (`.`), `goforth` offers two basic stack manipulation operators: `dup` to duplicate the head of the stack, and `swap` to interchange the top two values. To square a number, we'd run

```
goforth> 3 dup * .
9
```

We might use `swap` to reverse argument order in an operation that is order dependent (convince yourself that this is `6 / 2 = 3`):

```
goforth> 2 10 4 - swap / .
3
```

The most advanced feature in `goforth` is function definition. A function in Forth is called a "word." All of the operations encountered so far are built-in words. A new word is defined as `: <name> <body> ;`. When `:` is encountered, the program state is switched to compile mode. The `name` of the word is recorded in the program's environment for later use, and `body` holds other words and integers that define the new operation.

Suppose we wanted to write a function that squares a number. It would look like this:

```
goforth> : square dup * ;
goforth> 4 square .
16
```

Once a word is defined, it can be used to define other words:

```
goforth> : forth-power square square ;
goforth> 3 forth-power .
81
```

Get it? _forth_-power... 😂
