package cmd

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"

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
    fmt.Println("WARNING! This will destroy the environment.")
    fmt.Println("There is no undo.")
    fmt.Println("To confirm that you wish to proceed, please type yes: ")

    reader := bufio.NewReader(os.Stdin)
    confirmation, _ := reader.ReadString('\n')
    if strings.EqualFold("yes\n", confirmation) {
      fmt.Println("Really terminating...")
      destroyCommand := exec.Command("terraform", "destroy", "-auto-approve")
      // TODO: the command output should be stored in a log for easier troubleshooting
      destroyError := destroyCommand.Run()
      if destroyError != nil {
        log.Fatal(destroyError)
      }
    }
  },
}
