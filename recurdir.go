package main

import (
    "fmt"
    "io/ioutil"
    //"path/filepath"
    "os"
)

func main() {
    var d string = "/usr"
    f_list := []string{d}
    var dir_len int =1
    var i int = 0
    for i<dir_len {
        tmp,_ := os.Stat(f_list[i])
        fmt.Printf("%s:\n",tmp.Name())
        if tmp.Mode().IsDir() {
            list_buf,err := ioutil.ReadDir(f_list[i])
            if err!=nil {
                fmt.Printf("Error:%s",err.Error())
            } else {
                for _,f := range list_buf {
                    fmt.Println(f.Name())
                    if f.Mode().IsDir() {
                        f_list = append(f_list,f_list[i]+"/"+f.Name())
                    }
                }
            }
        } else {
            fmt.Printf("%s is regular file\n",f_list[i])
        }
        i+=1
        dir_len = len(f_list)
        //fmt.Println(f_list,i)
    }
    /*
    f,_ := os.Stat(d)
    fmt.Println(f.Mode().IsDir())
    fmt.Println( filepath.Dir("/home/wy/bin/fst") )
    */
}

func recur_dir() {

}

