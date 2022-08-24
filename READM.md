# github.com/jainachal03/go-lib

## Usage

### Initialize your module
    $ go mod init example.com/my-golib-demo

### Get the go-lib module
Note that you need to include v in the version tag

    $ go get github.com/jainachal03/go-lib@v0.1.0

    package main
    import (
        "fmt"

        "github.com/jainachal03/go-lib"
    )

    func main(){
            fmt.Println(lib.Add(1, 2))
            fmt.Println(lib.Sub(2, 3))
    }

### Testing
    $ go test
