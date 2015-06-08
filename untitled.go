package main

import (
  "fmt"
  "os/exec"

func main() {
	fmt.Println("Hello world")
	cmd := exec.Command("ls")
	d := cmd.Output()
	fmt.Println("%s",d)
	
}
