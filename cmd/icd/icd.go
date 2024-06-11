/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/env"
	"intellipath/internal/flows"
	"intellipath/internal/record/db"

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
	userInput := args[0]

	// Get GetENV<name>
	envValidator := env.NewValidator()
	envValidator.ValidateENVs()

	// Get the db, DEPENDS ON ENV
	database := db.GetDbInstance()
	defer database.Close()

	// Parallel cleanup of un-touched paths
	resultCh := make(chan error)
	dtimer, _ := strconv.Atoi(constants.INTELLIPATH_DB_DTIMER)
	go db.ParallelCleanUp(database, dtimer, resultCh)

	flowManager := pathfinder.NewFlowManager(database)
	outPath = flowManager.Manage(userInput)

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
