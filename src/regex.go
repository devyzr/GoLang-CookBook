package main

import (
	"fmt"
    "regexp"
)

func main() {
    regExStr := "\\\"[a-zA-Z]+\\\""
    compiledReg, err := regexp.Compile(regExStr)

    if err != nil {
        fmt.Println("Error compiling regex!")
    } else {
        fmt.Printf("The folling tests will be scanned with this regex:%s\n", regExStr)

        var result string
        var test string

        test = "Stacy is awesome!"
        result = compiledReg.FindString(test)
        fmt.Println(test)
        fmt.Printf("Returns: %s\n", result)
        
        test = "Greg is \"cool\""
        result = compiledReg.FindString(test)
        fmt.Println(test)
        fmt.Printf("Returns: %s\n", result)

        test = "This '\"' is a single quote!"
        result = compiledReg.FindString(test)
        fmt.Println(test)
        fmt.Printf("Returns: %s\n", result)
    }
}
