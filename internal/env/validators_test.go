package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func MockEnvironment() {
	os.Setenv("_INTELLIPATH_DIR", "/mock/dir")
	os.Setenv("_INTELLIPATH_DB_DTIMER", "5")
}

// RestoreEnvironment restores the original environment after testing.
func RestoreEnvironment() {
	os.Unsetenv("_INTELLIPATH_DIR")
	os.Unsetenv("_INTELLIPATH_DB_DTIMER")
}

func TestValidateENVs_Success(t *testing.T) {
	// Arrange
	MockEnvironment()
	defer RestoreEnvironment()

	validator := NewValidator()

	// Act
	err := validator.ValidateENVs()

	// Assert
	assert.NoError(t, err, "ValidateENVs should succeed with mock environment variables")
}

func TestValidateENVs_MissingIntellipathDir(t *testing.T) {
	// Arrange
	os.Unsetenv("_INTELLIPATH_DIR")

	validator := NewValidator()

	// Act
	err := validator.ValidateENVs()

	// Assert
	assert.Error(t, err, "_INTELLIPATH_DIR not found error should be returned")
	assert.EqualError(t, err, "_INTELLIPATH_DIR not found within environmental variables")
}

func TestValidateENVs_MissingIntellipathDTimer(t *testing.T) {
	// Arrange
	MockEnvironment()
	defer RestoreEnvironment()
	os.Unsetenv("_INTELLIPATH_DB_DTIMER")

	validator := NewValidator()

	// Act
	err := validator.ValidateENVs()

	// Assert
	assert.Error(t, err, "_INTELLIPATH_DB_DTIMER not found error should be returned")
	assert.EqualError(t, err, "_INTELLIPATH_DB_DTIMER not found within environmental variables")
}