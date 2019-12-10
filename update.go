package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const cacheDir = "/var/cache/insights/core/"
const corePath = cacheDir + "insights-core.egg"
const coreSigPath = corePath + ".asc"
const cacheControlFieldsPath = cacheDir + "cache_control_fields.json"
const coreURL = "https://cert-api.access.redhat.com/r/insights/v1/static/core/insights-core.egg"
const coreSigURL = coreURL + ".asc"

// update downloads a new core archive if available.
func update(cfg *config) error {
	var err error

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}

	client, err := newClient(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, coreURL, nil)
	if err != nil {
		return err
	}

	if _, err := os.Stat(cacheControlFieldsPath); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(cacheControlFieldsPath)
		if err != nil {
			return err
		}

		var cacheControlFields map[string]string
		err = json.Unmarshal(data, &cacheControlFields)
		if err != nil {
			return err
		}

		if etag, ok := cacheControlFields["ETag"]; ok {
			req.Header.Set("If-None-Match", etag)
		}

		if modified, ok := cacheControlFields["Last-Modified"]; ok {
			req.Header.Set("If-Modified-Since", modified)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusNotModified:
		return nil
	case http.StatusOK:
		break
	default:
		return &unexpectedResponseErr{statusCode: res.StatusCode, body: ""}
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	f, err := os.Create(corePath)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	f.Close()

	if err := downloadURL(client, coreSigURL, coreSigPath); err != nil {
		return err
	}

	{
		cacheControlFields := make(map[string]string)
		etag := res.Header.Get("ETag")
		if etag != "" {
			cacheControlFields["ETag"] = etag
		}
		modified := res.Header.Get("Last-Modified")
		if modified != "" {
			cacheControlFields["Last-Modified"] = modified
		}
		data, err = json.Marshal(&cacheControlFields)
		if err != nil {
			return err
		}

		f, err := os.Create(cacheControlFieldsPath)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err != nil {
			return err
		}
		f.Close()
	}

	valid, err := verify(corePath, coreSigPath)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("core verification failed")
	}

	return nil
}
