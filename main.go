package main

import (
	"errors"
	"os"
	"log"
	"syscall"
)

func init() {
	args := []string{"config", "core.hooksPath", ".polar/hooks"}
	err  := syscall.Exec("git", args, os.Environ())

	if err != nil {
		log.Fatal("Failed to change git-hooks directory into .polar/hooks")
	}

	_, err := os.Stat(".polar")

	if err == nil {
		return
	}

	err = os.MkdirAll(".polar/hooks")

	if err != nil {
		log.Fatal("Failed to create recursive directory .polar/hooks")
	}
}

func main() {
	// TODO: do some fancy shit here.. :)
}
