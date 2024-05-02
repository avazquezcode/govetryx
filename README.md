# GoVetryx
[![Go](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/avazquezcode/govetryx/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-81.9%25-brightgreen)

> [!NOTE]  
> This is a really simple language, just developed for fun (not production ready).

## Intro
GoVetryx is an interpreter written in Golang for the _Vetryx_ language.

## Vetryx
I have been reading recently the **AMAZING** book [Crafting Interpreters](https://www.amazon.com/dp/0990582930) and it inspired me a lot to try those ideas out.

This language is based on the *Lox* language implementation in _Java_ created by Robert Nystrom in the book mentioned above. I tried to follow the exact same implementation from the book (except for some tiny changes I did to the syntax and grammar, and some minor things I changed to experiment a bit on the topic). Since I changed some tiny things in the language, I added a different name for it: Vetryx.

It's been a while since I wanted to dive deeper into this topic, so this is just to do some practice around the concepts explained in the book.

## About the language
- Refer to this [doc](LANGUAGE.md) to understand the syntax of the language, and some of its rules.

Some things I changed in comparison to the _Lox_ language:

- [x] New lines should be used to terminate statements instead of semicolons (;) 
- [x] Parentheses to wrap if/while conditions are not necessary (but can be added if wanted)
- [x] Added modulus operator
- [x] Changed some reserved words & chars used for some operators
- [x] Added support for short variable declarator (eg: `a := 1`)
- [x] Added support for sleep, min and max native fns
- [x] Support `break` and `continue` in while loop

Coming next:
- [ ] Add support for structs
- [ ] Add support for arrays
- [ ] Add more native functions

## About the interpreter
This interpreter is coded in Golang. It's implemented as a Tree-Walk Interpreter (trying to follow the ideas and same implementation from the _Java_ implementation in the *Crafting Interpreters* book, so all the credit goes to Bob Nystrom here - althought I tried out some slightly different implementations in some parts just for fun & to experiment a bit on this topic).

## Playground

I developed a playground for it, using Next.js and NextUI.
You can test it online [here](https://govetryx.agustinvazquez.me/).

## Compiler
- I also intend to create a compiler later. But step by step :)

## Some words about Crafting Interpreters
Thanks a lot Robert Nystrom for writing such a pleasant book to read, and to learn about! Also the code in the book is really easy to follow! One of the nicest books about software that I've read in the past years!
