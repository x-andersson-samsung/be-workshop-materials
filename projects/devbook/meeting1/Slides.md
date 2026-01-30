---
title: "Devbook "
level: basic
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-01-30 17-09-14
slideNumber: "true"
---

 %% Required for proper codeblock width %%
<style>
.reveal pre {
  width: 110%;
}

code {
    font-size: 16px;
    line-height: normal;
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
- Discussing next homework
</grid>

:::

---

# Project

--

<grid drag="100 10" drop="0 0" align="left" >
### Description
</grid>

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

<grid drag="50 85" drop="0 10" align="left">
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
### Initialization
</grid>

<grid drag="50 85" drop="0 10" align="left">
```go []
var longVar int = 0
shortVar := 0

var (
	manyVars1 = 0
	manyVars2 = "hello"
)

const unchangechable = 0
```
</grid>

---
# Types

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
- nil for pointers and pointer based types (slices, maps)
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
Structures are a collection of fields.
```go []
type Point struct {
	X int
	Y int
}

v1 := Point{}
v2 := Point{1,2}
//v2 := Point{1} -- invalid, should 1 go to X or Y?
v3 := Point{X: 5}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Arrays
</grid>
	
<grid drag="100 85" drop="0 10" align="left">
The type **[n]T** is an array of **n** values of type **T**.
```go []
arr1 := [3]int{}				// [3]int{0,0,0}
arr2 := [3]int{1,2,3}		// [3]int{1,2,3}
arr3 := [4]int{1,2,3}		// [4]int{1,2,3,0}
//arr4 := [2]int{1,2,3}	// error - too many values
```
An array's length is part of its type, so arrays cannot be resized.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Slices
</grid>

<grid drag="100 85" drop="0 10" align="left">
A slice is a dynamically-sized, flexible view into the elements of an array. 

In simplified terms, it is a pointer to an array with length and capacity.
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

--

<grid drag="100 10" drop="0 0" align="left" >
### Slices 3
</grid>

<grid drag="100 85" drop="0 10" align="left">
When building slices we can skip the bounds to use defaults

```go []
var a [10]int

// For `a` these will be equivalent
a[0:10]
a[:10]
a[0:]
a[:]
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Maps
</grid>

<grid drag="100 85" drop="0 10" align="left">
A map maps keys to values. 

```go []
// Initialize 2 maps
m := make(map[int]string)
m2 := map[int]string{}

// Add value
m[1] = "Hello"

// get values
v := m[1]

// checking if value exists
v, ok := m[2] // Will return 0, false

// deleting key
delete(m, 1)
```
</grid>

---

# Flow control

--

<grid drag="100 10" drop="0 0" align="left" >
### If / else
</grid>

<grid drag="100 85" drop="0 10" align="left">
- Parentheses are optional
- Brackets are required

```go []
if b {
	fmt.Println("True")
} else {
	fmt.Println("False")
}

// We can also use if with "short statement". Following two checks are equivalent
var m map[int]string
_, ok := m[0]
if ok {
	// Do something
}

if v, ok := m[0]; ok {
	// Do something
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Loops
</grid>

<grid drag="100 85" drop="0 10" align="left">
- Only one type of loops - **for**

```go []
// Long form - "init; condition; post"
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// Short form - "condition"
for i < 10 {
	fmt.Println(i)
}



```
</grid>
