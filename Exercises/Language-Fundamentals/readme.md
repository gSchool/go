# Language Fundamentals

## Commentary

Go provices C-style /* */ block comments, and c++ style // line comments.

```go
package main

func main() {
	HelloWorld()
}

/*
   Special feature... putting a comments above top level declarations will be used
   in creating documentation with godoc
*/
func HelloWorld() {
	// Print out hello world, and consequently the most useless comment ever made.
	println("hello world")
}
```

## Constants, Variables, Basic Types and Values

## Constants

Constants in Go are created at compile time, and can only be numbers, characters(runes), strings or booleans.
They can not refer to a function that needs to be called at compile time.

[Runnable Example](http://play.golang.org/p/QeoJvxhBs4)

```go
package main

import "fmt"

const (
	foo = "1"
	bar = '2'
	bin = 2
)

func main() {
	fmt.Printf("%T %v\n", foo, foo)
	fmt.Printf("%T %v\n", bar, bar)
	fmt.Printf("%T %v\n", bin, bin)
}
```

File modes from [golang.org/pkg/os/#FileMode](http://golang.org/pkg/os/#FileMode)

```go
const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file (not backed up)
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky

        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

        ModePerm FileMode = 0777 // permission bits
)
```

## Variables

Variables are initialized just like constants, but are computed at run time.

[Runnable Example](http://play.golang.org/p/6-ALpAL5_q)

```go
package main

import "os"
import "path/filepath"

func main() {
	var (
		gopath   = os.Getenv("GOPATH")
		dataPath = filepath.Join(gopath, "data")
	)
	println(dataPath)
}
```

## Basic Types and Values

For an in depth look at types, see the [Language Specification for Types](http://golang.org/ref/spec#Types)

Basic types include:

* Boolean
* Numeric
* String
* Array
* Slice
* Struct
* Pointer
* Function
* Interface
* Map
* Channel

All types have a default value.


[Runnable Example](http://play.golang.org/p/pWN9tbcHBL)

```go
package main

import "fmt"

func main() {
	var (
		s string
		b bool
		i int
	)
	fmt.Printf("%q\n", s)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", i)
}
```


* Conditionals
  * If/Else
  * Switch
    * switch value { }
    * switch { case expression: ... }

* Collections
  * Arrays
  * Slices
  * Maps

* Looping
  * Traditional loop for i:=0;... {
  * Infinite Loop for {
  * Conditional Loop for <bool> {
  * Range Loop

* Functions
  * Double(int) int
  * Triple(int) int
  * anonymous

* Scope

* Goroutines

* Channels
