package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }//调用
func untrace(s string) { fmt.Println("leaving:", s) }//追踪

func a() {
    trace("a")
    defer untrace("a")
    fmt.Println("in a")
}

func b() {
    trace("b")
    defer untrace("b")
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
