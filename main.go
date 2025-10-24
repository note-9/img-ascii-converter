package main

import (
    "fmt"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

func main() {
    filePath := "~/Pictures/girl.jpg"
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file %v\n", err)
        return
    }
    
}
