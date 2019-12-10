package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

// downloadURL downloads the file at URL and saves it to filePath, truncating
// filePath if necessary.
func downloadURL(client *http.Client, URL, filePath string) error {
	var err error

	res, err := client.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return &unexpectedResponseErr{statusCode: res.StatusCode, body: string(data)}
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
