package main

import ."my_porto/Controlers"
import "fmt"
import "net/http"
import ."my_porto/Config"


func main() {


    http.Handle("/static/", 
        http.StripPrefix("/static/", 
            http.FileServer(http.Dir("styling"))))
    Connect()
	http.HandleFunc("/", Home)//page home 
    http.HandleFunc("/test", Test)//page home 
    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}

