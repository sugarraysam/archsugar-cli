package scenario_test

import (
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestNew(t *testing.T) {
	t.Parallel()

	s, err := helpers.CreateRandomScenario(t, TmpDir)
	require.Nil(t, err)

	// check files were created
	for _, tmpl := range s.Templates {
		require.True(t, tmpl.Exists())
		_, err := os.Stat(tmpl.GetDst())
		require.Nil(t, err)
	}

	require.Nil(t, s.Rm())
}

func TestErrScenarioAlreadyExists(t *testing.T) {
	t.Parallel()

	s, err := helpers.CreateRandomScenario(t, TmpDir)
	require.Nil(t, err)

	_, err = scenario.New(TmpDir, s.Name, s.Desc)
	require.True(t, errors.Is(err, scenario.ErrScenarioAlreadyExists))

	require.Nil(t, s.Rm())
}

func TestErrScenarioDoesNotExist(t *testing.T) {
	t.Parallel()

	_, err := scenario.Read(TmpDir, helpers.RandomScenarioName())
	require.True(t, errors.Is(err, scenario.ErrScenarioDoesNotExist))
}

func TestReadAndRm(t *testing.T) {
	t.Parallel()

	n := 5
	for i := 0; i < n; i++ {
		s1, err := helpers.CreateRandomScenario(t, TmpDir)
		require.Nil(t, err)
		s2, err := scenario.Read(TmpDir, s1.Name)
		require.Nil(t, err)

		assert.Equal(t, s1.Name, s2.Name)
		assert.Equal(t, s1.Desc, s2.Desc)
		assert.Equal(t, s1.Enabled, s2.Enabled)
		assert.Equal(t, s1.EnabledSymlink, s2.EnabledSymlink)
		require.Nil(t, s1.Rm())
		require.False(t, s2.Exists())
	}
}

func TestEnableDisable(t *testing.T) {
	t.Parallel()

	n := 5
	for i := 0; i < n; i++ {
		s, err := helpers.CreateRandomScenario(t, TmpDir)
		require.Nil(t, err)
		require.False(t, s.IsEnabled())

		// Enable
		require.Nil(t, s.Enable())
		require.True(t, s.IsEnabled())

		// Disable
		require.Nil(t, s.Disable())
		require.False(t, s.IsEnabled())

		// Rm
		require.Nil(t, s.Rm())
	}
}
