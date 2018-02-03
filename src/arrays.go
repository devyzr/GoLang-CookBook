package main

import (
    "fmt"
)

func main() {
    scores := make([]int, 0, 10)
    scores = scores[0:8]
    scores[7] = 9033
    fmt.Println(scores)
}

