package main

import (
    "bufio"
    "os"
)

func Exit() {
    os.Exit(1)
}

func FileArg() (string) {
    if (len(os.Args) > 1 && FileExists(os.Args[1])) {
        return os.Args[1]
    } else {
        return "Questions.txt"
    }
}

func OpenFile(FileName string) (error, []string) {
    var Output []string
    FileObj, Error := os.Open(FileName)
    if Error != nil {
        return Error, Output
    }
    
    defer FileObj.Close()
    
    ScanObj := bufio.NewScanner(FileObj)
    for ScanObj.Scan() {
        if (ScanObj.Text())[0] != '#' {
            Output = append(Output, ScanObj.Text())
        }
    }
    
    Error = ScanObj.Err()
    
    return Error, Output
}

func WriteFile(FileName string, Lines []string) (error) {
    FileObj, Error := os.Create(FileName)
    if Error != nil {
        return Error
    }
    
    defer FileObj.Close()
    
    for _, Line := range Lines {
        FileObj.WriteString(Line + "\n")
    }
    
    return Error
}

func FileExists(FileName string) (bool) {
    if _, Error := os.Stat(FileName); os.IsNotExist(Error) {
        return false
    } else {
        return true
    }
}
