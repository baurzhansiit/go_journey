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
		if err != nil {
			fmt.Println("There no git repository available")
			panic(err.Error())
		}
		fmt.Println(string(out[:len(out)-1]))
		_, err = exec.Command("git", "add", "./").Output()
		if err != nil {
			fmt.Println("err when add")
			panic(err)
		}
		_, err = exec.Command("git", "commit", "-am", os.Args[1]).Output()
		if err != nil {
			fmt.Println("err when commit")
			panic(err)
		}
		branchName := string(out[:len(out)-1])
		_, err = exec.Command("git", "push", "origin", branchName).Output()
		if err != nil {
			fmt.Println("err when push")
			panic(err)
		}
		fmt.Printf("2. Pushed commit to branch name: %s\n", out)
	}

}
