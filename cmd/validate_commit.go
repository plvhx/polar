// Copyright 2022 Paulus Gandung Prakosa <gandung@lists.infradead.org>
// All rights reserved.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"polar/src"
	"github.com/spf13/cobra"
)

var (
	commitEditMsgErr = errors.New(".git/COMMIT_EDITMSG not found.")
)

var validateFunc = func(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(".git/COMMIT_EDITMSG"); err != nil {
		panic(commitEditMsgErr)
	}

	buf, err := src.FetchCommitMessage()

	if err != nil {
		panic(err)
	}

	types, err := src.FetchSerializedCommitTypes()

	if err != nil {
		panic(err)
	}

	if err := src.ValidateCommit(buf); err != nil {
		fatal(buf, types)
	}
}

var validateCmd = &cobra.Command{
	Use:     "validate-commit",
	Short:   "Validate commit message",
	Long:    "Validate commit message",
	Args:    cobra.ExactArgs(0),
	Example: `polar-bear validate-commit`,
	Run:     validateFunc,
}

func fatal(buf []byte, types string) {
	fmt.Printf("\033[31m----------------------- Invalid commit message -----------------------\033[0m\n")
	fmt.Printf("commit message: \033[31m%s\033[0m\n", strings.TrimRight(string(buf), "\n"))
	fmt.Printf("correct format: \033[33m<type>(<scope>): <subject>\033[0m\n")
	fmt.Printf("available types: \033[35m%s\033[0m\n", strings.Replace(types, "|", ", ", -1))
	fmt.Printf("\033[31m----------------------------------------------------------------------\033[0m\n")
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
