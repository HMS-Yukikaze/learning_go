/* for ix,value :=range slice1{
    ...
}
第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。*/
package main

import "fmt"

func main(){
    var slice1 []int = make ([]int,4)
    slice1[0] = 1
    slice1[1] = 2
    slice1[2] = 3
    slice1[3] = 4
    for i,val :=range  slice1{
        fmt.Printf("Slice at %d is:%d\n",i,val)
    }
    fmt.Printf("the other type of For-range\n")

    seasons :=  []string{"Spring","Summer","Autumn","Winter"}
    for ix,season :=range seasons{
        fmt.Printf("Season %d is :%s\n",ix,season)
    }
    
    fmt.Printf("the third type of For-range\n")
    // _ 用于忽略索引
    var season string 
    for _,season = range seasons{
        fmt.Printf("%s\n",season)
    }
    //如果只要索引，也可以忽略第二个变量
    fmt.Printf("ignore the second argument:\n")
    for ix := range seasons{
        fmt.Printf("%d\t",ix)
    }
    items := [...]int{10,20,30,40,50}
    for _,item :=range items{
        item *=2
        fmt.Println("item=",item)
    }


}
