package pathfinder

import (
	"intellipath/internal/record"
	"intellipath/internal/record/db"

)

// Each flow should implement this
type PathMatcher interface {
	FindMatch() string
}

type FlowManager struct {
	db db.Database
}

func NewFlowManager(db db.Database) *FlowManager {
	return &FlowManager{db: db}
}

func (fm *FlowManager) Manage(userInput string) string {
	// Validate input

	chosenFlow := fm.findFlow(userInput)
	outPath := chosenFlow.FindMatch()
	return outPath
}

func (fm *FlowManager) findFlow(userInput string) PathMatcher {
	pathRecFormatter := record.NewPathFormatter()
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
