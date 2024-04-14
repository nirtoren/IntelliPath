/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package setup

import (
	"fmt"
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

		dbFile := "mydb.db"

		database, err := db.NewDatabase(dbFile)

		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			os.Exit(1)
		}

		defer database.Close()


		err = database.Initizlize()
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Database initialization complete.")

		var userID int64
		userID, err = database.InsertPath("/home/desktop")
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
