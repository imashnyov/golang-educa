package main

import (
	"fmt"
)

func main() {
	var pattern = []string{}
	var tgChannel = []string{}
	var args []string

	for _, url := range tgChannel {
		args = append(args, url)
	}
	result := curl(args, pattern)
	for _, value := range result {
		fmt.Println(value)
	}
}
