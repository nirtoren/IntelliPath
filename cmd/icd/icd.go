/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package icd

import (
	"fmt"
	"intellipath/internal/constants"
	"intellipath/internal/db"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"
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

	fmt.Println("icd called")

	// Stage 1: get the db
	database, err := db.GetDatabase(constants.DBname)
	if err != nil{
		fmt.Printf("An error has occured!")
	}

	// Stage 2: try to 'cd' into users input
	err = os.Chdir(UserPath)
	if err != nil {
		// Light flow
		// Check absolute path in db
		// if in DB -> Score up & Act
	} else {
		// Heavy flow
		// Check in DB + fuzzy + levinshtein
		// if exists -> get result -> try cd -> delete path if fails / Score up & Act.
		// if does not exists -> fail the process ( As 'cd' would fail )
	}

	// Stage 3: if err == nil than check path in db + fuzzy ...
	// If err != nil, meaning the path exists relativly to where the user is on,
	// Check db, if exists -> Score up and act, if doesnt exist, Save!

	var paths []string
	paths, err = database.GetAllPaths()
	if err != nil{
		fmt.Printf("An error has occured!")
	}

	// fmt.Println(fuzzy.RankFind(UserPath, paths))
	fmt.Println(fuzzy.FindNormalized(UserPath, paths))
	// fmt.Println(paths)

	// database.InsertPath(UserPath)

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
