package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"    
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
   
    time := s[:8] 
    ampm := s[8:]

    if s[:2] == "12" {
        if ampm == "AM" {
            return "00" + time[2:]
        } else {                    
            return time
        }                
    } else if (s[:2] == "00"){
        if ampm == "AM" {
            return time
        } else {                    
            return "12" + time[2:]
        }          
    } else {    
        if ampm == "AM" {
            return time
        } else {                    
            angkadepan, err := strconv.Atoi(s[:2])
            if err != nil {
                panic(err)
            }
            jam := angkadepan + 12
            
            return strconv.Itoa(jam) + time[2:]            
        }                        
    }       
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
s