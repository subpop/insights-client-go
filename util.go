package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

const machineIDFilePath = "/etc/insights-client/machine-id"

func getMachineID() string {
	if _, err := os.Stat(machineIDFilePath); os.IsNotExist(err) {
		UUID, err := uuid.NewUUID()
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Open(machineIDFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Fprintf(file, "%s", UUID)
		return UUID.String()
	}
	data, err := ioutil.ReadFile(machineIDFilePath)
	if err != nil {
		log.Fatal(err)
	}
	UUID, err := uuid.Parse(string(data))
	if err != nil {
		log.Fatal(err)
	}
	return UUID.String()
}
