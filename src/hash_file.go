package main

import(
    "fmt"
    "crypto/sha256"
    "os"
    "log"
    "io"
)

func main() {
    test_file := "test_file.txt"
    f, err := os.Open(test_file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    h := sha256.New()

    _, err = io.Copy(h, f)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The hash of '%s' is:\n%x", test_file, h.Sum(nil))
}