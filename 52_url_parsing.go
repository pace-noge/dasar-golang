package main

import (
    "fmt"
    "net/url"
)


func main() {
    var urlString = "http://www.google.com/hello?name=john wick&age=27"
    var u, e = url.Parse(urlString)
    if e != nil {
        fmt.Println(e.Error())
        return
    }

    fmt.Printf("url: %s\n", urlString)
    fmt.Printf("protocol: %s\n", u.Scheme)  // http
    fmt.Printf("host: %s\n", u.Host)    // www.google.com
    fmt.Printf("path: %s\n", u.Path)  // /hello

    var name = u.Query()["name"][0] // john wick
    var age = u.Query()["age"][0]
    fmt.Printf("name: %s, age: %s\n", name, age)
}
