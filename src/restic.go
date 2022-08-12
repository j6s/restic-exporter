package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type StatsResponse struct {
	TotalSize      int `json:"total_size"`
	TotalFileCount int `json:"total_file_count"`
}

type SnapshotResponse struct {
	Time string `json:"time"`
}

type Restic struct {
	Binary     string
	Name       string
	Repository string
	Password   string
	Env        map[string]string
}

func (restic Restic) Run(arguments []string, target interface{}) error {
	arguments = append(arguments, "--json")

	log.Printf("[%s] %s %s", restic.Name, restic.Binary, arguments)
	command := exec.Command(restic.Binary, arguments...)
	command.Env = append(
		os.Environ(),
		fmt.Sprintf("RESTIC_REPOSITORY=%s", restic.Repository),
		fmt.Sprintf("RESTIC_PASSWORD=%s", restic.Password),
	)

	for key, value := range restic.Env {
		command.Env = append(
			command.Env,
			fmt.Sprintf("%s=%s", key, value),
		)
	}

	output, err := command.Output()
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, target)
	if err != nil {
		return err
	}

	return nil
}

func (restic Restic) SnapshotTimestamp() (int64, error) {
	snapshots := make([]SnapshotResponse, 0)
	err := restic.Run([]string{"snapshots", "latest"}, &snapshots)
	if err != nil {
		return -1, err
	}

	time, err := time.Parse(time.RFC3339Nano, snapshots[0].Time)
	return time.Unix(), err
}
