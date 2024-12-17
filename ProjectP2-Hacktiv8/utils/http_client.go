package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestGET(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, values := range headers {
		req.Header.Set(key, values)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: %s", body)
		return nil, fmt.Errorf("Failed to get data, status code: %d", res.StatusCode)
	}

	return body, nil
}