package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
  Use: "update",
  Short: "Re-run the configuration scripts",
  Long: `Re-run the configuration scripts without provisioning new cloud resources.
  `,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Rad CLI tool v0.1")
    fmt.Println("running the configuration scripts on the environment servers...")
  },
}
