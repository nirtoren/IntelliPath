// matcher/matcher_test.go
package matcher

import (
	"intellipath/internal/record"
	"testing"
)

func TestFuzzyFind(t *testing.T) {
	// Helper function to create a PathRecord
	createRecord := func(path string, score int) *record.PathRecord {
		rec,_ := record.NewRecord(path, score)
		return rec
	}

	tests := []struct {
		name          string
		inputPath     string
		inputRecords  []*record.PathRecord
		expectedPaths []PathDist
		expectError   bool
	}{
		{
			name:      "no match",
			inputPath: "some/random/path",
			inputRecords: []*record.PathRecord{
				createRecord("/home/user/music", 10),
				createRecord("/home/user/docs", 20),
			},
			expectedPaths: nil,
			expectError:   true,
		},
		{
			name:      "single match",
			inputPath: "music",
			inputRecords: []*record.PathRecord{
				createRecord("/home/user/music", 10),
				createRecord("/home/user/docs", 20),
			},
			expectedPaths: []PathDist{
				{Path: "/home/user/music", LevDistance: 0},
			},
			expectError: false,
		},
		{
			name:      "multiple matches",
			inputPath: "mus",
			inputRecords: []*record.PathRecord{
				createRecord("/home/user/music", 10),
				createRecord("/docs/fser/musica", 20),
			},
			expectedPaths: []PathDist{
				{Path: "/home/user/music", LevDistance: 2},
				{Path: "/docs/fser/musica", LevDistance: 3},
			},
			expectError: false,
		},
		{
			name:      "no matching base path",
			inputPath: "random",
			inputRecords: []*record.PathRecord{
				createRecord("/home/user/music", 10),
				createRecord("/home/user/docs", 20),
				createRecord("/home/user/videos", 30),
			},
			expectedPaths: nil,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuzzyFind(tt.inputPath, tt.inputRecords)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("did not expect an error but got one: %v", err)
			}

			if len(got) != len(tt.expectedPaths) {
				t.Errorf("expected %d results, but got %d", len(tt.expectedPaths), len(got))
			}

			for i, expectedPath := range tt.expectedPaths {
				if got[i].Path != expectedPath.Path || got[i].LevDistance != expectedPath.LevDistance {
					t.Errorf("expected %v, but got %v", expectedPath, got[i])
				}
			}
		})
	}
}