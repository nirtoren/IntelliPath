package pathfinder

type Direct struct {
	absolutePath string
}

func NewDirectFlow(absolutePath string) *Direct {
	return &Direct{
		absolutePath: absolutePath,
	}
}

func (direct *Direct) FindMatch() string { // This should later on return a record
	var outPath string

	rec, err := direct.pathsdb.PathSearch(direct.absolutePath) // This should return a record if it exists
	if err != nil {
		return "", err
	}

	switch rec.GetPath() {
	case "": // In case no record was found
		record, err := record.NewRecord(direct.absolutePath, 0)
		if err != nil {
			return ""
		}

		if _, err = direct.pathsdb.InsertRecord(record); err != nil {
			return ""
		}
		outPath = direct.absolutePath

	case direct.absolutePath: // In case a matching record was found
		if err := direct.pathsdb.UpdateScore(rec); err != nil {
			return ""
		}
		outPath = direct.absolutePath

	}

	return outPath
}
