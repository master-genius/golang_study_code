package main

import "fmt"
import "master/trial"


func main() {
    rt := trial.CPURuntime(loop_run,"qweou")
    fmt.Println(rt)
}

func loop_run(s string) {
    var i int = 0
    for i<10000;i++ {
        fmt.Println(s)
    }
}

