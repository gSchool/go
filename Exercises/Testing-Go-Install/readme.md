#Test your installation

Check that Go is installed correctly by building a simple program, as follows.

Create a file named `hello.go` and put the following program in it:

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
Then run it with the go tool:

```bash
$ go run hello.go
hello, world
```

If you see the "hello, world" message then your Go installation is working.
