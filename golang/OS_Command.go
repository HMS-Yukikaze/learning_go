package main


import(
    "bytes"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func main(){
    //func Command (name string ,arg ...string) *Cmd
    cmd := exec.Command("echo","hello")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    //func (c *Cmd) Run() error 执行Cmd中包含的命令，阻塞直到命令执行完成
    err := cmd.Run()
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("in all caps :%q\n",out.String())
}


