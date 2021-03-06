package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Show findyshark version",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(rootCmd.Use + " " + VERSION)
  },
}

func init() {
  rootCmd.AddCommand(versionCmd)
}