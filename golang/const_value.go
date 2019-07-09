package main

import "fmt"

func  main(){
    const(
        a="abc"
        x     //常量组中未初始化的变量，视作和上一个变量相同
    )
    const Pi=3.14123124
    var n int
    fmt.Println((n + 5))
    fmt.Println(Pi)
    fmt.Println(a)
    fmt.Println(x)
}
