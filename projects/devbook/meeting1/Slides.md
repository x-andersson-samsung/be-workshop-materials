---
title: "Devbook "
level:
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-01-26 13-56-37
slideNumber: "true"
---

 %% Required for proper codeblock width %%
<style>
.reveal pre {
  width: 110%; /* increase width */
}
</style>

%% Start of slides %%


# Devbook 
## Meeting 1

---

# Introduction

--

<grid drag="100 10" drop="0 0" align="left" >
 ### Goal
</grid>


<grid drag="100 85" drop="0 10" align="left">
- A project based introduction to BE concepts
- Each project taking around 5 meetings
- Finished application at the end
</grid>

note: 
Some projects might require some knowledge or attending a previous project.


--

<grid drag="100 10" drop="0 0" align="left" >
 ### Meeting
</grid>

<grid drag="100 85" drop="0 10" align="left">
Meetings will be split into:
- Discussing homework
- New theory
- Practical exercise - building on project
- Discussing new homework
</grid>

:::

---

# Project

--

<grid drag="100 10" drop="0 0" align="left" >
_**DevBook**_ is a simple app where developers can manage useful resources like articles, tutorials, tools, and libraries. 

Users can create resources with titles, URLs, descriptions, categories, and tags.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Steps
</grid>

<grid drag="100 85" drop="0 10" align="left">
1. Basic functionality with Terminal UI
2. REST API
3. Permanent storage with PostgreSQL
4. Containerization with Docker
5. Testing and Documentation
</grid>

---

# Go

--

<grid drag="100 10" drop="0 0" align="left" >
### What is Go?
</grid>

<grid drag="100 85" drop="0 10" align="left">
- Compiled
- Statically typed
- Garbage-collected
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Syntax
</grid>

<grid drag="100 85" drop="0 10" align="left">
```go []
package main

import (
	"fmt"
)

func main() {
	name := "World"
	fmt.Printf("Hello %s!\n", name)
}

```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Base Types
</grid>

<grid drag="100 85" drop="0 10" align="left">
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
note: The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Zero values
</grid>

<grid drag="100 85" drop="0 10" align="left">
Uninitialized values are given initial zero value depending on type:
- 0 for numeric types,
- false for the boolean type
- "" (the empty string) for strings.
- nil for pointers and pointer based types
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Pointers
</grid>

<grid drag="100 85" drop="0 10" align="left">
Pointers hold the memory address.
```go []
// Declare a variable
a := 1

// Declare a pointer
ptr := &a

// Change the value
*ptr = 2

fmt.Println(a) // Will print "2"
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Structs
</grid>

<grid drag="100 85" drop="0 10" align="left">
```go []
type Point struct {
	X int
	Y int
}

v1 := Point{}
v2 := Point{1,2}
v3 := Point{X: 5}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Arrays
</grid>

<grid drag="100 85" drop="0 10" align="left">
```go []
arr1 := [3]int{}              // [3]int{0,0,0}
arr2 := [3]int{1,2,3}     // [3]int{1,2,3}
arr3 := [4]int{1,2,3}     // [4]int{1,2,3,0}
//arr4 := [2]int{1,2,3}  // error - too many values
```
An array's length is part of its type, so arrays cannot be resized.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Slices
</grid>

<grid drag="100 85" drop="0 10" align="left">
A slice is a dynamically-sized, flexible view into the elements of an array. 
```go []
// create a new array of size 5 and return slice on it
slice := make([]int, 5) 

// slice is a reference to underlying array
numbers := [6]int{0,1,2,3,4,5}

var s []int = numbers[1:4] // []int{1,2,3}
s[0] = 9
s[1] = 8

fmt.Println(numbers)
// 0, 9, 8, 3, 4, 5
```
</grid>

note: Slice contains 3 values -> pointer to beginning, size, capacity

--


<grid drag="100 10" drop="0 0" align="left" >
### Slices 2
</grid>

<grid drag="100 85" drop="0 10" align="left">
Slice can be extended to add new values.
```go []
// define a slice of size 3 and cap 5
slice := make([]int, 3, 5) 

// Let's add 2 values
slice = append(slice, 1)
slice = append(slice, 2) // []int{0,0,0,1,2}



// This append will go beyond capacity 
// and will cause reallocation and copy
slice = append(slice, 3)

// []int{0,0,0,1,2,3}
```
</grid>