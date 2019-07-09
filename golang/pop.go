package main

import "fmt"

func main(){
    s :=make([]int,3,10)
    fmt.Println(s)
    fmt.Println(len(s),cap(s))
    ar:=[5]int{3,2,1,4,5}
//  外层控制遍历的次数，内层循环进行交换
    for i:=0;i<len(ar)-1;i++{
        fmt.Printf("第%d次遍历\n",i)
        for j:=0;j<(len(ar)-i-1);j++{
            if ar[j]<ar[j+1]{
                fmt.Printf("ar[%d]=%d < ar[%d]=%d\n",j,ar[j],j+1,ar[j+1])
                temp:=ar[j]
                ar[j]=ar[j+1]
                ar[j+1]=temp
            }
        }
    }
    for i:=0;i<len(ar);i++{
        fmt.Println(ar[i])
    }
}
