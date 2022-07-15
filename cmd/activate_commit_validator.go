// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package cmd

import (
	"github.com/plvhx/polar/src"
	"github.com/spf13/cobra"
)

var (
	hookName    = "prepare-commit-msg"
	hookCommand = "polar validate-commit"
)

var activateFunc = func(cmd *cobra.Command, args []string) {
	if err := src.CreateHook(hookName, hookCommand); err != nil {
		panic(err)
	}
}

var activateCmd = &cobra.Command{
	Use:     "activate-commit-validator",
	Short:   "Activate prepare-commit-msg hook",
	Long:    "Activate prepare-commit-msg hook for validating commit message.",
	Example: "polar activate-commit-validator",
	Run:     activateFunc,
}

func init() {
	rootCmd.AddCommand(activateCmd)
}
