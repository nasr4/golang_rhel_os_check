package main

import (
       "os"
       "fmt"
)

//os check for file path
func OsCheck() (OsVersion []byte, err error) {
        f, err := os.Open("/etc/redhat-release")
        //b1 := make([]byte, 60)
        //f.Read(b1)
        if err != nil {
                return nil, err
        }

        if _, err := f.Seek(40, 0); err != nil {
                return nil, err
        }

        b1 := make([]byte, 1)
        os, err := f.Read(b1)
        if err != nil {
                return nil, err
        }

        return b1[:os], err
        //osCheck = fmt.Printf("%s",string(b1[40]))
}

func main() {
        os, err := OsCheck()
        if err != nil {
                 return
        }

        fmt.Printf("%c", os)
}
