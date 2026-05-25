---
title: "Devbook "
level: basic
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-04-22 11-28-41
slideNumber: "true"
---

 %% Required for proper codeblock width %%
<style>

li,p {
	font-size: 32px;
}

code {
    font-size: 16px;
    line-height: normal;
}

/* left-align all content in Slides */
.reveal .slides {
    text-align: left;
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


<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Meetings will be split into:
- Discussing homework
- New theory
- Practical exercise & building on project
- Discussing next homework
</grid>

---

# Project

--

<grid drag="100 10" drop="0 0" align="left">
### Description
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
_**DevBook**_ is a simple app where developers can manage useful resources like articles, tutorials, tools, and libraries. 

Users can create resources with titles, URLs, descriptions, categories, and tags.
</grid>

--

<grid drag="100 10" drop="0 0" align="left">
### Steps
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
1. Basic functionality with Terminal UI
2. REST API
3. Permanent storage with PostgreSQL
4. Containerization with Docker
5. Testing, Documentation & Best practices
</grid>

---

# Go

--

<grid drag="100 10" drop="0 0" align="left" >
### What is Go?
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Compiled
- Expressive
- Statically typed
- Garbage-collected
</grid>

--


<grid drag="100 10" drop="0 0" align="left" >
### Why Go?
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Easy to learn
- Rich std library
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Syntax
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
package main

import (
	"fmt"
)

func name() string {
	return "World"
}

func main() {
	name := name()
	fmt.Printf("Hello %s!\n", name)
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Variables
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
var longVar int = 0
shortVar := 0

var (
	manyVars1 = 0
	manyVars2 = "hello"
)

const unchangeable = 0
```
</grid>

---
# Types

--

<grid drag="100 10" drop="0 0" align="left" >
### Base Types
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Uninitialized values are given initial zero value depending on type:
- **_0_** - for numeric types,
- **_false_** - for the boolean type
- **_""_** - (the empty string) for strings.
- **_nil_** - for pointers and pointer based types (slices, maps, interfaces)
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Type casting
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go does not perform implicit casting. If you want to change variable type you have to do it explicitly.

```go[]
var (
	i int
	i32 int32
	f64 float64
 )

// f64 + i   // Not allowed, cannot add float to integer
f64 + float64(i) // Will change `i` to float64

// i32 + i   // Not allowed, `i32` and `i` are of different types 
i32 + int32(i) // Will change `i` to int32

type MyInt int // Create an alias type on int
var mInt MyInt

// myInt + i  // Not allowed, `myInt` is it's own type now
int(myInt) + i       // Change `myInt` to int
myInt + MyInt(i)  // Change `i` to MyInt
```

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Pointers
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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
### Errors
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Errors are first class types in Go
```go []
// Declare a variable
var err error

// Can be passed to and returned by functions
func errHandler(err error) error {
	return err
}


```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Error chaining
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Errors can be chained and joined.
```go []
var (  
    err1    = errors.New("someError")  
    err2    = errors.New("someError")  
    errWrap = fmt.Errorf("%s:%w", "module", err1)  
)

// What are the messages of err1, err2, errWrap ?

// True of False?
err1 == err2, 
errors.Is(err1, err2),
err1.Error() == err2.Error(),

err1 == errWrap,
errors.Is(errWrap, err1),
err1.Error() == errWrap.Error()

errors.Is(errWrap, err1)
errors.Is(err1, errWrap))
```
</grid>

note:
false, false, false
false,true,false
true, false

--

<grid drag="100 10" drop="0 0" align="left" >
### Structs
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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
	
<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

<grid drag="100 40" drop="0 10" align="left" justify-content="center">
A slice is a dynamically-sized, flexible view into the elements of an array. 

In simplified terms, it is a pointer to an array with length and capacity.
</grid>

<grid drag="100 50" drop="0 40" align="left" justify-content="center">
<!-- element class="fragment fade-in-then-out" data-fragment-index="1" -->
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

<grid drag="100 50" drop="0 40" align="left" justify-content="center">
<!-- element class="fragment fade-in-then-outout" data-fragment-index="2" -->
```mermaid
flowchart TB
    subgraph Slice_Header
        A[Pointer: 0x2004]
        B[Length: 3]
        C[Capacity: 4]
    end
    
    subgraph Array_in_Heap
        D[0x2000: 1]
        E[0x2004: 2]
        F[0x2008: 3]
        G[0x200C: 4]
        H[0x2010: 5]
    end
    
    A -- points to --> E
``` 
</grid> 

--

<grid drag="100 10" drop="0 0" align="left" >
### Slices 2
</grid>

<grid drag="100 85" drop="0 15" align="left" justify-content="center">
Slice can be extended to add new values.
```go []
// define a slice of size 3 and cap 5 and add 2 values
slice := make([]int, 3, 5)  
slice = append(slice, 1)
slice = append(slice, 2)

// This append will go beyond capacity 
// and will cause reallocation and copy
slice = append(slice, 3)
```

When building slices we can skip the bounds

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

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

--

<grid drag="100 10" drop="0 0" align="left" >
### Collection functions
</grid>

<grid drag="100 85" drop="0 15" align="left" justify-content="center">
You can get the size and capactity of a collection using built-in **_len_** and **_cap_**  functions.
```go []
// define a slice of size 3 and cap 5 and add 2 values
var (
	arr [5]int
	slice = arr[1:4]
	m = map[int]int{0:1} 
)

len(arr) // Returns 5
cap(arr) // Returns 5

len(slice) // Returns 3
cap(slice) // Returns 4

len(m) // Returns 1
// cap(m) // Maps don't have capacity

```

</grid>


---

# Flow control

--

<grid drag="100 10" drop="0 0" align="left" >
### If / else
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
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

// Forever
for {}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Loops 2
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- We can also iterate through arrays, slices and maps

```go []
arr  := [1]byte{'a'}
m := map[string]string {
	"key1": "value1",
}

for idx, val := range arr {
	fmt.Println(idx, val) // Will print "0 a"
}
for key, val := range map {
	fmt.Println(key, val) // Will print "key1 value1"
}

// Short form - only indexes, keys
for idx := range arr {}
for key := range m {}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Switch
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- No default fallthrough - will only execute selected case
- Switch cases can be any value or expression

```go []
var name string
switch name {
	case "John":
		// Do something
	case "Jane":
	case "TestUser":
		fallthrough
	default:
		// Handle default
}

today := time.Now().Weekday()
switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Switch 2
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Switch does not need condition

```go []
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Functions
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">

```go []
func name(arg1 string, arg2 int64) int64 {
	return 0
}

func joinTypes(str1, str2 string) int64

func multipleReturns() (int64, error)

func noReturns() {}

func ignoredArg(_ int64) {}
```

Functions can have named returns

```go[]
func named() (i int64, err error) {
	if true {
		return 1, nil
	}
	return // If no values are provided it will return current states of i, err
}

```


</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Passing arguments
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
We have 2 ways of passing data to a function: 
- by value
- by reference

```go []
func byValue(a int) {}
func byRef(a *int) {}

v := 0

byValue(a) // Will copy value to function argument
byRef(&a) // Will pass a pointer to value

```
Can make a big impact when passing large structures

Things always passed by reference:
- arrays - (pointer, length)
- slices - (pointer, length, capacity)
- maps - (pointer)
- strings - (pointer, length)
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Struct methods
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
You can add methods to any type you have declared.

```go []
type MyType struct {}

// Receiver (m) is passed by copy 
// you will not be able to modify the structure
func (m MyType) MyMethod() {}

// Receiver is passed by reference
func (m *MyType) MySecondMethod() {}

// You can also alias existing types to add your own methods
type myInt int

func (m myInt) IsNegative() bool { return m < 0 }
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Interfaces
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go has **_implicit_** interfaces. It means that you don't have to state that a struct implements
an interface. It is checked automatically based on implemented methods.

```go []
type Reader interface {
	Read(p []byte) (n int, err error)
}

type File struct {}

func (f *File) Read(p []byte) (int, error) {
	// Read from file and put data in `p`
}

func ReadAll(r Reader) (data []byte, err error) {
	_, err = r.Read(data)
	return data, err
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Interfaces - Good practice
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
_Accept interfaces, Return Structs_

```go []
type Reader interface {
	Read(p []byte) (n int, err error)
}

type File struct {}

func NewFileReader(path string) *File {}

func ReadAll(r Reader) (data []byte, err error) {}
```

Allows for greatest flexibility

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Defer
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Will be called after function returns
- Function calls are pushed onto a stack (LIFO order)

```go []
defer fmt.Println("!")
defer fmt.Print("world")

fmt.Print("Hello ")

// Will print "Hello world!"
```
</grid>

---

# Code structure

--

<grid drag="100 10" drop="0 0" align="left" >
### Packages
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- All source files in the same folder must have the same package
- File with **_main_** function must belong to package **_main_**
- They can be included using **_import_** directive

```go []
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Folders
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```[]
root
|- cmd         - main packages for executables, with subdirectories for each binary
|- internal    - packages only intended for use within the current project
|- pkg         - library code that can be imported by external projects
|- api         - API definitions (protobuf files, OpenAPI etc.)
|- web         - static assets and templates for web apps
|- conf        - config files
|- scripts     - build & deployment scripts, automation
|- build       - CI files (Dockerfiles, Makefiles)
|- docs        - documentation
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Public vs Private
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
All files within the same package have access to all information within the package.

For other packages you can only access **_exported_** values.

Values starting with uppercase letter are exported.

```go []
const packetTTL = time.Second * 3600 // private value
const PacketTTL = time.Second * 3600 // public value
```

The same goes for functions and types.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Main file
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1-2|4-9|11-14]
// Declare main package
package main

// Import IO related libraries
import (
	"fmt"   // Formatted I/O
)

// Program entrypoint
func main() {
	fmt.Println("Hello World!")
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Tests
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1|3-5|8|9-12|13|14-20]
package main // or main_test

import (
	"testing" // Test types and functions
)

// Function must start with `Test*` and accept only `*testing.T`
func TestFibonacci(t *testing.T) {  
    type testCase struct {  
       input int  
       want  int  
    }  
    cases := []testCase{{...}}  
    for _, tc := range cases {  
       t.Run(fmt.Sprintf("Fibonacci(%d)", tc.input), func(t *testing.T) {  
          if got := Fibonacci(tc.input); got != tc.want {  
             t.Errorf("Fibonacci(%d) = %d, want %d", tc.input, got, tc.want)  
          }  
       })  
    }  
}
```
</grid>

---

# Tooling

--

<grid drag="100 10" drop="0 0" align="left" >
### Go tooling
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Go provides many tools through the **_go_** command

```bash
go build   # building executables
go run     # build and run 
go test    # run tests / benchmarks
go mod     # manage modules
go get     # download a package
go fmt     # format go code
go help    # find out more about a command
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Important for us
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```bash
# Initialize new module
go mod init {name}

# Run code in current folder (must contain main package with main function)
go run .

# Test code
go test       # run all tests in current folder
go test ./... # run all tests in current folder and subfolders

# Formatting
go fmt # format current folder

```
</grid>


---
# Exercises

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 1
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write a function called **_Fibonacci_** accepting an **_int_** and returning n-th fibonacci number.

```go []
// Sample calls
Fibonacci(0) // returns 0
Fibonacci(1) // returns 1
Fibonacci(2) // returns 1
Fibonacci(3) // returns 2
```

Check if your code passes tests.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 2
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write a function called **_AtoI_** accepting a **_string_** representation of a number and returning an **_int_**

You can assume that the passed value will be correct.

```go []
// Sample calls
AtoI("0") // returns 0
AtoI("123") // returns 123
AtoI("-432") // returns -432
```

Check if your code passes tests.

Hints:
- You can treat string as an array.
- You can subtract bytes to get digit value. ('1' - '0' == 1)

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 3
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write a function called **_ArrStats_** accepting an **_[]int_** and returning 3 values:
- min int
- max int
- average float64

```go []
// Sample calls
ArrStats([]int{0,1,2,3,4}) // returns (0, 4, 2.5)
ArrStats([]int{}) // returns (0, 0, 0)
```

Check if your code passes tests.
</grid>


--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 4
</grid>



<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write a function called **_MergeMap_** accepting 2 **map[string]string** and returning one with keys from both. Keys from first map takes precedence.

```go []
// Sample calls
MergeMap(
	map[string]string{"k1": "v1"},
	map[string]string{"k2": "v2"},
) // returns map[string]string{"k1":"v1", "k2":"v2"}

MergeMap(
	map[string]string{"k1": "v1", "k2":"v2a"},
	map[string]string{"k2": "v2b", "k3":"v3"},
) // returns map[string]string{"k1":"v1", "k2":"v2a", "k3":"v3"}
```

Check if your code passes tests.
</grid>

---

# Back to Devbook

--

<grid drag="100 10" drop="0 0" align="left" >
### Plan for today
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
We want to have a simple TUI app allowing us to:

1. Store items in memory
2. Operate on items:
	1. Add
	2. Remove
	3. List
3. Basic terminal based UI
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Create a folder for the project 

```
mkdir devbook
cd devbook
go mod init devbook
```

Initialize a module with **_go.mod_** file

```
module devbook

go 1.25.3
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### go.mod & go.sum
</grid>

<grid drag="100 10" drop="0 15" align="left" justify-content="center">
Introduced in Go 1.11 for managing external dependencies
</grid>

<grid drag="100 85" drop="0 25" align="left" justify-content="top">
<!-- element class="fragment fade-out" data-fragment-index="1" -->

**_go.mod_** stores module information

```
module devbook  

go 1.25

require (
        github.com/oklog/ulid/v2 v2.1.0
)

require (
        golang.org/x/text v0.18.0 // indirect
)

```

</grid>


<grid drag="100 85" drop="0 25" align="left" justify-content="top">
<!-- element class="fragment fade-in" data-fragment-index="1" -->
**_go.sum_** stores hashes of dependencies for validation
```
github.com/oklog/ulid/v2 v2.1.0 h1:+9lhoxAP56we25tyYETBBY1YLA2SaoLvUFgrP2miPJU=
github.com/oklog/ulid/v2 v2.1.0/go.mod h1:rcEKHmBBKfef9DhnvX7y1HZBYxjXb0cP5ExxNsTT1QQ=
golang.org/x/text v0.18.0 h1:XvMDiNzPAl0jr17s6W9lcaIhGUfUORdGCNsuLmPG224=
golang.org/x/text v0.18.0/go.mod h1:BuEKDfySbSR4drPmRPG/7iBdf8hvFMuRexcpahXilzY=

```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Main file
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1-2|4-9|11-14]
// Declare main package
package main

// Import IO related libraries
import (
	"fmt"   // Formatted I/O
	"bufio" // Buffered I/O - useful for user input
	"os"    // OS specific Stdin and Stdout
)

// Program entrypoint
func main() {
	fmt.Println("Hello World!")
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Item structure and store
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1-7|9-21]
// Item definition
type Item struct {  
    Name        string  
    Description string  
  
    URL string  
}  

// Simple in-mem store with starting items
var store = map[string]Item{  
    "item1": {  
       Name:        "item1",  
       Description: "search engine",  
       URL:         "https://google.pl",  
    },  
    "item2": {  
       Name:        "item2",  
       Description: "other search engine",  
       URL:         "https://bing.com",  
    },  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Handling User input
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1|3-11|13-21]
const prompt = "> "

func readLine() (string, error) {  
    scanner := bufio.NewScanner(os.Stdin)  
    scanner.Scan()  
    if err := scanner.Err(); err != nil {  
       return "", err  
    }  
  
    return scanner.Text(), scanner.Err()  
}

func AskForString(question string) (string, error) {  
    _, err := fmt.Printf("%s\n%s", question, prompt)  
    if err != nil {  
       return "", err  
    }  
  
    answer, err := readLine()  
    return strings.TrimSpace(answer), err
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Asking for a choice
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func AskForMapChoice(question string, choices map[string]string) (string, error) {  
    // Print question  
    _, err := fmt.Println(question)  
    if err != nil {  
       return "", err  
    }  
  
    // Check for longest key for formatting  
    maxLength := 0  
    for key := range choices {  
       if len(key) > maxLength {  
          maxLength = len(key)  
       }  
    }  
  
    // Sort by key  
    type choiceStruct struct {  
       key   string  
       value string  
    }
      
    sortedChoices := make([]choiceStruct, 0, len(choices))  
    for k, v := range choices {  
       sortedChoices = append(sortedChoices, choiceStruct{key: k, value: v})  
    }  
    sort.Slice(sortedChoices, func(i, j int) bool { return sortedChoices[i].key < sortedChoices[j].key })  
  
    // Print choices  
    for _, choice := range sortedChoices {  
       fmt.Printf("%*s: %s\n", maxLength, choice.key, choice.value)
    }  
  
    // Print prompt  
    fmt.Print(prompt)
  
    // Read answer  
    answer, err := readLine()  
    if err != nil {  
       return "", err  
    }  
  
    // Validate  
    if _, ok := choices[answer]; !ok {  
       fmt.Println("Invalid choice")  
       return AskForMapChoice(question, choices)  
    }  
  
    return answer, nil  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Menu "view"
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|2-8|11-14|16-34]
func menuView() error {  
    choice := map[string]string{  
       "a": "Add",  
       "e": "Edit", 
       "d": "Delete",  
       "l": "List",  
       "q": "Exit",  
    }  
  
    for {
       choice, err := AskForMapChoice("Choose option", choices)  
       if err != nil {  
          return err  
       }  
  
       switch choice {  
       case "a":  
          if err = addItemView(); err != nil {  
             return err  
          }  
       case "e":  
          fmt.Println("Not yet implemented")  
       case "d":  
          if err = deleteItemView(); err != nil {  
             return err  
          }  
       case "l":  
          printItems(ListItems())  
       case "q":  
          return nil  
	  default:
		  fmt.Println("Invalid choice")
       }  
       fmt.Println()  
    }  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Add / Remove / List "views"
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go [|1-17|19-30|32-35]
func addItemView() error {  
    name, err := AskForString("Name")  
    if err != nil {return err}  
  
    description, err := AskForString("Description")  
    if err != nil {return err}  
  
    url, err := AskForString("URL")  
    if err != nil {return err}  
  
    store[name] = Item{
	    Name:        name,  
		Description: description,  
		URL:         url,
	}
    return nil  
}  

func deleteItemView() error {  
    name, err := AskForString("Name")  
    if err != nil {return err}  
  
    if _, ok := store[name]; !ok {  
       fmt.Println("Item does not exist")  
       return nil
    }  
  
    delete(store, name)  
    return nil  
}

func printItems() {  
    for _, item := range store {  
       fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)  
    }  
}  
```
</grid>

---

# Refactor

--

<grid drag="100 10" drop="0 0" align="left" >
### Issues
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Mixing responsibilities in one file
	- Logic
	- User input
	- Storage
- Global store variable 
- Tight coupling

- Validation & Performance - not a focus
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Mixing responsibilities - problem
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- "Views" are handling user input and data storage
- Hard to test - tests have to get "interactive"
- Harder to reuse code
- Harder to change / refactor code
</grid>

--


<grid drag="100 10" drop="0 0" align="left" >
### Mixing responsibilities - solution
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Split into packages by domain:
```
devbook
|- cmd         - for our main file
|  |- main.go
|-pkg          - place for our packages
|  |- devbook   - our models
|  |  |- store.go
|  |  |- item.go
|  |- tui       - handling user interface & interaction
|  |  |- input.go
|  |  |- views.go
```

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Global Store - problem
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Hidden dependencies
- Testability issues
- Concurrency hazards
- Tight coupling
</grid>

note:
- **Hidden dependencies**
    - Functions may look pure but secretly depend on global state.
    - Harder to reason about behavior because inputs/outputs aren’t explicit.
- **Testability suffers**
    - Tests become order-dependent if they share the same global store.
    - Harder to run tests in parallel (common in Go) without state leaking between tests.
    - Mocking becomes awkward: you end up patching globals and carefully resetting them.
- **Concurrency hazards (especially in Go)**
    - Shared mutable global state can cause **data races** unless you guard it (mutexes/atomics/channels).
    - Even with locks, you can introduce deadlocks, contention, or subtle timing bugs.
- **Tight coupling and reduced modularity**
    - Many packages/modules become coupled to “the global store,” making refactors painful.
    - Harder to reuse components in other projects because they assume a global exists.
- **Initialization order and lifecycle problems**
    - “Who initializes the store, and when?” becomes a recurring bug source.
    - Cleanup is tricky (closing DB connections, flushing buffers, stopping goroutines).
    - In Go, `init()` + globals can create fragile startup sequences.


--

<grid drag="100 10" drop="0 0" align="left" >
### Global Store - solution
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Move to an independent structure
```go [|1|3-6|8-10|12-14|16-23|25-26|]
type Store map[string]Item  
  
func (store Store) GetByName(name string) (Item, bool) {  
    item, ok := store[name]  
    return item, ok  
}  
  
func (store Store) Add(item Item) {  
    store[item.Name] = item  
}  
  
func (store Store) DeleteByName(name string) {  
    delete(store, name)  
}  
  
func (store Store) List() []Item {  
    // Expressive way  
    arr := make([]Item, 0, len(store))  
    for _, item := range store {  
       arr = append(arr, item)  
    }  
  
    return arr  
  
    // Compact way  
    //return slices.Collect(maps.Values(store))
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Tight coupling - problem
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- View uses direct implementation
- Harder to change to different source of data
- Harder to test
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Tight coupling - solution
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Make views use an interface instead
```go []
type ItemStore interface {

}
```
</grid>

---

# Homework

--

<grid drag="100 10" drop="0 0" align="left" >
### Assignments
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Add editing items
- Complete all exercises
- Take a look at:
	- [Tour of Go](https://go.dev/tour/welcome/1) - great introduction to Go
</grid>

---