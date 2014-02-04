package docker

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
  "strings"
)


type (
	Docker interface {
		GetEvents() chan *Event
	}

  dockerClient struct {
		url string
	}

	Event struct {
		ContainerId string `json:"id"`
		Status      string `json:"status"`
		Image       string `json:"from"`
    Time        string `json:"time"`
	}
)

func (client *dockerClient) GetEvents() chan *Event {
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
  url := strings.Split(client.url, "://")
	conn, err := net.Dial(url[0], url[1])
	if err != nil {
		return nil, err
	}
	return httputil.NewClientConn(conn, nil), nil

}

func NewClient(url string) (Docker, error) {
	return &dockerClient{url}, nil
}


