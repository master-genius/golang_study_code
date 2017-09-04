package main

import (
    "fmt"
    "io/ioutil"
    //"path/filepath"
    "os"
    "time"
)

func main() {
    //defer cpu_run_time(time.Now())
    var d string = "/usr"
    var run_time int64
    var rt_float float64
    run_time = run_time_count(recur_dir,d)
    rt_float = float64(run_time)
    rt_float = rt_float/1000000000
    fmt.Println("run time:",rt_float," seconds")
    /*
    f,_ := os.Stat(d)
    fmt.Println(f.Mode().IsDir())
    fmt.Println( filepath.Dir("/home/wy/bin/fst") )
    */
}

func recur_dir(d string) {
    tmp,err := os.Stat(d)
    if err!=nil {
        fmt.Println(err.Error())
        return
    }
    if !tmp.Mode().IsDir() {
        fmt.Println(d,"is not a dir")
        return
    }

    f_list := []string{d}
    var dir_len int =1
    var i int = 0
    for i<dir_len {
        //fmt.Printf("%s:\n",f_list[i])
        list_buf,err := ioutil.ReadDir(f_list[i])
        if err!=nil {
            fmt.Printf("Error:%s",err.Error())
        } else {
            for _,f := range list_buf {
                //fmt.Println(f.Name())
                if f.Mode().IsDir() {
                    f_list = append(f_list,f_list[i]+"/"+f.Name())
                }
            }
        }
        i+=1
        dir_len = len(f_list)
        //fmt.Println(f_list,i)
    }
}

func run_time_count(f func(string),args string) int64 {
    var start_time,end_time int64
    start_time = time.Now().UnixNano()
    f(args)
    end_time = time.Now().UnixNano()
    return (end_time - start_time)
}

func cpu_run_time(t time.Time) {
    rt := time.Since(t)
    fmt.Println(rt)
}

