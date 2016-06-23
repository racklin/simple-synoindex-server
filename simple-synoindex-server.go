package main

import (
    "io"
    "net/http"
    "fmt"
    "strings"
    "github.com/go-ini/ini"
    "log"
    "os/exec"
)

var (
    cfg *ini.File
    volumeMappings map[string]string
)


func init() {

    // get current execute file path
    execdir := GetCurrentExecDir()

    inifile := fmt.Sprintf("%s/simple-synoindex-server.ini", execdir)
    cfg , _ = ini.LooseLoad(inifile)

    volumeMappings = cfg.Section("mappings").KeysHash()

}

func remappingPath(srcPath string) string {

    newPath := srcPath

    for vPath, mPath := range volumeMappings {
        newPath = strings.Replace(newPath, vPath, mPath, 1)
        if newPath != srcPath { return newPath; }
    }

    return newPath
}

func SynoIndex(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "ok\n")
    req.ParseForm()
    args := req.Form["args"]

    args[1] = remappingPath(args[1])

    // execute /usr/syno/bin/synoindex 
    cmd := exec.Command("/usr/syno/bin/synoindex", args...)

    err :=  cmd.Run()

    if err != nil {
        log.Printf("SynoIndex Error: %s")
    }

}

func main() {

    srvIp := cfg.Section("main").Key("SERVER_IP").MustString("172.17.0.1")
    srvPort := cfg.Section("main").Key("SERVER_PORT").MustString("32699")
    srvListen := fmt.Sprintf("%s:%s", srvIp, srvPort)

    http.HandleFunc("/synoindex", SynoIndex)
    log.Fatal(http.ListenAndServe(srvListen, nil))

}

