## Kata02

### Problem statement

[Source](http://codekata.com/kata/kata02-karate-chop/)

Write a binary chop method that takes an integer search target and a sorted array of integers. It should return the integer index of the target in the array, or -1 if the target is not in the array. 

### Dependencies

- Go Lang
- https://github.com/stretchr/testify

### Run instructions

- Testing `go test`
- Coverage `go test -coverprofile=$t $@ && go tool cover -html="/tmp/go-cover.$$.tmp" && unlink "/tmp/go-cover.$$.tmp"`
