package main

import "fmt"
import "os"
import "io/ioutil"

import "github.com/spf13/viper"

func copyKey(config *viper.Viper) {

    homeDir := os.Getenv("HOME");
    sshKey := homeDir + "/.ssh/id_rsa"

    //check if key exists
    if !fileExists(sshKey) {
        fmt.Println("Could not find your id_rsa key")
        return
    }

    key, err := os.Open(sshKey)
    checkError(err)

    //check if a key already exists
    if fileExists("./keys/id_rsa") {
        fmt.Println("You already have an id_rsa key in your keys folder")
        return
    }

    if !fileExists("./keys") {
        fmt.Println("No keys directory exists (./keys)")
        return
    }

    //get contents of key
    fmt.Println("Reading your private key")
    keyContents, err := ioutil.ReadAll(key)

    fmt.Println("Writing your key to keys/id_rsa")
    newKey, err := os.Create(cwd() + "/keys/id_rsa")
    checkError(err)

    newKey.Write(keyContents)
    newKey.Sync()
    newKey.Close()
}
