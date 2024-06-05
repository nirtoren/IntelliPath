/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package setup

import (
	"fmt"
	"intellipath/internal/constants"
		"log"
	"os"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the deinit command
var DeinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Are you sure you want to de-init Intellipath? [y/n]")
		var choice string
    	n, err := fmt.Scanln(&choice)
		if n > 1 {
			log.Fatal("Answer shoul dbe `y` or `n`")
		}
		if err != nil {
			log.Fatal(err)
		}

		if choice == "y" {
			// Delete the database and the executable
			value := os.Getenv(constants.INTELLIPATH_DIR)
			os.RemoveAll(value + constants.INTELLIPATH)			
			
			// Delete the icd() and the export PATH from bashrc
			fmt.Println("Intellipath de-initialized successfully!")

		} else {
			fmt.Println("De-initialization aborted!")
			os.Exit(1)
		}

	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
