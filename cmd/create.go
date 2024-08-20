package cmd

import (
	"fmt"
	"log"

	"github.com/avran02/package-manager/internal/app"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Upload local package to a remote server",

	Run: createPackage,
}

func createPackage(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("No package file")
	}
	for _, pkgFile := range args {
		app.New().CreatePackage(pkgFile)
		fmt.Println("Created package:", pkgFile)
	}
	fmt.Println("Done!")
}

func init() {
	rootCmd.AddCommand(createCmd)
}
