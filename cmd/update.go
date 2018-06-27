package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
  "strings"
)

type config struct {
  AppName               string `yaml:"APP_NAME"`
  HostName              string `yaml:"HOSTNAME"`
  RailsEnv              string `yaml:"RAILS_ENV"`
  ScmUrl                string `yaml:"SCM_URL"`
  InstanceSize          string `yaml:"INSTANCE_SIZE"`
  DbVolumeSize          string `yaml:"DB_VOLUME_SIZE"`
  MntVolumeSize         string `yaml:"MNT_VOLUME_SIZE"`
  DataVolumeSize        string `yaml:"DATA_VOLUME_SIZE"`
  WebWorkerMemoryLimit  string `yaml:"WEB_WORKER_MEMORY_LIMIT"`
  WebWorkerCount        string `yaml:"WEB_WORKER_COUNT"`
}

func (c *config) loadConfigFromFile(configFilename string) *config {
  yamlFile, err := ioutil.ReadFile(configFilename)
  if err != nil {
    log.Printf("yamlFile.Get err #%v ", err)
  }
  err = yaml.Unmarshal(yamlFile, c)
  if err != nil {
    log.Fatalf("Unmarshal: %v", err)
  }

  return c
}

func prepareConfigFile(configFilename string, configVars config) {
  read, err := ioutil.ReadFile(configFilename)
  if err != nil {
    panic(err)
  }

  // FIXME: Is there a better way to do this?
  newContents := strings.Replace(string(read), "$APP_NAME", configVars.AppName, -1)
  newContents = strings.Replace(string(newContents), "$RAILS_ENV", configVars.RailsEnv, -1)
  newContents = strings.Replace(string(newContents), "$SCM_URL", configVars.ScmUrl, -1)
  newContents = strings.Replace(string(newContents), "$HOSTNAME", configVars.HostName, -1)
  newContents = strings.Replace(string(newContents), "$INSTANCE_SIZE", configVars.InstanceSize, -1)
  newContents = strings.Replace(string(newContents), "$DB_VOLUME_SIZE", configVars.DbVolumeSize, -1)
  newContents = strings.Replace(string(newContents), "$MNT_VOLUME_SIZE", configVars.MntVolumeSize, -1)
  newContents = strings.Replace(string(newContents), "$DATA_VOLUME_SIZE", configVars.DataVolumeSize, -1)
  newContents = strings.Replace(string(newContents), "$WEB_WORKER_MEMORY_LIMIT", configVars.WebWorkerMemoryLimit, -1)
  newContents = strings.Replace(string(newContents), "$WEB_WORKER_COUNT", configVars.WebWorkerCount, -1)

  err = ioutil.WriteFile(configFilename, []byte(newContents), 0)
  if err != nil {
    panic(err)
  }
}

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

    fmt.Println("Parsing rad.yml...")
    var configVars config
    configVars.loadConfigFromFile("rad.yml")

    configFiles := []string{
      "configure.sh",
      "app.conf",
      "nginx.conf",
      "database.yml",
    }
    for _, configFile := range configFiles {
      prepareConfigFile(configFile, configVars)
    }
  },
}
