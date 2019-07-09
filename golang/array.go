package main

import std "fmt"

func main(){
    var a [5]int
    std.Println("emp:",a)
    
    a[4]=100
    std.Println("set:",a)
    std.Println("get:",a[4])
    std.Println("len",len(a))
    b :=[5]int{1,2,3,4,5}
    std.Println("dcl:",b)

    var twoD [2][3]int
    for i := 0;i<2;i++{
        for j:=0;j<3;j++ {
            twoD[i][j]=i+j
            std.Println("woD[i][j]",twoD[i][j])
        }
    }
    std.Println("2d:",twoD)
}
