package main

import "os/exec"
import "os"
import "bufio"
import "fmt"
import "strconv"

import "github.com/spf13/viper"

func checkError(e error) {
    if e != nil {
        fmt.Println(e.Error())
        doContinue := prompt("Continue with installation [y/n]? ")
        if doContinue != "y" {
            os.Exit(1)
        }
    }
}

func cwd() string {
    dir, err := os.Getwd()
    checkError(err)
    return string(dir)
}

func fileExists(filePath string) bool {
    //fmt.Println("Checking " + filePath + " exists")
    _, err := os.Stat(filePath)
    return err == nil
}

func prompt(question string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(question)
    text, _ := reader.ReadString('\n')
    //strip trailing newline
    return text[0:len(text) - 1]
}

func shell(path string, args ...string) {
    commandArgs := args[1:]
    command := exec.Command(args[0], commandArgs...)

    if path != "" {
        command.Dir = path
    }

    _, err := command.Output()
    checkError(err)
}

func getConfigStringOrPrompt(key string, config *viper.Viper, message string) string {
    var value string

    if config.IsSet(key) {
        value = config.GetString(key)
    } else {
        value = prompt(message)
    }

    return value;
}

func getConfigIntOrPrompt(key string, config *viper.Viper, message string)  int {
    var value int

    if config.IsSet(key) {
        value = config.GetInt(key)
    } else {
        intVal, err := strconv.ParseInt(prompt(message), 10, 32)
        checkError(err)
        value = int(intVal)
    }

    return value
}

func getConfigBoolOrPrompt(key string, config *viper.Viper, message string) bool {
    var value bool

    if config.IsSet(key) {
        value = config.GetBool(key)
    } else {
        var err error
        value, err = strconv.ParseBool(prompt(message))
        checkError(err)
    }

    return value
}
