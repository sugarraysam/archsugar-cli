package scenario_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestReadAllScenarios(t *testing.T) {
	// Create 'n' scenarios, and save name + desc in a map
	n := 5
	expected := make(map[string]string)
	for i := 0; i < n; i++ {
		s, err := helpers.CreateRandomScenario(t, TmpDir)
		require.Nil(t, err)
		expected[s.Name] = s.Desc
	}

	xs, err := scenario.ReadAll(TmpDir)
	require.Nil(t, err)
	for _, s := range xs {
		require.Equal(t, expected[s.Name], s.Desc)
		require.Nil(t, s.Rm())
	}
}

func TestAllScenariosAreSorted(t *testing.T) {
	// Create 'n' scenarios, and save names
	n := 5
	var expected []string
	for i := 0; i < n; i++ {
		s, err := helpers.CreateRandomScenario(t, TmpDir)
		require.Nil(t, err)
		expected = append(expected, s.Name)
	}
	sort.Strings(expected)

	xs, err := scenario.ReadAll(TmpDir)
	require.Nil(t, err)
	var names []string
	for _, s := range xs {
		names = append(names, s.Name)
		require.Nil(t, s.Rm())
	}
	assert.Equal(t, names, expected)
}

func TestListOutput(t *testing.T) {
	// create n scenarios
	n := 5
	var expected []*scenario.Scenario
	for i := 0; i < n; i++ {
		s, err := helpers.CreateRandomScenario(t, TmpDir)
		require.Nil(t, err)
		expected = append(expected, s)
	}

	// verify all created scenarios Name present in output
	xs, err := scenario.ReadAll(TmpDir)
	require.Nil(t, err)
	require.Equal(t, len(expected), len(xs))
	var b strings.Builder
	require.Nil(t, xs.List(&b))
	out := b.String()
	for _, s := range expected {
		require.Contains(t, out, s.Name)
	}

	// clean
	for _, s := range xs {
		require.Nil(t, s.Rm())
	}
}

func TestEnableAllAndDisableAll(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		_, _ = helpers.CreateRandomScenario(t, TmpDir)
	}
	xs, err := scenario.ReadAll(TmpDir)
	require.Nil(t, err)
	for _, s := range xs {
		require.False(t, s.IsEnabled())
	}

	// EnableAll
	require.Nil(t, xs.EnableAll())
	for _, s := range xs {
		require.True(t, s.IsEnabled())
	}

	// DisableAll
	require.Nil(t, xs.DisableAll())
	for _, s := range xs {
		require.False(t, s.IsEnabled())

		// clean
		require.Nil(t, s.Rm())
	}
}
