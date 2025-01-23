package main

import (
	"fmt"
	"task2/implementation/client"
)

func main() {
	MyClient := client.MakeNewClient("http", "localhost", "8080")
	var err = true
	for i := 0; err; i++ {
		version := MyClient.RunGetAPIVersionRequest()
		if version == "" {
			fmt.Println("failed to get API version")
			err = false
		} else {
			fmt.Println(version)
		}
		str := MyClient.DecodeString(`{"inputString": "SGF2ZSBhIG5pY2UgZGF5"}`)
		if str == "" {
			fmt.Println("failed to post")
			err = false
		} else {
			fmt.Println(str)
		}
		status, ok := MyClient.RunHardOpRequest()
		if !ok {
			fmt.Println(ok, "worked too long")
		} else {
			fmt.Println(ok, status)
		}
	}
}
