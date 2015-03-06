package main

import (
	"bufio"
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
)

func main() {
	temp := os.TempDir()
	fmt.Println(temp)
	//	temp1, err := ioutil.TempDir("/tmp", "test")
	//	if err != nil {
	//		fmt.Println("Error")
	//}
	//	fmt.Println(temp1)
	//	tempfile, err := ioutil.TempFile("/tmp", "test")
	//	if err != nil {
	//		fmt.Println("Error")
	//	}
	//	fmt.Println(tempfile)
	file, err := os.Open("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
