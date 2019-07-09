package main

import "fmt"

func f(a [3]int){ fmt.Println(a)}
func fp(a *[3]int){ fmt.Println(a) }

func main(){
    var ar [3]int
    var ar1 =new([5]int)//go语言中数组是一种值类型，所以可以通过new创建
    ar2 := *ar1
    ar2[2]=100
    fmt.Println("ar1",ar1)
    fmt.Println("ar2",ar2)
    ar[2]=1
    f(ar)
    fp(&ar)
/*
    Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建： var arr1 = new([5]int)。
    那么这种方式和 var arr2 [5]int 的区别是什么呢？arr1 的类型是 *[5]int，而 arr2的类型是 [5]int。
    这样的结果就是当把一个数组赋值给另一个时，需要在做一次数组内存的拷贝操作。
*/
}

