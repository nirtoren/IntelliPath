/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/record"
	"intellipath/internal/utils"
	flow "intellipath/internal/flows"
	"os"

	"github.com/spf13/cobra"
)

// icdCmd represents the icd command
var IcdCmd = &cobra.Command{
	Use:   "icd",
	Short: "Changing directory, a wrapper for 'cd' command",
	Long:  `Long for icd command`,
	Run:   RunIcd,
}

func RunIcd(cmd *cobra.Command, args []string) {
	var outPath string
	pathFormatter := utils.NewPathFormatter()

	userPath := args[0]

	// Get the db
	database, err := record.GetDatabase(constants.DBpath)
	if err != nil {
		fmt.Printf("An error has occured!")
	}

	defer database.Close()

	// Parallel cleanup of un-touched paths
	resultCh := make(chan error)
	go record.ParallelCleanUp(database, resultCh)


	// Check if users input actually path exists
	absolutePath := pathFormatter.ToAbs(userPath)
	isPathExists := pathFormatter.IsExists(absolutePath)

	if isPathExists {
		directFlow := flow.InitDirectFlow(database, absolutePath)
		outPath, err = directFlow.Act()
	} else {
		fuzzyFlow := flow.InitFuzzyFlow(database, pathFormatter.ToBase(userPath))
		outPath, err = fuzzyFlow.Act()
	}

	err = <-resultCh
	if err != nil {
		fmt.Println("error while cleaning up old records.")
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
