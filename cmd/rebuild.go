package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(rebuildCmd)
}

var rebuildCmd = &cobra.Command{
  Use: "rebuild",
  Short: "Reprovision the cloud resources",
  Long: `Reprovision the cloud resources.

    If the cloud resources are not yet running, then they're provisioned. If they're already running then some or all of them may be destroyed and replaced by new ones.

    This is the same as running rad init --force
  `,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Rad CLI tool v0.1")
    fmt.Println("rebuilding the environment...")
  },
}
