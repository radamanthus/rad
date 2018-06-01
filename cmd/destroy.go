package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
  Use: "destroy",
  Short: "Deprovision the cloud resources for the environment",
  Long: `Deprovision the cloud resources for the environment
  `,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Rad CLI tool v0.1")
    fmt.Println("Destroying the environment...")
  },
}
