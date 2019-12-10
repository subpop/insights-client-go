package main

import (
	"os"
	"os/exec"
)

func collect() error {
	cmd := exec.Command("/usr/bin/env",
		"python",
		"-m",
		"insights.collect",
		"--compress")
	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),
		"LANG=" + os.Getenv("LANG"),
		"PYTHONPATH=" + corePath,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
