package git

import (
    "fmt"
    "github.com/urfave/cli/v2"
    "bytes"
    "os/exec"
)

func CleanBranchLocal(c *cli.Context) error {
    var stdout, stderr bytes.Buffer

    gitFetch := exec.Command("git", "fetch", "-p")
    gitFetch.Stdout = &stdout
    gitFetch.Stderr = &stderr

    err := gitFetch.Run()
    if err != nil {
        fmt.Println("erro")
        fmt.Println(err)
    }

    gitBranch := exec.Command("git", "branch", "-vv")
    gitBranch.Stdin = &stdout
    gitBranch.Stdout = &stdout
    gitBranch.Stderr = &stderr

    err = gitBranch.Run()
    if err != nil {
        fmt.Println("erro")
        fmt.Println(err)
    }

    grep := exec.Command("grep", "origin/.*: gone]")
    grep.Stdin = &stdout
    grep.Stdout = &stdout
    grep.Stderr = &stderr

    err = grep.Run()
    if err != nil {
        fmt.Println(err)
    }

    awk := exec.Command("awk", "{print $1}")
    awk.Stdin = &stdout
    awk.Stdout = &stdout
    awk.Stderr = &stderr

    err = awk.Run()
    if err != nil {
        fmt.Println(err)
    }

    gitBranchDelete := exec.Command("xargs", "git", "branch", "-D")
    gitBranchDelete.Stdin = &stdout
    gitBranchDelete.Stdout = &stdout
    gitBranchDelete.Stderr = &stderr

    err = gitBranchDelete.Run()
    if err != nil {
        fmt.Println(err)
    }

    return nil
}