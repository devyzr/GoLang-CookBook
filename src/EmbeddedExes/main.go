// Unpacks a microsoft calculator from the 
// neighbouring 'assets.go' file.
// Both packages must be called main.
// Doesn't work with 'run', must be built.
package main

import ("fmt"
        "io/ioutil"
        )

func main() {
    exe_path := "calc.exe" 
    fmt.Println("Writing Omnibus script to %s\n", exe_path) 
    data, err := Asset("assets/calc.exe") 
    if err != nil { 
        fmt.Println("Problem loading asset.")
    }
    
    err = ioutil.WriteFile(exe_path, []byte(data), 0644) 
    if err != nil {
        fmt.Println("There was a problem writing the .exe")
    } else {
        fmt.Println("Exe written successfully!")
    }
}