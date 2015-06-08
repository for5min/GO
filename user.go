package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	u, err := user.Lookup(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
