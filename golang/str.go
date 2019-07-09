package main

import (
    "fmt"
    "strings"
)
func main(){
    s:="aaa"
    var str string ="this is an example of  a string"
    fmt.Println(s[0]=='\x61',s[1]=='b',s[2]==0x63)
    fmt.Printf("T/F? Does the string \"%s\" have prefix %s ?",str,"th")
    //HasPrefix 判断字符串 s 是否以 prefix 开头：
    //strings.HasPrefix(s, prefix string) bool
    fmt.Printf("%t\n",strings.HasPrefix(str,"th"))//判断是否以“th”开头

    //Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
    //strings.Index(s, str string) int
    fmt.Printf("%d\n",strings.Index(str,"a"))
    i:=strings.Count("is",str)
    fmt.Println("count:",i)
    fmt.Printf("%s\n",strings.Repeat("huaq",3))

    //分割字符串
    //str1 := "are you ok ?"
    s1 := strings.Fields(str)
    //fmt.Printf("After Field:%v\n",s1)
    //fmt.Printf("After Field:%s\n",s1)
    for _,val :=range s1{
        fmt.Printf("%s -",val)
    }
    //从字符串中读取
    s_str:=str
    fmt.Printf("org:%s\n",s_str)
    p :=strings.NewReader(str)
    fmt.Println("1.A",p)
}

