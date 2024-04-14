/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/db"

	"github.com/spf13/cobra"
)

// icdCmd represents the icd command
var IcdCmd = &cobra.Command{
	Use:   "icd",
	Short: "Changing directory, a wrapper for 'cd' command",
	Long:  `Long for icd command`,
	Run: RunIcd,
}

func RunIcd(cmd *cobra.Command, args []string) {
	UserPath := args[0]

	//Start goroutine to fetch from DB
	//Start goroutine named: DB:
	//                      - Update DB if path exists/scoring
	fmt.Println("icd called")

	database, err := db.GetDatabase(constants.DBname)

	database.GetPath(UserPath)
	if err != nil{
		fmt.Printf("An error has occured!")
	}

	database.InsertPath(UserPath)

	// os.Stdout.WriteString("~/")
}


func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// icdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// icdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
