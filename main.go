package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
  "github.com/cpuguy83/docker-event-coordinator/docker"
  "github.com/cpuguy83/docker-event-coordinator"
)

type (


)

func (event *Event) handle() {
	fmt.Println(event.Time)
}

func (client *dockerClient) GetEvents() <-chan *Event {
	eventChan := make(chan *Event, 100)
	go func() {
		defer close(eventChan)

		conn, _ := client.newConnection()
		defer conn.Close()
		req, _ := http.NewRequest("GET", "/events", nil)

		resp, _ := conn.Do(req)
		defer resp.Body.Close()

		dec := json.NewDecoder(resp.Body)
		for {
			var event *Event
			if err := dec.Decode(&event); err != nil {
				if err == io.EOF {
					break
				}
				continue
			}
			eventChan <- event
		}
	}()
	return eventChan
}

func (client *dockerClient) newConnection() (*httputil.ClientConn, error) {
	conn, err := net.Dial("tcp", "192.168.42.43:4243")
	if err != nil {
		return nil, err
	}
	return httputil.NewClientConn(conn, nil), nil

}

func NewClient(url string) (Docker, error) {
	return &dockerClient{url}, nil
}

func main() {
	client, nil := NewClient("tcp://192.168.42.43:4243")
	events := client.GetEvents()
	for event := range events {
		go event.handle()
	}
}
