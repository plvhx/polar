// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package src

import (
	"errors"
	"os"
)

var (
	hookTypeErr = errors.New("Invalid hook type.")
)

func CreateHook(hook, contents string) error {
	if !Validate(hook) {
		return hookTypeErr
	}

	err := os.Chdir(".polar/hooks")

	if err != nil {
		return err
	}

	f, err := os.OpenFile(hook, os.O_RDWR | os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := f.WriteAt([]byte(contents), 0); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}