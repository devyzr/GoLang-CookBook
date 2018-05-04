// taken from https://www.devdungeon.com/content/web-scraping-go
// http_request_with_cookie.go

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
    if err != nil {
        log.Fatal(err)
    }

    // Create a new cookies with the only required fields
    myCookie := &http.Cookie {
        Name: "cookieKey1",
        Value: "value1",
    }

    // Add the cookie to your request
    request.AddCookie(myCookie)

    // Ask the request to tell us about itself,
    // just to confirm the cookie attached properly
    fmt.Println(request.Cookies())
    fmt.Println(request.Header)
}