package helpers_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

func TestRandomDigit(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		d := helpers.RandomDigit()
		require.Less(t, d, helpers.MaxDigit)
		require.GreaterOrEqual(t, d, helpers.MinDigit)
	}
}

func TestTmpDir(t *testing.T) {
	n := 10
	osTemp := os.TempDir()
	for i := 0; i < n; i++ {
		dir := helpers.TmpDir()
		require.Contains(t, dir, "archsugar-")
		require.Contains(t, dir, osTemp)
	}
}
