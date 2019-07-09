package main

import(
    "fmt"
    "strings"
)
func main(){
    var str string = "123456789"
    var b []byte
    b = make([]byte,20)
    p:=strings.NewReader(str)
    fmt.Println(p)
    buf,err := p.Read(b)
    if err !=  nil{
        fmt.Println(err)
    }
    fmt.Println(buf)
}
