/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/db"
	"intellipath/internal/flows"
	"os"
	"path/filepath"

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
	var outPath string
	UserPath := args[0]
	
	// Stage 1: get the db
	database, err := db.GetDatabase(constants.DBname)
	if err != nil{
		fmt.Printf("An error has occured!")
	}

	// Stage 2: Check if users input path exists
	// absolutePath := interfaces.Formatter.ToAbs(UserPath)
	absolutePath, err := filepath.Abs(UserPath)
	if err != nil {
		fmt.Printf("An error has occured!")
	}

	if _, err = os.Stat(absolutePath); os.IsNotExist(err) {
		heavyFlow := flow.InitHeavyFlow(database, filepath.Base(UserPath))
		outPath, err = heavyFlow.Act()
	} else {
		lightFlow := flow.InitLightFlow(database, absolutePath)
		outPath, err = lightFlow.Act()
	}

	os.Stdout.WriteString(outPath)
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
