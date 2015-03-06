package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//broker:
//exchange:
//  name: scaleworks
//  type: topic
//  durable: true
//jobs:
//  sync_nodes:
//    interval: 3m
//    routing_key: sync_nodes
//    command_template:
//  sync_clouds:
//    interval: 1d
//    routing_key: sync_nodes

type Config struct {
	Broker   string
	Exchange struct {
		Name    string
		Type    string
		Durable bool
	}
	Jobs struct {
		Sync_nodes struct {
			Interval         string
			Routing_key      string
			Command_template string
		}
		Sync_clouds struct {
			Interval    string
			Routing_key string
		}
	}
}

func main() {

	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Can't find the file")
		os.Exit(1)
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := Config{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Value: %v\n", config)

}
