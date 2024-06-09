package env

import (
	"os"
	"testing"

	"intellipath/internal/constants"
)


func TestValidateENVs(t *testing.T) {
	// Helper function to set and unset environment variables
	setEnv := func(key, value string) {
		err := os.Setenv(key, value)
		if err != nil {
			t.Fatalf("Failed to set environment variable %s: %v", key, err)
		}
	}

	unsetEnv := func(key string) {
		err := os.Unsetenv(key)
		if err != nil {
			t.Fatalf("Failed to unset environment variable %s: %v", key, err)
		}
	}

	// Test when both environment variables are set
	t.Run("Both ENVs set", func(t *testing.T) {
		setEnv(constants.INTELLIPATH_DIR, "/path/to/intellipath")
		setEnv(constants.INTELLIPATH_DB_DTIMER, "7")

		defer unsetEnv(constants.INTELLIPATH_DIR)
		defer unsetEnv(constants.INTELLIPATH_DB_DTIMER)

		validator := NewValidator()
		validator.ValidateENVs() // Should not panic

		// If the function returns without panicking, the test is successful
	})

	// Test when INTELLIPATH_DIR is not set
	t.Run("INTELLIPATH_DIR not set", func(t *testing.T) {
		unsetEnv(constants.INTELLIPATH_DIR)
		setEnv(constants.INTELLIPATH_DB_DTIMER, "7")

		defer unsetEnv(constants.INTELLIPATH_DB_DTIMER)

		validator := NewValidator()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic when INTELLIPATH_DIR was not set")
			}
		}()
		validator.ValidateENVs() // Should panic
	})

	// Test when INTELLIPATH_DB_DTIMER is not set
	t.Run("INTELLIPATH_DB_DTIMER not set", func(t *testing.T) {
		setEnv(constants.INTELLIPATH_DIR, "/path/to/intellipath")
		unsetEnv(constants.INTELLIPATH_DB_DTIMER)

		defer unsetEnv(constants.INTELLIPATH_DIR)

		validator := NewValidator()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic when INTELLIPATH_DB_DTIMER was not set")
			}
		}()
		validator.ValidateENVs() // Should panic
	})

	// Test when both environment variables are not set
	t.Run("Both ENVs not set", func(t *testing.T) {
		unsetEnv(constants.INTELLIPATH_DIR)
		unsetEnv(constants.INTELLIPATH_DB_DTIMER)

		validator := NewValidator()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic when both environment variables were not set")
			}
		}()
		validator.ValidateENVs() // Should panic
	})
}