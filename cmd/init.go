package cmd

import (
  "fmt"
  "log"
  "os/exec"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
  Use: "init",
  Short: "Provision cloud resources for the first time",
  Long: `Provision cloud resources for the first time.

    If there are already existing resources for this environment, then nothing is done.

    If you have already provisioned the environment and would like to re-run the configuration script, run "rad update" instead.
  `,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Rad CLI tool v0.1")
    fmt.Println("Initializing the environment...")
    initCommand := exec.Command("terraform", "init", "-input=false")
    // TODO: the command output should be stored in a log for easier troubleshooting
    initError := initCommand.Run()
    if initError != nil {
      log.Fatal(initError)
    }
  },
}
