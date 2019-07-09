package main

import "fmt"

func (st *Stack) Pop() int {
    v := 0
    for ix := len(st) - 1; ix >= 0; ix-- {
        if v = st[ix]; v != 0 {
            st[ix] = 0
            return v
        }
    }
}    
func main() {
    fmt.Println("Hello, 世界")
    fmt.Println("go"+"ole")
    fmt.Println("7.0/3.0=",7.0/3.0)
}
