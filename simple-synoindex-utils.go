package main

import (
    "os"
    "net/url"
    "path/filepath"
)

func GetCurrentExecDir() (string) {

    execdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    return execdir

}


// Get os arguments without first one
func GetArguments() ([]string) {

    v :=[]string{}
    for i := 1; i < len(os.Args); i++ {
        v = append(v, os.Args[i])
    }

    return v
}


func EncodeArguments(args []string) (string) {
    
    v := url.Values{}

    for _, ar := range args {
        v.Add("args", ar)
    }
    qs := v.Encode()

    return qs

}

