package docker

import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"net"
//	"net/http"
//	"net/http/httputil"
  //"docker-event-coordinator/event"
)


type (
	Docker interface {
	//	GetEvents() chan *Event
	}

  dockerClient struct {
		url string
	}
)
