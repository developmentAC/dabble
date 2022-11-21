# Dabble: A basic shell written in GO
## Date: 21 November 2022
### Oliver Bonham-Carter, [Allegheny College](https://allegheny.edu/)
### email: obonhamcarter@allegheny.edu

---
<!-- ![logo](graphics/beagleTM_logo2.png) -->

												
[![MIT Licence](https://img.shields.io/bower/l/bootstrap)](https://opensource.org/licenses/MIT)
[![BLM](https://img.shields.io/badge/BlackLivesMatter-yellow)](https://blacklivesmatter.com/)

GitHub link: https://github.com/developmentAC/dabble

## Table of contents

* [Overview](#overview)

## Overview

The Shell is a software program that runs in the terminal (i.e., a command line interpreter).
Shells are import pieces of software as they allow users to access parts of the
OS with command. Naturally, a file in the path could be run to achieve the same
result but when using a shell, parts of the individual parts of the shell
may be untied in some way to work together.

In this (on going) project, `Dabble` is a shell to demonstrate how a common shell might work. In addition, the code allows for functions to be shred between commands.

## Executing the Program

To run `Dabble`, you must have go installed. Dabble was created on a MacOS with `go version go1.13.5 darwin/amd64` 

To run: `go run dabble`

And you will see the following,

```
 go run dabble.go 
	 :: OBC's Dabble Terminal ::

 [+] Enter a system command (help) :help

 [+] System commands:
	 [+=+] pwd :Show current path
	 [+=+] exit :End program
	 [+=+] help :Show the menu of commands
	 [+=+] square :Draw a square
	 [+=+] ffile :Create a file with a fortune in it
	 [+=+] ran :Create n random numbers between min and max bounds
	 [+=+] prime :Create a random prime number
	 [+=+] hello :Say hello

 [+] Enter a system command (help) :
```

## Functionality

* Find prime numbers in a defined range.

```
	 [+] Finding all prime numbers between two values
	 Enter a lower bounds :3

	 Enter an upper bounds for a number :10

 3 
 5 
 7 
```

* Find a range of random numbers in a defined range.
```
	 We will find n random numbers between lower a and upper bounds
	 Enter a lower bounds :2

	 Enter an upper bounds for a number :10

	 Number of random numbers to create :4

3
9
9
5
```



## A work in progress

Check back often to see the evolution of the project!! BeagleTM is a work-in-progress. Updates to the methods and tests for the code will come soon and I will continue to update the repository with updates. If you would like to contribute to this project, __then please do!__ For instance, if you see some low-hanging fruit or task that you could easily complete, that could add value to the project, then I would love to have your insight.

Otherwise, please create an Issue for bugs or errors. Since I am a teaching faculty member at Allegheny College, I may not have all the time necessary to quickly fix the bugs and so I would be very happy to have any help that I can get from the OpenSource community for any technological insight. Much thanks in advance. I hope that this project helps you find the knowledge from PubMed that you require for your project. :-)
