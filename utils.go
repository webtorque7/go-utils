package main

import (
    "fmt"
    "os"
    "github.com/spf13/viper"
)

func main() {

    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("Please specify a command")
        return
    }

    config := viper.New()
    //load in config
    config.SetConfigName("config")
    config.AddConfigPath("./")
    config.SetConfigType("yaml")

    err := config.ReadInConfig()
    checkError(err)

    //call correct command
    switch args[0] {

    case "copy-key":
        copyKey(config)

    case "create-site":
        createSite(config)

    case "update-self":
        updateSelf()
    default:
        fmt.Println("Please specify a command")
    }


    // resp, err := http.Get("http://www.webtorque.co.nz")
    // checkError(err);
    //
    // defer resp.Body.Close()
    // body, err := ioutil.ReadAll(resp.Body)
    // checkError(err)
    //
    // file, err := os.Create("wt.html")
    // checkError(err)
    // file.Write(body)
    // file.Sync()
    // file.Close()
    //
    // fmt.Printf(string(body))
}
