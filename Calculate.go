package main

import (
    "fmt"
    "strings"
    "math"
    "strconv"
)

var AllowedCharacters = "1234567890."
var Types = []string{"equilateral", "isosceles", "scalene"}

func in(Source string, LookingFor rune) (int) {
    for Index, Letter := range Source {
        if LookingFor == Letter {
            return Index
        }
    }
    
    return -1
}

func in_array(Source []string, LookingFor string) (int) {
    for Index, Element := range Source {
        if LookingFor == Element {
            return Index
        }
    }
    
    return -1
}

func outliers(SourceA string, SourceB string) (int) {
    Counter := 0
    
    for _, Letter := range SourceB {
        if in(SourceA, Letter) == -1 {
            Counter ++
        }
    }
    
    return Counter
}

func Solve(WantNumbers int, CSV string) {
    CSV = strings.Replace(CSV, " ", "", -1)
    Numbers := strings.Split(CSV, ",")
    
    Error := 0
    
    if (WantNumbers > 3 || WantNumbers < 1) {
        Error ++
    }
    
    for _, Number := range Numbers {
        if strings.Contains(Number, "..") {
            Error += 2
        }
    }
    
    if len(Numbers) != WantNumbers {
        Error += 4
    }
    
    if Error != 0 {
        fmt.Printf("Error code %d\n", Error)
    }
    
    var A float64
    var B float64
    var C float64
    
    switch WantNumbers {
        case 1:
            A, _ = strconv.ParseFloat(Numbers[0], 32)
            B = A
            C = A
        case 2:
            A, _ = strconv.ParseFloat(Numbers[0], 32)
            B, _ = strconv.ParseFloat(Numbers[1], 32)
            C = B
        case 3:
            A, _ = strconv.ParseFloat(Numbers[0], 32)
            B, _ = strconv.ParseFloat(Numbers[1], 32)
            C, _ = strconv.ParseFloat(Numbers[2], 32)
    }
    
    Permimeter := (A + B + C)
    Semipermimeter := (Permimeter / 2)
    Area := math.Sqrt(Semipermimeter * (Semipermimeter - A) * (Semipermimeter - B) * (Semipermimeter - C))
    
    fmt.Printf("Type: %s\n", Types[WantNumbers - 1])
    fmt.Printf("Area: %f\n", Area)
    fmt.Printf("Permimeter: %f\n\n", Permimeter)
}
