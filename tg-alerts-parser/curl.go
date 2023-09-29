package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func curl(args []string, pattern []string) []string {
	var result []string
	// var oneMinuteResult []string

	// Infinite loop
	// for {
	for _, url := range args {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "fetch: %s returned status %v\n", url, resp.Status)
			os.Exit(1)
		}

		midle_result := tgParser(resp, pattern)
		result = append(result, midle_result...)
		// fmt.Println(result)
		// time.Sleep(2 * time.Second)
	}
	return result
}
