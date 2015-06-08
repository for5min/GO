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

func ReadConfig() {

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
