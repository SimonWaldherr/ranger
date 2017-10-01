# ranger

![Golang RegExp Ranger](https://raw.githubusercontent.com/SimonWaldherr/ranger/master/ranger.gif)

[ranger.go](https://simonwaldherr.de/go/ranger) generates regexp code for numeric ranges and is inspired by [dimka665/range-regex](https://github.com/dimka665/range-regex)  

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/633b21cb192a4c64a893b9e8bae50a07)](https://www.codacy.com/app/SimonWaldherr/ranger?utm_source=github.com&utm_medium=referral&utm_content=SimonWaldherr/ranger&utm_campaign=badger)
[![Build Status](https://travis-ci.org/SimonWaldherr/ranger.svg?branch=master)](https://travis-ci.org/SimonWaldherr/ranger) 
[![Go Report Card](https://goreportcard.com/badge/github.com/simonwaldherr/ranger)](https://goreportcard.com/report/github.com/simonwaldherr/ranger) 
[![codebeat badge](https://codebeat.co/badges/bb574430-ee9e-4d62-a6d0-6daff78a5c08)](https://codebeat.co/projects/github-com-simonwaldherr-ranger-master) 
[![BCH compliance](https://bettercodehub.com/edge/badge/SimonWaldherr/ranger?branch=master)](https://bettercodehub.com/) 
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/SimonWaldherr/ranger) 

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
