package main

// import "os"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "bytes"

const bitBucketAPI = "https://api.bitbucket.org/2.0/"

func gitSetupLocalRepo(path string) {
    shell(path, "git", "init")
    gitAdd(path, ".")
    gitCommit(path, "First commit")
}

func gitSetupRemoteRepo(user string, password string, account string, name string, slug string) interface{} {
    config := []byte("{\"scm\": \"git\", \"name\": \"" + name + "\",\"is_private\":true}")

    req, err := http.NewRequest("POST", bitBucketAPI + "repositories/" + account + "/" + slug, bytes.NewBuffer(config))
    req.Header.Add("Content-Type", "application/json")
    req.SetBasicAuth(user, password)

    client := &http.Client{}
    resp, err := client.Do(req)
    checkError(err);

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    checkError(err)

    var responseData interface{}
    err = json.Unmarshal(body, &responseData)
    checkError(err)

    return responseData
}

func gitClone(path string, repo string) {
    shell(path, "git", "clone", "--depth=1", repo, path)
}

func gitAdd(path string, files string) {
    shell(path, "git", "add", files)
}

func gitCommit(path string, message string) {
    shell(path, "git", "commit", "-m", message)
}

func gitAddRemote(path string, remoteName string, remoteAddr string) {
    shell(path, "git", "remote", "add", remoteName, remoteAddr)
}

func gitPush(path string, remote string, branch string) {
    shell(path, "git", "push", remote, branch)
}

func gitBranch(path string, branch string) {
    shell(path, "git", "checkout", "-b", branch)
}
