package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ INFO ] Add string argument eg. <g 'comment'> before Pushing to Github")

	} else {
		fmt.Printf("Github pushing with commants %s\n", os.Args[1])
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}
}
