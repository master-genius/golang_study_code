package main

import (
    "net/http"
    "log"
)

func main() {
    start_http_server()
}

func start_http_server() {
    http.HandleFunc("/index",handler_main)
    http.HandleFunc("/",handler_main)
    http.HandleFunc("/info",handler_info)
    http.HandleFunc("/test",handler_test)

    fh := http.FileServer( http.Dir("/home/wy/go/resource/") )
    http.Handle("/resource/", http.StripPrefix("/resource/",  fh) )

    err := http.ListenAndServe(":8080",nil)
    if err!=nil {
        log.Fatal("ListenAndServe: ",err)
    }
}

func handler_main(rw http.ResponseWriter, req *http.Request) {
    var html string = "<p>Hello! This is the main page.</p>"
    html += "<p><a href=\"/info\">Show Info</a></p>"+
        "<p><a href=\"/test\">Show test</a></p>"
    rw.Write([]byte(create_html_page(html)))
}

func handler_info(rw http.ResponseWriter, req *http.Request) {
    var info string = "<p>Name: Bravewang<br>Job: programmer<br>"+
            "Email: 1146040444@qq.com</p>"+
            "<a href=\"/index\">Home</a>"
    rw.Write([]byte(create_html_page(info)))
}

func handler_test(rw http.ResponseWriter, req *http.Request) {
    var info string = "<a href=\"/index\">Home</a>"+
            "<p><h3>Test page</h3></p>"+
            "<p>I am a programming master.</p>"+
            "<p>hahahahahahahahahahaha.Hello world."+
            "abcdefghijklmn</p>"
    rw.Write([]byte(create_html_page(info)))
}

func create_html_page(content string) string {
    var header string = "<!DOCTYPE html><html><head><meta charset=\"utf-8\">"
    header += "<meta name=\"viewport\" content=\"width=device-width\">"
    header += "<script src=\"/resource/brutal_1709.js\"></script>"
    header += "<link href=\"/resource/fdflex.min.css\" rel=\"stylesheet\" type=\"text/css\">"
    header += "<style>a{color:#009acd;text-decoration:none;}</style>"
    header += "</head>"

    header += "<body><div class=\"row\"><div class=\"column small-12\">"
    var footer string = "</div></div><div class=\"row\">"+
            "<div class=\"column small-12\" id=\"test-info\"></div></div>"+
        "<script>brutal.autod('#test-info','Hello,this is brutal.js test info.');</script>"+
        "</body></html>"
    return header + content + footer
}
