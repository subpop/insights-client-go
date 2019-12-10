package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const inventoryHostURL = "https://cloud.redhat.com/api/inventory/v1/hosts?insights_id="

func show(cfg *config) error {
	var err error

	client, err := newClient(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return err
	}

	id := getMachineID()

	req, err := http.NewRequest(http.MethodGet, inventoryHostURL+id, nil)
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

	var system map[string]interface{}
	err = json.Unmarshal(data, &system)
	if err != nil {
		return err
	}

	fmt.Println(system)

	return nil
}
