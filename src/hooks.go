package src

import (
	"os"
)

var (
	hookTypeErr = errors.New("Invalid hook type.")
)

func CreateHook(hook, contents string) err {
	if !Validate(hook) {
		return hookTypeErr
	}

	f, err := os.OpenFile(hook, os.O_RDWR | os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	finfo, err := f.Stat()

	if err != nil {
		return err
	}

	if err := f.Truncate(finfo.Size()); err != nil {
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