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
	Name       string
	Repository string
	Password   string
}

func (restic Restic) Run(arguments []string, target interface{}) error {
	arguments = append(arguments, "--json")

	log.Printf("[%s] restic %s", restic.Name, arguments)
	command := exec.Command("restic", arguments...)
	command.Env = append(os.Environ(), fmt.Sprintf("RESTIC_REPOSITORY=%s", restic.Repository), fmt.Sprintf("RESTIC_PASSWORD=%s", restic.Password))
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
