package main

import (
    "crypto/md5"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    data := "iwrupvqb"
    targetPrefix := "00000"
    var hash [16]byte
    counter := 1

    for {
        newData := data + strconv.Itoa(counter)
        
        // Compute the MD5 hash
        hash = md5.Sum([]byte(newData))
        
        // Convert the hash to a hexadecimal string
        hashString := fmt.Sprintf("%x", hash)
        
        // Check if the hash starts with the desired prefix
        if strings.HasPrefix(hashString, targetPrefix) {
            fmt.Printf("Found hash: %s\n", hashString)
            fmt.Printf("With counter: %d\n", counter)
            break
        }

        // Increment the counter for the next iteration
        counter++
    }
}
