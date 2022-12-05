package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
)

func main() {
    // open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        fmt.Println(i)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}