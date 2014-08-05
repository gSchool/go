# The `go` tools

## The `go` command

The `go` command is a tool for managing Go source code.

Usage:
        go command [arguments]

The commands are:

    build       compile packages and dependencies
    clean       remove object files
    env         print Go environment information
    fix         run go tool fix on packages
    fmt         run gofmt on package sources
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         run go tool vet on packages


Use `go help [command]` for more information about a command.

Additional help topics:

    c           calling between Go and C
    filetype    file types
    gopath      GOPATH environment variable
    importpath  import path syntax
    packages    description of package lists
    testflag    description of testing flags
    testfunc    description of testing functions

Use `go help [topic]` for more information about that topic.


Difference between the naked call:

```bash
go fmt
```

And the recursive call
```bash
go fmt ./...
```

### go fmt

Usage:

```bash
go fmt [-n] [-x] [packages]
```

Fmt runs the command 'gofmt -l -w' on the packages named by the import paths.
It prints the names of the files that are modified.

For more about gofmt, see 'godoc gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

Simple usage:

```bash
go fmt
```

### go build

For an extensive reference, see the [doc](http://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

From the root of your package:

```bash
go build
```

### go test

For an extensive reference, see the [doc](http://golang.org/cmd/go/#hdr-Test_packages)

From the root of your package:

```bash
go test
```

### go vet

Usage:

```bash
go vet [-n] [-x] [packages]
```

Vet runs the Go vet command on the packages named by the import paths.

For more about vet, see 'godoc code.google.com/p/go.tools/cmd/vet'.
For more about specifying packages, see 'go help packages'.

To run the vet tool with specific options, run 'go tool vet'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

Simple usage:

```bash
go vet
```


