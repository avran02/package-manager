package cmd

import (
	"fmt"
	"log"

	"github.com/avran02/package-manager/internal/app"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Download package from remote server to a local machine",

	Run: updatePackage,
}

func updatePackage(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("No package file")
	}
	for _, pkgFile := range args {
		app.New().UpdatePackage(pkgFile)
		fmt.Println("Updated package:", pkgFile)
	}
	fmt.Println("Done!")
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
