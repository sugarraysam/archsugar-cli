package helpers_test

import (
	"os/user"
	"testing"

	"github.com/sugarraysam/archsugar-cli/helpers"

	"github.com/stretchr/testify/require"
)

func TestBaseDir(t *testing.T) {
	user, err := user.Current()
	require.Nil(t, err)
	require.Contains(t, helpers.BaseDir, user.Username)
	require.Contains(t, helpers.BaseDir, helpers.ArchsugarDefaultPath)
}
