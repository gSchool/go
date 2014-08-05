# Language Fundamentals

* Commentary

Go provices C-style /* */ block omments, and c++ style // line comments.

```go

package main

func main () {
	HelloWorld()
}

/*
	Special feature... putting a comments above top level declarations will be used 
	in creating documentation with godoc
*/
func HelloWorld() {
	// Print out hello world, and consequently the most useless comment ever made.
	println('hello world')
}
```

* Constants, Variables, Basic Types and Values

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
