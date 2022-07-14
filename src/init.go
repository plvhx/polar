// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package src

import (
	"errors"
	"os"
	"syscall"
)

var (
	hookChangeDirErr = errors.New("Failed to change git-hooks directory into .polar/hooks")
	hookCreateDirFailErr = errors.New("Failed to create recursive directory .polar/hooks")
)

func Init() error {
	args := []string{"config", "core.hooksPath", ".polar/hooks"}
	err  := syscall.Exec("git", args, os.Environ())

	if err != nil {
		return hookChangeDirErr
	}

	_, err = os.Stat(".polar")

	if err == nil {
		return nil
	}

	err = os.MkdirAll(".polar/hooks", 0775)

	if err != nil {
		return hookCreateDirFailErr
	}

	return nil
}
