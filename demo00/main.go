package main

import (
    "fmt"
    "github.com/wuqingze/demo00/util"
    utilHello "github.com/wuqingze/demo00/util/hello"
    toolHello "github.com/wuqingze/demo00/tool/hello"
)

func main() {
    fmt.Println("Hello, world!")
    util.SayHello()
    utilHello.SayHello()
    toolHello.SayHello()
}
