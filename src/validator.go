// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package src

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
