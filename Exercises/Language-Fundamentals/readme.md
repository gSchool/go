# Language Fundamentals

## Commentary

Go provices C-style /* */ block comments, and c++ style // line comments.


[Runnable Example](http://play.golang.org/p/LjhZArm6Cm)

```go
package main

func main() {
	helloWorld()
}

/*
   Special feature... putting a comments above top level declarations will be used
   in creating documentation with godoc
*/
func helloWorld() {
	// Print out hello world, and consequently the most useless comment ever made.
	println("hello world")
}
```

## Constants, Variables, Basic Types and Values

### Constants

Constants in Go are created at compile time, and can only be numbers, characters(runes), strings or booleans.
They can not refer to a function that needs to be called at compile time.

[Runnable Example](http://play.golang.org/p/lfWlj1XhsR)

```go
package main

import "fmt"

const (
	foo = "A"
	bar = 'A' // any guesses what this will be?
	bin = 2
)

func main() {
	fmt.Printf("%T %v\n", foo, foo)
	fmt.Printf("%T %v\n", bar, bar)
	fmt.Printf("%T %v\n", bin, bin)
	
	// And for fun
	fmt.Printf("%T %v\n", bar, string(bar))
	
}
```

You can create enumerated constants in Go by using the `iota` enumerator.  Iota starts at zero, and can be used as part of an expression.


[Runnable Example](http://play.golang.org/p/qopDSeeCe9)

```go
const (
	Apple int = iota
	Orange
	Banana
)
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

### Variables

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

### Basic Types and Values

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


### Conditionals

#### If Statement

[Runnable Example](http://play.golang.org/p/yEU809ox4S)

```go
	// Basic if statement
	if x > 0 {
		fmt.Println(x)
	} else {
		fmt.Println("x was not greater than zero")
	}

	// If statement with initialization
	if f, err := openFile(""); err != nil {
		fmt.Printf("Error opening file: %q\n", err)
	} else {
		fmt.Printf("Opened file %v\n", f)
	}
```

#### Switch Statement

[Runnable Example](http://play.golang.org/p/_abBjoFA2y)

```go
	// Basic switch statement
	switch tag {
	default:
		fmt.Println("Couldn't match anything")
	case 1:
		fmt.Println("I'm one!")
	case 3:
		fmt.Println("I'm three!")
	}
```

[Runnable Example](http://play.golang.org/p/9JOnGaJfYr)

```go
	switch x := f(); {
	case x < 0:
		fmt.Println(-x)
	default:
		fmt.Println(x)
	}
```
[Runnable Example](http://play.golang.org/p/7mIn6oqPI2)

```go
	switch t := v.(type) {
	default:
		fmt.Printf("Unexpected type %T\n", t)
	case string:
		fmt.Printf("%q is a string!\n", v)
	case int:
		fmt.Printf("%d is an int!\n", v)
	case bool:
		fmt.Printf("%v is a bool!\n", v)
	}
```

### Collections

  * Arrays
  * Slices
  * Maps

### Looping

[Runnable Example](http://play.golang.org/p/6QoMGUMF8-)

```go
	for i := 5; i >= 0; i-- {
		fmt.Printf("Is %d odd? %v\n", i, odd(i))
	}
```

  * Traditional loop for i:=0;... {
  * Infinite Loop for {
  * Conditional Loop for <bool> {
  * Range Loop

### Functions

  * Double(int) int
  * Triple(int) int
  * anonymous

### Scope

### Goroutines

### Channels
