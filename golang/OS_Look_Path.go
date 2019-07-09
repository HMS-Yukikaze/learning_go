package main

import(
    "fmt"
    "log"
    "os/exec"
)

func main(){
    path,err := exec.LookPath("cat")
    if err !=nil{
        log.Fatal("error")
    }
    fmt.Printf("%s",path)
}

