package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("12 * * * *", func() { fmt.Println("hello") })
	c.Start()
}
