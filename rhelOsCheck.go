package main

import (
       "os"
       "fmt"
)

//os check for file path
func OsCheck() (OsVersion []byte, err error) {
        f, err := os.Open("/etc/redhat-release")
        if err != nil {
                return nil, err
        }

        if _, err := f.Seek(40, 0); err != nil {
                f.Close()
                return nil, err
        }

        b1 := make([]byte, 1)
        os, err := f.Read(b1)
        if err != nil {
                f.Close()
                return nil, err
        }

        f.Close()
        return b1[:os], err
}

func main() {
        os, err := OsCheck()
        if err != nil {
                 return
        }

        fmt.Printf("%c", os)
}
