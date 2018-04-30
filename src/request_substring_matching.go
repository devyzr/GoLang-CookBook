// taken from https://www.devdungeon.com/content/web-scraping-go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func main() {
    // Make HTTP GET request
    response, err := http.Get("https://www.devdungeon.com/")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Get the response body as a string
    dataInBytes, err := ioutil.ReadAll(response.Body)
    pageContent := string(dataInBytes)

    // Find a substr
    titleStartIndex := strings.Index(pageContent, "<title>")
    if titleStartIndex == -1 {
        log.Fatal("No title element found")
    }
    // We offset the title start index by 7 to skip over '<title>'
    titleStartIndex += 7

    // Find the index of the close tag
    titleEndIndex := strings.Index(pageContent, "</title>")
    if titleEndIndex == -1 {
        log.Fatal("No closing tag for title found")
    }

    // (Optional)
    // Copy the substring in to a separate variable so the
    // variables with the full document data can be garbage collected
    pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

    // Print the result
    fmt.Printf("Page title: %s\n", pageTitle)
}