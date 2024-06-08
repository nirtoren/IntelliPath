package pathfinder

import (
	"intellipath/internal/database"
	"intellipath/internal/utils"
)

// Each flow should implement this
type PathMatcher interface {
	FindMatch() string
}

type FlowManager struct {
	db database.Database
}

func NewFlowManager(db database.Database) *FlowManager {
	return &FlowManager{db: db}
}

func (fm *FlowManager) Manage(userInput string) string {

	chosenFlow := fm.findFlow(userInput)
	chosenFlow.FindMatch()
	return "hh"
}

func (fm *FlowManager) findFlow(userInput string) PathMatcher {
	pathRecFormatter := utils.NewPathFormatter()
	absolutePath := pathRecFormatter.ToAbs(userInput)
	isPathExists := pathRecFormatter.IsExists(absolutePath)

	var pathMatcher PathMatcher
	if isPathExists {
		pathMatcher = NewDirectFlow(absolutePath, fm.db)
	} else {
		pathMatcher = NewFuzzyFlow(pathRecFormatter.ToBase(userInput), fm.db)
	}

	return pathMatcher
}

func (fm *FlowManager) savePath() string {

	return "hh"
}

func (fm *FlowManager) updatePathScore() string {

	return "hh"
}

// func (fm *FlowManager) databaseCleanUp() {
// 	var err error
// 	resultCh := make(chan error)
// 	dtimer, _ := strconv.Atoi(constants.INTELLIPATH_DB_DTIMER)
// 	go database.ParallelCleanUp(fm.db, dtimer, resultCh)

// 	err = <-resultCh
// 	if err != nil {
// 		fmt.Println("error while cleaning up old records.")
// 	}
// }
