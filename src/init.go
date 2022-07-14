// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package src

import (
	"errors"
	"os"
	"os/exec"
)

var (
	gitBinaryNotFoundErr = errors.New("'git' not found.")
	hookChangeDirErr = errors.New("Failed to change git-hooks directory into .polar/hooks")
	hookCreateDirFailErr = errors.New("Failed to create recursive directory .polar/hooks")
)

func Init() error {
	path, err := exec.LookPath("git")

	if err != nil {
		return gitBinaryNotFoundErr
	}

	cmd := exec.Command(path, "config", "core.hooksPath", ".polar/hooks")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return err
	}

	if err := os.MkdirAll(".polar/hooks", 0775); err != nil && !os.IsExist(err) {
		return hookCreateDirFailErr
	}

	return nil
}
