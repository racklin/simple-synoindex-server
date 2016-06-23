package main

import (
    "os"
    "net/http"
    "fmt"
    "github.com/go-ini/ini"
)



func main() {

    if (len(os.Args) == 1 ) {
        fmt.Println("Usage:\nsynoindex [OPTIONS]")
        os.Exit(1)
    }

    // get current execute file path
    execdir := GetCurrentExecDir()

    qs := EncodeArguments(GetArguments())

    inifile := fmt.Sprintf("%s/simple-synoindex-server.ini", execdir)

    cfg , _ := ini.LooseLoad(inifile)

    srvIp := cfg.Section("main").Key("SERVER_IP").MustString("172.17.0.1")
    srvPort := cfg.Section("main").Key("SERVER_PORT").MustString("32699")

    reqUrl := fmt.Sprintf("http://%s:%s/synoindex?%s", srvIp, srvPort, qs) 

    req, err := http.Get(reqUrl)
    if err != nil {
        panic(err)
    }

    req.Body.Close()
    os.Exit(0)
}

