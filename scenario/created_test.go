package scenario_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestErrScenarioDoesNotExist(t *testing.T) {
	_, err := scenario.NewCreatedScenario(getRandomName())
	require.True(t, errors.Is(err, scenario.ErrScenarioDoesNotExist))
}

func TestNewCreatedScenarioAndRm(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		name, desc := createRandomScenario(t)
		t.Run(name, func(t *testing.T) {
			s, err := scenario.NewCreatedScenario(name)
			require.Nil(t, err)
			assert.Equal(t, s.GetName(), name)
			assert.Equal(t, s.GetDesc(), desc)
			err = s.Rm()
			require.Nil(t, err)
		})
	}
}

func TestEnableDisable(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		name, desc := createRandomScenario(t)
		t.Run(name, func(t *testing.T) {
			s, err := scenario.NewCreatedScenario(name)
			require.Nil(t, err)
			require.Equal(t, s.GetName(), name)
			require.Equal(t, s.GetDesc(), desc)
			require.False(t, s.IsEnabled())

			// Enable
			err = s.Enable()
			require.Nil(t, err)
			require.True(t, s.IsEnabled())

			// Disable
			err = s.Disable()
			require.Nil(t, err)
			require.False(t, s.IsEnabled())
		})
	}
}
