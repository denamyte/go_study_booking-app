package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func contextExample() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// create HTTP request
	request, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// get data from HTTP response
	imageData, err := ioutil.ReadAll(res.Body)

	fmt.Printf("downloaded image size %d\n", len(imageData))
}
