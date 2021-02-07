package scenario_test

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestCreate(t *testing.T) {
	name := getRandomName()
	expectedFiles := []string{
		path.Join(scenario.VarsBasedir, fmt.Sprintf("%s.yml", name)),
		path.Join(scenario.TasksBasedir, fmt.Sprintf("%s.yml", name)),
	}
	s, err := scenario.NewUncreatedScenario(name, "")
	require.Nil(t, err)
	err = s.Create()
	require.Nil(t, err)

	// check files were created
	for _, f := range expectedFiles {
		_, err := os.Stat(f)
		assert.Nil(t, err)
	}

	// clean
	for _, f := range expectedFiles {
		_ = os.Remove(f)
	}
}

func TestErrScenarioAlreadyExists(t *testing.T) {
	name := getRandomName()
	s, err := scenario.NewUncreatedScenario(name, "")
	require.Nil(t, err)

	err = s.Create()
	require.Nil(t, err)

	_, err = scenario.NewUncreatedScenario(name, "")
	require.True(t, errors.Is(err, scenario.ErrScenarioAlreadyExists))
}
