package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "journal",
  Short: "do some shit",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    // don't need to do anything yet
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
