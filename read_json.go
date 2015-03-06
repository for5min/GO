package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Record struct {
	Uid      string `json:"Uid"`
	Gid      string `json:"Gid"`
	Username string `json:"Username"`
	Name     string `json:"Name"`
	HomeDir  string `json:"HomeDir"`
}

//func main(r io.Reader) (x *Record, err error) {
//	x = new(Record)
//	err = json.NewDecorder(r).Decode(x)
//	return
func main() {
	var data Record
	file, err := ioutil.ReadFile("text.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Uid)
}
