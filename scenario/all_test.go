package scenario_test

import (
	"bufio"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestGetAllScenarios(t *testing.T) {
	// Create 'n' scenarios, and save name + desc in a map
	n := 7
	expected := make(map[string]string)
	for i := 0; i < n; i++ {
		name, desc := createRandomScenario(t)
		expected[name] = desc
	}

	xs, err := scenario.GetAllScenarios()
	require.Nil(t, err)
	for _, s := range xs.GetScenarios() {
		assert.Equal(t, expected[s.GetName()], s.GetDesc())
		// Clean
		err = s.Rm()
		require.Nil(t, err)
	}
}

func TestGetAllScenariosAreSorted(t *testing.T) {
	// Create 'n' scenarios, and save names
	n := 7
	var expected []string
	for i := 0; i < n; i++ {
		name, _ := createRandomScenario(t)
		expected = append(expected, name)
	}
	sort.Strings(expected)

	xs, err := scenario.GetAllScenarios()
	require.Nil(t, err)
	var names []string
	for _, s := range xs.GetScenarios() {
		names = append(names, s.GetName())
		// Clean
		err = s.Rm()
		require.Nil(t, err)
	}
	assert.Equal(t, names, expected)
}

func TestListOutput(t *testing.T) {
	// init n dummyScenarios ++ sort
	n := 7
	var expected dummyScenarios
	for i := 0; i < n; i++ {
		expected = append(expected, newRandomDummyScenario(t))
	}
	sort.Sort(expected)

	// list all scenarios to my writer
	xs, err := scenario.GetAllScenarios()
	require.Nil(t, err)
	require.Equal(t, len(expected), len(xs.GetScenarios()))
	var b strings.Builder
	err = xs.List(&b)
	require.Nil(t, err)

	// verify list output is accurate
	scanner := bufio.NewScanner(strings.NewReader(b.String()))
	for _, expectedLine := range []string{scenario.Header, scenario.Separator} {
		scanner.Scan()
		require.True(t, strings.HasPrefix(scanner.Text(), expectedLine))
	}
	for i := 0; i < len(expected) && scanner.Scan(); i++ {
		tc := expected[i]
		expectedLine := scenario.FormatLine(tc.name, scenario.DisabledToken, tc.desc)
		if tc.enabled {
			expectedLine = scenario.FormatLine(tc.name, scenario.EnabledToken, tc.desc)
		}
		require.True(t, strings.HasPrefix(scanner.Text(), expectedLine))
	}
	require.Nil(t, scanner.Err())

	// clean
	for _, s := range xs.GetScenarios() {
		err = s.Rm()
		require.Nil(t, err)
	}
}

func TestEnableAndDisableAll(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		_, _ = createRandomScenario(t)
	}
	xs, err := scenario.GetAllScenarios()
	require.Nil(t, err)
	for _, s := range xs.GetScenarios() {
		require.False(t, s.IsEnabled())
	}

	// EnableAll
	err = xs.EnableAll()
	require.Nil(t, err)
	for _, s := range xs.GetScenarios() {
		require.True(t, s.IsEnabled())
	}

	// DisableAll
	err = xs.DisableAll()
	require.Nil(t, err)
	for _, s := range xs.GetScenarios() {
		require.False(t, s.IsEnabled())
		// clean
		err = s.Rm()
		require.Nil(t, err)
	}
}
