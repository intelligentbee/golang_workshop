package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please specify a file path")
        return
    }

    filepath := os.Args[1]

    content, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
        return
    }

    fmt.Println(string(content))
}
