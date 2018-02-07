package main

import "flag"

// note, that variables are pointers
var boolFlag = flag.Bool("bool", false, "Description of flag")

func main() {
    flag.Parse()
    println(*boolFlag)
}     