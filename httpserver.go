package main

import (
    "net/http"
    "log"
)

func main() {
    start_http_server();

}

func start_http_server() {
    http.HandleFunc("/index",handler_main)
    http.HandleFunc("/",handler_main)
    http.HandleFunc("/info",handler_info)

    err := http.ListenAndServe(":8080",nil)
    if err!=nil {
        log.Fatal("ListenAndServe: ",err)
    }
}

func handler_main(rw http.ResponseWriter, req *http.Request) {
    rw.Write([]byte("Hello,world"))
}

func handler_info(rw http.ResponseWriter, req *http.Request) {
    var info string = "Name: Bravewang\nJob: programmer\nEmail: 1146040444@qq.com"
    rw.Write([]byte(info))
}
