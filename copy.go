package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, _ := os.Open("test.txt")
	w, _ := os.Create("write.txt")
	num, err := io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
