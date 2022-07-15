// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package src

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	allowedHooks = []string{
		"applypatch-msg",
		"commit-msg",
		"fsmonitor-watchman",
		"post-update",
		"pre-applypatch",
		"pre-commit",
		"pre-merge-commit",
		"prepare-commit-msg",
		"pre-push",
		"pre-rebase",
		"pre-receive",
		"update",
	}
)

var (
	fileModeMask         = fs.FileMode(0x1ff)
	gitCommitMsgFile     = ".git/COMMIT_EDITMSG"
	commitTypeConfigFile = ".polar-commitmsg-types.yaml"

	commitMsgNotMatchErr = errors.New("Commit message not match.")
)

func Validate(name string) bool {
	match := false

	for _, hook := range allowedHooks {
		if hook == name {
			match = true
			break
		}
	}

	return match
}

func FetchSerializedCommitTypes() (string, error) {
	finfo, err := os.Stat(commitTypeConfigFile)

	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(commitTypeConfigFile, os.O_RDONLY, finfo.Mode()&fileModeMask)

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)

	io.Copy(buf, f)

	if err := f.Close(); err != nil {
		return "", err
	}

	types := make(map[string][]string, 0)

	err = yaml.Unmarshal(buf.Bytes(), &types)

	if err != nil {
		return "", err
	}

	return strings.Join(types["types"], "|"), nil
}

func buildMatchingPattern(types string) (string, error) {
	return fmt.Sprintf("^(?:%s)(?:\\(.*\\))(?:\\!)?(?:\\: )(.*)", types), nil
}

func FetchCommitMessage() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	finfo, err := os.Stat(gitCommitMsgFile)

	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(gitCommitMsgFile, os.O_RDONLY, finfo.Mode()&fileModeMask)

	if err != nil {
		return nil, err
	}

	io.Copy(buf, f)

	if err := f.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func ValidateCommit(buf []byte) error {
	types, err := FetchSerializedCommitTypes()

	if err != nil {
		return err
	}

	pat, err := buildMatchingPattern(types)

	if err != nil {
		return err
	}

	matched, err := regexp.Match(pat, buf)

	if err != nil {
		return err
	}

	if !matched {
		return commitMsgNotMatchErr
	}

	return nil
}
