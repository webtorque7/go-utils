package main

import (
    // "os"
    "fmt"
    "github.com/kardianos/osext"
)


func updateSelf() {
    self, err := osext.Executable()
    checkError(err)
    fmt.Println(self)
}
