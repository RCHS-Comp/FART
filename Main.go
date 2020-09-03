package main

import (
    "fmt"
    "time"
    "strings"
)

func main() {
    Error, Lines := OpenFile(FileArg())
    if Error != nil {
        Exit()
    }
    
    var Split []string
    Location := 0
    
    fmt.Println("(c) Harry Nelsen 2020")
    fmt.Println("The format is as follows:")
    fmt.Println("A      ->   A,A,A")
    fmt.Println("A,B    ->   A,B,B")
    fmt.Println("A,B,C  ->   A,B,C\n")
    
    StartTime := time.Now()
    for _, Line := range Lines {
        Split = strings.Split(Line, " ")
        Location = in_array(Types, Split[0])
        if !(len(Split) != 2 || Location == -1) {
            Solve(Location + 1, Split[1])
        }
    }
    EndTime := time.Now()
    fmt.Printf("Finished in %s\n", EndTime.Sub(StartTime))
}
