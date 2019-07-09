package main

import "fmt"

func main(){
    const(
        Sunday=iota
        Monday
        Tuesday =9
        Wednesday
        Thursday=iota
        Friday
        Saturday
    )
    fmt.Println(Sunday)
    fmt.Println(Monday)
    fmt.Println(Tuesday)
    fmt.Println(Wednesday)
    fmt.Println(Thursday)
    fmt.Println(Friday)
    fmt.Println(Saturday)
}
