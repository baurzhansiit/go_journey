package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ErrHandler(err error) {
	if err != nil {
		log.Fatal("some err\n", err)
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ INFO ] Add string argument eg. <g 'comment'> before Pushing to Github")

	} else {
		fmt.Printf("1. Github pushing with commits:  %s\n", os.Args[1])
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		ErrHandler(err)
		add, err := exec.Command("git", "add", ".").Output()
		ErrHandler(err)
		// branchName := string(out[:len(out)-1])
		push, err := exec.Command("git", "push", "origin", "main").Output()
		ErrHandler(err)
		fmt.Printf("2. Pushed commit to branch name: %s,%s,%s\n", out, add, push)
	}
	// listarg := []string{"some", "text"}
	// fmt.Println(strings.Join(listarg, "' '"))

}

// git add .
// #    git commit -am "$msg"
// #    git push origin $currentbranch
