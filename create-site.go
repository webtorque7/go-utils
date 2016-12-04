package main

import "os"
// import "os/exec"
import "fmt"
import "strings"
import "sync"
import "github.com/spf13/viper"

//settings
const bitBucketAccount = "webtorque-dev"
const installerRepo = "git@github.com:webtorque7/silverstripe-installer"
var branches = []string{"test", "uat", "develop"}

func createSite(config *viper.Viper) {

    installDir := getConfigStringOrPrompt("createSite.installDir", config, "Where would you like to install the site")

    siteName := prompt("Please enter a name for the site: ")
    slug := strings.Replace(strings.ToLower(siteName), " ", "-", -1)
    installPath := installDir + "/" + slug

    //check the install directory exists
    if !fileExists(installDir) {
        fmt.Println("Directory " + installDir + " doesn't exist")
        return
    }

    //check we aren't overwriting an existing directory
    if fileExists(installPath) {
        fmt.Println("Path " + installPath + " already exists")
        return
    }

    //create project folder
    fmt.Println("Making project folder " + installPath)
    os.Mkdir(installPath, 0755)

    //change cwd to install path to make future commands easier
    os.Chdir(installPath)

    fmt.Println("Setting up site in " + installPath)
    gitClone(installPath, installerRepo)

    //remove git files, setup gitInore
    prepareGit()

    //initialise repo, setup remote, do first push
    setupGit(installPath, siteName, slug, config)

    var wg sync.WaitGroup
    wg.Add(2)

    //setup branches
    go func() {
        setupGitBranches(installPath, config.GetStringSlice("createSite.branches"))
        wg.Done()
    }()
    //start vagrant
    go func() {
        fireupVagrant(installPath)
        wg.Done()
    }()

    wg.Wait()

    return
}

func prepareGit() {
    os.RemoveAll(".git")
    os.Remove(".gitignore")
    os.Rename(".gitignore.install", ".gitignore")
}

func setupGitBranches(path string, branches []string) {
    for _, branch := range branches {
        fmt.Println("Setting up branch " + branch)
        gitBranch(path, branch)
        gitPush(path, "origin", branch)
    }
    fmt.Println("Finished setting up branches")
}

func setupGit(path string, siteName string, slug string, config *viper.Viper) {

    fmt.Println("Initialising git")
    gitSetupLocalRepo(path)

    //get git user details
    bitBucketUser := getConfigStringOrPrompt("bitBucket.user", config, "Please enter your bitbucket username: ")
    bitBucketPassword := prompt("Please enter your bitbucket password: ")
    bitBucketAccount := getConfigStringOrPrompt("bitBucket.account", config, "Please enter the bitbucket account to crete the new repository: ")

    gitSetupRemoteRepo(bitBucketUser, bitBucketPassword, bitBucketAccount, siteName, slug)

    //add origin and push
    fmt.Println("Adding git remote and pushing first commit")
    gitAddRemote(path, "origin", "git@bitbucket.org:" + bitBucketAccount + "/" + slug)
    gitPush(path, "origin", "master")
}

func fireupVagrant(path string) {
    fmt.Println("Starting vagrant")
    shell(path, "vagrant", "up")
    fmt.Println("Vagrant started")
}
