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

	fmt.Println("Programa para ver si estan activas las url")
	fmt.Println("Type a file name:")
	fileName, err := reader.ReadString('\n')

	// Limpiar la url
	fileName = strings.TrimSpace(fileName)

	// Saber si ha dado erro
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Split the file to a list of urls
	urlList = strings.Split(proccessFile(fileName), "\n")

	// Call function to print the result
	impressResult(urlList)

}

// Function to proccess the file
func proccessFile(fileName string) string {
	data, err := os.ReadFile(fileName)

	// If there is an error return empty
	if err != nil {
		return ""
	}

	// Return data convert to string
	return string(data)
}

// Function to validate the url
func validateUrl(url string) bool {

	// Create a http cliente with five seconds of life
	var httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	// Hacer la peticion get a la url
	response, err := httpClient.Get(url)

	// Validar si ha dado error
	if err != nil {
		return false
	}

	// Cerrar la peticion
	defer response.Body.Close()

	// Validar si la respuesta es 200
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}

}

// Function to print the result
func impressResult(urlList []string) {

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

			// Validate the url
			exist := validateUrl(url)

			// Print the resault
			if exist {
				fmt.Println("El sitio " + url + " esta activo")
			} else {
				fmt.Println("El sitio " + url + " no existe o no esta activo")
			}
		}(url)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
