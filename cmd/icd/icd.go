/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	flow "intellipath/internal/flows"
	"intellipath/internal/record"
	"intellipath/internal/utils"
	"os"
	"strconv"

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
	var err error

	pathFormatter := utils.NewPathFormatter()
	validator := utils.NewValidator()
	validator.ValidateENV()
	
	userPath := args[0]

	// Get the db
	value := os.Getenv(constants.INTELLIPATH_DIR)
	
	database := record.GetDbInstance(value+constants.DBpath)
	// ADD ERROR IN CASE DATABASE NOT FOUND

	defer database.Close()

	// Parallel cleanup of un-touched paths
	resultCh := make(chan error)
	dtimer, _ := strconv.Atoi(constants.INTELLIPATH_DB_DTIMER)
	go record.ParallelCleanUp(database, dtimer, resultCh)


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
