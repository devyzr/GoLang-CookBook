// taken from https://www.devdungeon.com/content/web-scraping-go

package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main(){
    // Create HTTP Client with timeout
    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    // Create and modify HTTP request before sending
    response, err := client.NewRequest("GET", "https://www.devdungeon.com/", nil)
    if err != nil {
        log.Fatal(err)
    }
    request.Header.Set("User-Agent", "Not Firefox")

    // Make request
    response, err := client.Do(request)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Copy data from response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}