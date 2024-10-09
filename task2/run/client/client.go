package main

import (
	"task2/implementation/client"
)

func main() {
	urls := []string{"http://localhost:8080/version",
		"http://localhost:8080/decode",
		"http://localhost:8080/hard-op",
	}
	client.DoRequests(urls)
}
