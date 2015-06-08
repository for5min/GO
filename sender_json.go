package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//{
//  "broker": "",
//  "exchange": {
//    "name": "scaleworks",
//    "type": "topic",
//    "durable": "true"
//  },
//  "jobs": {
//    "sync_nodes": {
//      "interval": "3m",
//      "routing_key": "sync_nodes",
//      "command_template": {
//	"job_name": "sync_nodes",
//	"when_distributed": "Time.now"
//      }
//    },
//    "sync_clouds": {
//      "interval": "1d",
//      "routing_key": "sync_nodes"
//    }
//  }
//}

type Config struct {
	Broker   string
	Exchange ExchangeType
	Jobs     JobsType
}

type ExchangeType struct {
	Name    string
	Type    string
	Durable string
}

type JobsType struct {
	Sync_nodes  Sync_nodesType
	Sync_clouds Sync_cloudsType
}

type Sync_nodesType struct {
	Interval         string
	Routing_key      string
	Command_template Command_templateType
}

type Command_templateType struct {
	Job_name         string
	When_distributed string
}

type Sync_cloudsType struct {
	Interval    string
	Routing_key string
}

func main() {

	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Can't find the file")
		os.Exit(1)
	}

	jsonFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := Config{}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	name := config.Exchange.Name()
	fmt.Printf("Value: %v\n", config)
	fmt.Printf(name)

}
