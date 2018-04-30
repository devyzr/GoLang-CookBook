package main

import (
	"flag"
	"fmt"
)

var bigFlag = flag.Bool("Flag", false, "The big flag!")

func init() {
    // Note that a help message is necessary and both will be printed individually.
	flag.BoolVar(bigFlag, "f", false, "The small flag!")
}

func main() {
    flag.Parse()
	fmt.Println(*bigFlag)
}
