package main
import (
	"fmt"
	"os"
)

var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	fmt.Println(HOME, USER, GOROOT)
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is : %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
