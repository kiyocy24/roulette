/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a item file",
	Long: `Edit registered items and weights on file.
Vim must be installed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("vim", filepath)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		return c.Run()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
