package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	// Variables declared
	var fileName string
	reader := bufio.NewReader(os.Stdin)
	var urlList []string

	fmt.Println("Program to check if URLs are active")
	fmt.Println("Type a file name:")
	fileName, err := reader.ReadString('\n')

	// Clean the URL
	fileName = strings.TrimSpace(fileName)

	// Check if an error occurred
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Split the file to a list of URLs
	urlList = strings.Split(processFile(fileName), "\n")

	// Call function to print the result
	printResults(urlList)

}

// Function to process the file
func processFile(fileName string) string {
	data, err := os.ReadFile(fileName)

	// If there is an error return empty
	if err != nil {
		return ""
	}

	// Return data converted to string
	return string(data)
}

// Function to validate the URL
func validateURL(url string) bool {

	// Create an HTTP client with five seconds of timeout
	var httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	// Make the GET request to the URL
	response, err := httpClient.Get(url)

	// Check if an error occurred
	if err != nil {
		return false
	}

	// Close the response body
	defer response.Body.Close()

	// Check if the response status is 200 OK
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

// Function to print the result
func printResults(urlList []string) {

	// Create a wait group
	var wg sync.WaitGroup

	for _, url := range urlList {

		// Pass empty lines
		if url == "" {
			continue
		}

		wg.Add(1)

		// Run in a goroutine
		go func(url string) {

			// Substract one for the group when the goroutine finishes
			defer wg.Done()

			// Validate the URL
			exist := validateURL(url)

			// Print the result
			if exist {
				fmt.Println("The site " + url + " is active")
			} else {
				fmt.Println("The site " + url + " does not exist or is not active")
			}
		}(url)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
