package main

import "fmt"

func main(){
    result:=0
    for i:=0;i<=10;i++{
        result=fun(i)
        fmt.Printf("fibonacii:%d\t is %d\n",i,result)
    }

}
func fun(num int)(res int){
    if num<=1{
        res=1
    }else{
        res=fun(num-1)+fun(num-2)
    }
    return
}
