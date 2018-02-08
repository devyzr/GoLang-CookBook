package main

import (
    "fmt"
    "strings"
)

func main() {
    // "normal" test case
    s := "test.string.one"
    sep := "."
    res := strings.Split(s, sep)
    fmt.Printf("Test 1: s(\"%s\") contains sep(\"%s\").\n", s, sep)
    fmt.Println(res)
    rangePrintLoop(res)

    // s does not contain sep and sep is not empty
    s = "TestStringTwo"
    sep = "."
    res = strings.Split(s, sep)
    fmt.Printf("Test 2: s(\"%s\") does not contain sep(\"%s\").\n", s, sep)
    fmt.Println(res)
    rangePrintLoop(res)

    // sep is empty
    s = "TestStringThree"
    sep = ""
    res = strings.Split(s, sep)
    fmt.Printf("Test 3: s(\"%s\") is not empty but sep(\"%s\") is.\n", s, sep)
    fmt.Println(res)
    rangePrintLoop(res)

    // both s and sep are empty
    s = ""
    sep = ""
    res = strings.Split(s, sep)
    fmt.Printf("Test 4: both s(\"%s\") and sep(\"%s\") are empty.\n", s, sep)
    fmt.Println(res)
    rangePrintLoop(res)

    // "normal" test case, except sep is first val of s
    s = ".test"
    sep = "."
    res = strings.Split(s, sep)
    fmt.Printf("Test 5: First char of s(\"%s\") is sep(\"%s\").\n", s, sep)
    fmt.Println(res)
    rangePrintLoop(res)    
}

func rangePrintLoop(res []string) {
    fmt.Println("Now testing with a range print loop:")
    for i, s := range res {
        fmt.Printf("Elem %d: %s\n", i, s)
    }
    fmt.Println("")
}