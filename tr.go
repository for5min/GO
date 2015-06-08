package main

import (
	"fmt"
	"os/exec"
)

func main() {
	argv := []string{"-a"}
	c := exec.Command("ls", argv...)
	d, _ := c.Output()
	fmt.Println(string(d))
}
