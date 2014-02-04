package main

import (
	"fmt"
  "github.com/cpuguy83/docker-event-coordinator/docker"
)


func (event *Event) handle() {
	fmt.Println(event.Time)
}



func main() {
	client, nil := NewClient("tcp://192.168.42.43:4243")
	events := client.GetEvents()
	for event := range events {
		go event.handle()
	}
}


