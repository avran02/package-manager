package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dotEnvExample = `SSH_HOST=localhost
SSH_PORT=2222
SSH_USER=packager
SSH_PASSWORD=123
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file on user home directory",

	Run: initConfig,
}

func initConfig(cmd *cobra.Command, args []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFileDir := home + "/.config/package-manager/"

	if err := os.MkdirAll(filepath.Dir(configFileDir), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(configFileDir+".env", []byte(dotEnvExample), 0644); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
