package go_sheep

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HTTPTransport struct {
	Error
	Client http.Client
}

func NewHTTPTransport() Transport {
	return &HTTPTransport{
		Client: http.Client{
			Timeout: time.Second * 2,
		},
		Error: Error{
			File:    "http.go",
			Package: "go_sheep",
		},
	}
}

func (h *HTTPTransport) SendMessage(n Node, m Message) error {
	msgByt, err := json.Marshal(m)
	if err != nil {
		return h.NewError("Could not Marshal message for HTTP sending", "SendMessage", err)
	}
	reqBody := bytes.NewReader(msgByt)

	_, err = http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/%s", n.IP, n.Port, "endpoint"), reqBody)
	if err != nil {
		return h.NewError("Could not build HTTP requrest", "SendMessage", err)
	}

	return nil
}
