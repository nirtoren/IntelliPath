/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package setup

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/db"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		database, err := db.GetDatabase(constants.DBname)

		if err != nil {
			fmt.Printf("error was occured during icd init.")
			os.Exit(1)
		}

		// Testing insertion on initialization of the database
		var userID int64
		rec, _ := db.NewRecord("~/", 0)
		userID, err = database.InsertRecord(rec)
		if err != nil {
			fmt.Printf("Error on insertion to database: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("insertion completed. %d\n", userID)

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
