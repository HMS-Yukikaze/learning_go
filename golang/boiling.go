package main

import "fmt"


func main(){
    const freezingF,boilingF =32,212.0
    var f =boilingF
    var c = (f-32)*5/9
    fmt.Printf("boiling=%g F =%g C\n",f,c)
    fmt.Printf("%g F =%g C\n",boilingF,fToC(boilingF))
}

func fToC(f float64) float64{
    return (f-32)*5/9
}
