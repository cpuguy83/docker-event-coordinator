package main

import (
  "github.com/cpuguy83/docker-event-coordinator/docker"
)

func main() {
	client, _ := docker.NewClient(*url)
	events := client.GetEvents()
	for event := range events {
		go handleEvent(event)
	}
}

func handleEvent(event *docker.Event) {
  // find registered event handlers and run them
}


