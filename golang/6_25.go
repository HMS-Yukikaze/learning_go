package main

import "fmt"

const (
    _ = iota
    KB float64= 1<<(iota *10)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
func main(){
    var a int =65
    b := string (a)
    fmt.Println(b)
    fmt.Println(KB)
    fmt.Println(MB)
    fmt.Println(GB)
    fmt.Println(PB)
    fmt.Println(PB)
}
