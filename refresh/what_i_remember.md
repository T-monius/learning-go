# What I Remember

## General

- Go is a compiled language, so it's not possible to use a REPL
  - To test the code, it needs to be run, and logging can be utilized
- Typed
- Pascal Case

## Logging
```go
[someLib].print()
[someLib].printLn()
```
General format string functionality

## Spacing
2 space? but not opinionated?

## Comments
Single line `//`
Multi-line, can't remember
Relied on for making docs (or is that Python?)

## Project Setup

`go mod init` creates a module
`go run` runs a file


## Variables

```go
var str string = "hi"
str1 := "hi"
```

## Flow Control

```go
if ?

for (el := dataStructure) {
  // ...
}
```


## Data Structures

Strings
- Immutable?
- Didn't learn to many string manipulation functions

Arrays
- Fixed length
- `len` and `cap` functions get length and capacity
- Typed
```go
var arr array[int] = []?
var arr array[]int = []
```

SomeOtherArrayLikeDS
- How mutable array-like ds are represented in Go

Struct
- Basis for OO
- Acts like a dictionary
- Key value relationships
```go
var dict struct{int} = {}
```

Interface
- Relationship w/ struct for OO
- 
