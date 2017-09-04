package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

func main() {
    defer cpu_runtime(time.Now())
    var file_path string = os.Getenv("HOME") + "/tmp/big_file.test"
    f,err := os.Stat(file_path)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }
    if !f.Mode().IsRegular() {
        fmt.Println("Error: file is not exists")
        return
    }

    data,err := ioutil.ReadFile(file_path)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(string(data))
}

func cpu_runtime(t time.Time) {
    fmt.Println( time.Since(t) )
}


