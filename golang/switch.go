package main

import "fmt"


func main(){
    var a int64=11
    var b int64=22
    a,b = mult_returnval(a,b)
    fmt.Println(a)
    fmt.Println(b)
}
//func  函数名（arg1 类型，arg2 类型...）（返回参数列表，...）
func mult_returnval(i int64,j int64)(a2 int64,b2 int64){
    a2 = 2*i
    b2 = 3*j
    return a2,b2
    //return 
}

