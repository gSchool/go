# Installing Go

Download the latest installer from go.  This is the preferred method.

[http://golang.org/dl/](http://golang.org/dl/)

## Alternate Install Methods - Not preferred

*Using Homebrew:*

```bash
brew update
brew install go
```

*Using APT*

```bash
sudo apt-get update
sudo apt-get install golang
```

*Note: This can leave you with an old version of go as they don't update
the package manager as timely as they should*

# Environment - GOPATH

The GOPATH environment variable specifies the location of your
workspace. It is likely the only environment variable you'll need to set
when developing Go code.

To get started, create a workspace directory and set GOPATH accordingly.
Your workspace can be located wherever you like, but we'll use $HOME/go
in this document. Note that this must not be the same path as your Go
installation.

*NOTE:* This is only necessary if the method of installation used above did NOT do this.

## Instructions for Linux or Mac

```bash
mkdir $HOME/go
export GOPATH=$HOME/go
```

## Instructions for Windows

From a command prompt:

```bash
makedir "%USERPROFILE%\go"
```

Go to the Control Panel > Sytem > Advanced Tab > Environment Variables.

Add a new User Variable

Variable name: GOPATH
Variable value: %USERPROFILE%\go

# Add Go Bin directory to PATH

For convenience, add the workspace's bin subdirectory to your PATH:

*NOTE:* If you used an installer, you typically do not need to do this step.

```bash
export PATH=$PATH:$GOPATH/bin
```


