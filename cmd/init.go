// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package cmd

import (
	"github.com/plvhx/polar/src"
	"github.com/spf13/cobra"
)

var initFunc = func(cmd *cobra.Command, args []string) {
	if err := src.Init(); err != nil {
		panic(err)
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize polar-bear",
	Long:  "polar-bear is a tool to help you manage your git hooks.",
	Run:   initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
