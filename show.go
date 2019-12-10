package main

import (
	"encoding/json"
	"fmt"
)

const inventoryHostURL = "https://cert.cloud.redhat.com/api/inventory/v1/hosts?insights_id="
const insightsSystemReportURL = "http://cert.cloud.redhat.com/api/insights/v1/system/%v/reports/"

func show(cfg *config) error {
	var err error

	client, err := newClient(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return err
	}

	machineID := getMachineID()

	data, err := get(client, inventoryHostURL+machineID)
	if err != nil {
		return err
	}

	var result struct {
		Results []map[string]interface{} `json:"results"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	id, ok := result.Results[0]["id"].(string)
	if !ok {
		return &invalidKeyTypeErr{key: "id", val: result.Results[0]["id"]}
	}

	data, err = get(client, fmt.Sprintf(insightsSystemReportURL, id))
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
