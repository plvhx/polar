// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "polar",
	Short: "Git hooks made easy.",
	Long: "polar is a tool to help you manage your git hooks."
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		return
	}
}

func init() {
	rootCmd.Flags()
}
