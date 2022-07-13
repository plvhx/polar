// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package cmd

import (
	"errors"
	"os"

	"github.com/plvhx/polar-bear/src"
	"github.com/spf13/cobra"
)

var (
	hookInitErr = errors.New("Run 'polar init' first.")
)

var addFunc = func(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(".polar"); err != nil {
		panic(hookInitErr)
	}

	if err := src.CreateHook(args[0], args[1]); err != nil {
		panic(err)
	}
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a new hook",
	Long: "Adds a new hook to polar and installs it.",
	Args: cobra.ExactArgs(2),
	Example: `polar add pre-commit "
echo 'this is will run at pre-commit state.'
"`,
	Run: addFunc,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
