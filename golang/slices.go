package main

import "fmt"

func main(){
    s:=make([]string,3)
    fmt.Println("emp:=",s)
    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:",s)
    fmt.Println("get:",s[2])
    fmt.Printf("len:%d ,cap:%d\n",len(s),cap(s))//切片越界会报panic错误
    fmt.Printf("-%p\n",s)
    s =append(s,"d")//用append实现在切片中追加元素
    fmt.Printf("--%p\n",s)//当容量不足，补齐容量后，切片的地址发生了改变
    fmt.Printf("*len:%d ,cap:%d\n",len(s),cap(s))//用append追加元素时，若容量不足，则自动补齐
    s =append(s,"e",`"f"`,"g")
    fmt.Println("append:",s)
    fmt.Printf("**len:%d ,cap:%d\n",len(s),cap(s))//补齐规则：加上当前cap（）个容量，即为当前不够翻一倍
    fmt.Printf("---%p\n",s)
    c := make([]string,len(s))
    copy(c,s)
    fmt.Printf("copy:%p\n",s)//copy（）复制切片时，地址未发生改变
    fmt.Println("cpy",c)
    
    l:= s[2:5]
    fmt.Println("sl1:",l)
    
    l =s[:5]
    fmt.Println("sl2:",l)

    l =s[2:]
    fmt.Println("sl3",l)
    
    t:=[]string{"g","h","i"}
    fmt.Println("dcl",t)

    twoD :=make([][]int,3)
    for i:=0;i<3;i++{
        innerLen :=i + 1
        twoD[i] =make([]int,innerLen)
        for j:=0;j<innerLen;j++{
            twoD[i][j]=i+j
        }
        fmt.Println("2d:",twoD)
    }
}
