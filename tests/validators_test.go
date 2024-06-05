package internal_test

import (
	"os"
	"testing"

	"intellipath/internal/constants"
	"intellipath/internal/utils"
	"github.com/stretchr/testify/assert"
)


func TestValidateENV(t *testing.T) {
	validator := utils.NewValidator()
	validator.ValidateENV()
	
	t.Run("Environment variables exist", func(t *testing.T) {
		os.Setenv(constants.INTELLIPATH_DIR, "/path/to/intellipath")
		os.Setenv(constants.INTELLIPATH_DB_DTIMER, "some_value")
		defer os.Unsetenv(constants.INTELLIPATH_DIR)
		defer os.Unsetenv(constants.INTELLIPATH_DB_DTIMER)

		err := validator.ValidateENV()
		assert.NoError(t, err)
	})

	t.Run("INTELLIPATH_DIR not set", func(t *testing.T) {
		os.Unsetenv(constants.INTELLIPATH_DIR)
		os.Setenv(constants.INTELLIPATH_DB_DTIMER, "some_value")
		defer os.Unsetenv(constants.INTELLIPATH_DB_DTIMER)

		err := validator.ValidateENV()
		assert.EqualError(t, err, "_INTELLIPATH_DIR not found within environmental variables")
	})

	t.Run("INTELLIPATH_DB_DTIMER not set", func(t *testing.T) {
		os.Setenv(constants.INTELLIPATH_DIR, "/path/to/intellipath")
		os.Unsetenv(constants.INTELLIPATH_DB_DTIMER)
		defer os.Unsetenv(constants.INTELLIPATH_DIR)

		err := validator.ValidateENV()
		assert.EqualError(t, err, "_INTELLIPATH_DB_DTIMER not found within environmental variables")
	})

	t.Run("No environment variables set", func(t *testing.T) {
		os.Unsetenv(constants.INTELLIPATH_DIR)
		os.Unsetenv(constants.INTELLIPATH_DB_DTIMER)

		err := validator.ValidateENV()
		assert.EqualError(t, err, "_INTELLIPATH_DIR not found within environmental variables")
	})
}
