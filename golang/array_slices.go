package main

import "fmt"

func main(){
    //数组：数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列；只有长度一致，类型一致的数组才能成为同一类型
    //数组声明的两种方式
    //1.  var 数组名 [len]type
    var ar [3]int
    //初始化的变型:
    var ar1 = [...]int{1,2,3}
    var ar2 = [3]int{4,5,6}
    var ar3 = []int{7,8,9}//注：这种形式没有指定大小，故，实际初始化得到的是切片
    //用for循环遍历初始化,打印数组中元素
    for i:=0;i<len(ar);i++{
        ar[i] = i*2
    }
    for i:=0;i<len(ar);i++{
        fmt.Println("ar[]=",ar[i])
    }
    //打印数组元素的第二种方式
    fmt.Println(&ar1)
    fmt.Println(&ar2)
    fmt.Println(&ar3)
    //切片：切片（slice）是对数组一个连续片段的引用 ,切片是一个长度可变的数组
    //声明切片的格式是： var identifier []type（不需要说明长度）

    var array1 [6]int //声明数组
    array1[0]=1
    array1[1]=2
    array1[2]=3
    array1[3]=4
    array1[4]=5
    array1[5]=6
    var slice1 []int =array1[2:5]
    fmt.Println("slice1",&slice1)//array1[2:5]不包括第五个元素，&slice1 -> &[3,4,5]
    //2   用make（）创建一个切片
    slice2 :=make([]int ,10)
    fmt.Println("caps=",cap(slice2))
    //2.2 切片在内存中的组织方式：1.指向原数组的指针，2.长度len（） 3.容量cap（）

}
