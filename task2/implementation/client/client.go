package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const timeout = 15 * time.Second

type Response struct {
	Str string `json:"outputString"`
}

func RunGetAPIVersionRequest(url string) string {
	resp, err := http.Get(url)
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

func RunHardOpRequest(url string) (int, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	return resp.StatusCode, true
}

func RunPostRequest(url, str string) string {
	reqBody := bytes.NewBuffer([]byte(str))
	resp, err := http.Post(url, "application/json", reqBody)
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
func DoRequests(urls []string) {
	var err = true
	for i := 0; err; i++ {
		index := i % len(urls)
		switch index {
		case 0:
			version := RunGetAPIVersionRequest(urls[index])
			if version == "" {
				fmt.Println("failed to get API version")
				err = false
			} else {
				fmt.Println(version)
			}
		case 1:
			str := RunPostRequest(urls[index], `{"inputString": "SGF2ZSBhIG5pY2UgZGF5"}`)
			if str == "" {
				fmt.Println("failed to post")
				err = false
			} else {
				fmt.Println(str)
			}
		case 2:
			status, ok := RunHardOpRequest(urls[index])
			if !ok {
				fmt.Println(ok, "worked too long")
			} else {
				fmt.Println(ok, status)
			}
		}
	}
}
