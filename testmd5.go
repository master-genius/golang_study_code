package main

import "fmt"
import "crypto/md5"

func main() {
    var a string = "1234";
    m := md5.Sum([]byte(a))
    fmt.Printf("%x\n",m)
}
