package main

import "fmt"

func main(){
    var a="init"
    fmt.Println(a)
    var b,c int=1,2
    fmt.Println(b,c)
    fmt.Println("b+c=",b+c)
    var x int32 =12
    var y float32 =23.12
    var m int32
    m=x+int32(y)
    fmt.Println("int32(y)",int32(y))
    fmt.Println("m",m)
}
