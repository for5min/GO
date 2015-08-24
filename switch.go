package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("I'm Mac")
	case "linux":
		fmt.Println("I'm Linux")
	}
}