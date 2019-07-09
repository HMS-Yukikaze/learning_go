package main

import "fmt"

func main(){
    fun1()
}

func fun1(){
    fmt.Println("head")
    /*defer*/ fun2()
    fmt.Println("end")
    fun3()
}

func fun2(){
    fmt.Println("this is function2")
}
//defer 栈：后进先出
func fun3(){
    for i:=0; i<5 ;i++{
        defer fmt.Println(i)
    }
}
/*
1.关闭文件流 
    open a file
    defer file.Close()
2.解锁一个加锁的资源 
    mu.Lock()
    defer mu.Unlock()
3.打印最终报告
    printHeader()  
    defer printFooter()
4.关闭数据库链接
    // open a database connection  
    defer disconnectFromDB()
*/

