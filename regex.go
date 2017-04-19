//Demonstrate the regular expression parser

package main

import (
    "./regex"
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("examples/regex.txt")
    if err != nil {
        panic(1)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    //Get first line, the regex
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ":")
        reg := line[0]
        examples := line[1:]

        nfal := regex.ParseRegex(reg, "")
        fmt.Printf("Regex: %s\n", reg)
        for _,example := range examples {
            accept, _ := nfal.DeltaFunc(example, new(regex.Trace))
            fmt.Printf("\tExample: %s, %v\n", example, accept)
        }
    }


}
