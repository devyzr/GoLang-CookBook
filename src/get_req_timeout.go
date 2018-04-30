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

    // Make request
    response, err := client.Get("https://www.devdungeon.com/")
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