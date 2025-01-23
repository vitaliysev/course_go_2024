package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	address string
}

const timeout = 15 * time.Second

type Response struct {
	Str string `json:"outputString"`
}

func MakeNewClient(protocol string, host string, port string) Client {
	return Client{protocol + "://" + host + ":" + port}
}
func (client Client) RunGetAPIVersionRequest() string {
	resp, err := http.Get(client.address + "/version")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (client Client) RunHardOpRequest() (int, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", client.address+"/hard-op", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	return resp.StatusCode, true
}

func (client Client) DecodeString(str string) string {
	reqBody := bytes.NewBuffer([]byte(str))
	resp, err := http.Post(client.address+"/decode", "application/json", reqBody)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var res Response
	err2 := json.Unmarshal(body, &res)
	if err2 != nil {
		return ""
	}
	return res.Str
}
