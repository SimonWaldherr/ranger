# ranger

![Golang RegExp Ranger](https://raw.githubusercontent.com/SimonWaldherr/ranger/master/ranger.gif)

[ranger.go](https://simonwaldherr.de/go/ranger) generates regexp code for numeric ranges and is inspired by [dimka665/range-regex](https://github.com/dimka665/range-regex)  

[![Build Status](https://travis-ci.org/SimonWaldherr/ranger.svg?branch=master)](https://travis-ci.org/SimonWaldherr/ranger)

## what

generates a [regular expression](https://en.wikipedia.org/wiki/Regular_expression) that matches any number between two given numbers.

## how

first import ```simonwaldherr.de/go/ranger``` and then   
call ```ranger.Compile(23,37)``` and you will get ```^(2[3-9]|3[0-7])$```.

## why

because we can

## who

[me](https://simonwaldherr.de)

## is it any good

[yes](http://news.ycombinator.com/item?id=3067434).

## is it "Production Ready™"?

maybe. maybe not
