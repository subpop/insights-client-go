package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

const uploadURL = "https://cert.cloud.redhat.com/api/ingress/v1/upload"

// upload submits archivePath to the Insights service for analysis.
func upload(cfg *config, archivePath string) error {
	var err error

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return err
	}

	client, err := newClient(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return err
	}

	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodPost, uploadURL, f)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK:
		break
	default:
		return &unexpectedResponseErr{statusCode: res.StatusCode, body: string(data)}
	}

	return nil
}
