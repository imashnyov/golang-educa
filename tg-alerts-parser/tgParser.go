package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func tgParser(resp *http.Response, pattern []string) []string {
	var parsedMessages []string

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: error parsing %s", err)
		os.Exit(1)
	}

	// Find and process the time elements
	doc.Find("time").Each(func(index int, element *goquery.Selection) {
		// Get the "datetime" attribute
		datetime, exists := element.Attr("datetime")
		if !exists {
			fmt.Println("Datetime attribute not found.")
			return
		}

		// Parse the datetime string into a time.Time value
		parsedTime, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			fmt.Printf("Error parsing datetime: %v\n", err)
			return
		}

		// Add 3 hours to the time
		modifiedTime := parsedTime.Add(3 * time.Hour)

		// Format the modified time with date and time
		modifiedTimeString := modifiedTime.Format("02 Jan 15:04")

		// Find and print the text from the div.tgme_widget_message_text.js-message_text
		messageText := doc.Find("div.tgme_widget_message_text.js-message_text").Eq(index).Text()

		// Combine date, time, and message text
		result := fmt.Sprintf("%s: %s\n", modifiedTimeString, messageText)

		// Define the regular expression pattern to match
		for _, p := range pattern {
			// Compile the regular expression
			re := regexp.MustCompile(p)
			// Split the input into lines
			lines := strings.Split(result, "\n")
			// Iterate through the lines and append those that contain the pattern
			for _, line := range lines {
				if re.MatchString(line) {
					parsedMessages = append(parsedMessages, line)
				}
			}
		}
	})

	return parsedMessages
}
