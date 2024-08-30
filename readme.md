# ranger

![Golang RegExp Ranger](https://raw.githubusercontent.com/SimonWaldherr/ranger/master/ranger.gif)

[ranger.go](https://simonwaldherr.de/go/ranger) generates regexp code for numeric ranges and is inspired by [dimka665/range-regex](https://github.com/dimka665/range-regex)  


[![Go Report Card](https://goreportcard.com/badge/simonwaldherr.de/go/ranger)](https://goreportcard.com/report/simonwaldherr.de/go/ranger)
[![codebeat badge](https://codebeat.co/badges/bb574430-ee9e-4d62-a6d0-6daff78a5c08)](https://codebeat.co/projects/github-com-simonwaldherr-ranger-master) 
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/SimonWaldherr/ranger) 

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

## is it fast?

no, definitely not!  
take a look at the benchmark in [ranger_test.go](https://github.com/SimonWaldherr/ranger/blob/master/ranger_test.go) before using it.
