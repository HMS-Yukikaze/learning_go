/*Go 里面有三种类型的函数：

普通的带有名字的函数
匿名函数或者lambda函数
方法（Methods)*/
package main

import "fmt"

func main(){
    fmt.Printf("sum= %d\n",sum(1,2,3))
}
func sum(a int,b int,c int)int{
    return a+b+c
}
