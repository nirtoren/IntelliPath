package pathfinder

import (
	"intellipath/internal/database"
)

// Each flow should implement this
type PathMatcher interface {
	FindMatch() string
}

type FlowManager struct {
	pathmatcher PathMatcher
	db database.Database
}

func NewFlowManager(db database.Database, pathmatcher PathMatcher) *FlowManager {
	return &FlowManager{pathmatcher: pathmatcher,
						db: db,}
}

func (fm *FlowManager) FindClosestPath(userInput string) string {
	fm.pathmatcher.FindMatch()
	return "hh"
}

func (fm *FlowManager) savePath() string {

	return "hh"
}

func (fm *FlowManager) updatePathScore() string {

	return "hh"
}


