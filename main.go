package main

import (
	"log"
	"net/http"
)

// this define a common
type restClientInterface interface {
	Get() (*http.Response, error)
}

type restClient struct {
	URL string
}

func (client *restClient) Get() (response *http.Response, err error) {
	response, err = http.Get(client.URL)
	return
}

func fetch(thisClient restClientInterface) (response *http.Response, err error) {
	response, err = thisClient.Get()
	return
}

func main() {
	println("Starting the application!")

	thisClient := &restClient{}
	thisClient.URL = "https://jsonplaceholder.typicode.com/todos/404"
	response, err := fetch(thisClient)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response)
	println("Application ended")
}
