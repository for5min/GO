package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello world")
	cmd := exec.Command("ls")
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%s", d)
}
