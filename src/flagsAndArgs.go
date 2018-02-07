package main

import "flag"

// Set string and boolean flag
var bFlag = flag.Bool("bFlag", false, "An example boolean flag")
var sFlag = flag.String("sFlag", "Empty", "An example string flag")

func main() {
    flag.Parse()
    println(*bFlag)
    println(*sFlag)
    // Iterate over the arguments after the flags.
    if len(flag.Args()) > 0 {
        for _, elem := range flag.Args() {
            println(elem)
        }
    }
}