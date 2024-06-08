package record


type PathRecConvertor struct{}

func NewPathRecConvertor() *PathRecConvertor {
	return &PathRecConvertor{}
}

func (c *PathRecConvertor) RecToPath(record *PathRecord) string {
	return record.GetPath()
}

func (c *PathRecConvertor) RecordsToPaths(records []*PathRecord) []string {
	var paths []string

	for _, rec := range records {
		paths = append(paths, rec.GetPath())
	}

	return paths
}