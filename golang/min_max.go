package main

import "fmt"

func main(){
    a := 1
    b := 0
    var min int
    var max int
    min,max=compare(a,b)
    fmt.Println(min)
    fmt.Println(max)
}

func compare(a int,b int)(min int ,max int){
    if a>b {
        return a,b
    }else{
        return b,a
    }
}
