package scenario_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Create dir structure under /tmp to avoid cluttering the real project
// also provide some helper funcs
func TestMain(m *testing.M) {
	// Setup - create new directory structure under /tmp
	tmpDir := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", helpers.GetRandomDigit()))
	helpers.BaseDir = tmpDir
	scenario.VarsBasedir = path.Join(tmpDir, "roles/main/vars/master")
	scenario.TasksBasedir = path.Join(tmpDir, "roles/main/tasks/master")
	scenario.EnabledBasedir = path.Join(scenario.TasksBasedir, "enabled")

	for _, path := range []string{
		scenario.VarsBasedir,
		scenario.TasksBasedir,
		scenario.EnabledBasedir,
	} {
		if err := os.MkdirAll(path, 0755); err != nil {
			os.Exit(1)
		}
	}

	// Run
	rc := m.Run()

	// Teardown
	_ = os.RemoveAll(helpers.BaseDir)
	os.Exit(rc)
}

func getRandomName() string {
	return fmt.Sprintf("neverdont-%d", helpers.GetRandomDigit())
}

var randomDescriptions = []string{
	"who do you be?",
	"how much does it look like?",
	"I will code you the moon",
	"where is your horse?",
	"",
	"how much would could a programmer chop chop chop?, how much would could a programmer chop chop chop?",
}

func getRandomDescription() string {
	i := rand.Intn(len(randomDescriptions))
	return randomDescriptions[i]
}

// returns description and name
func createRandomScenario(t *testing.T) (string, string) {
	for {
		name := getRandomName()
		desc := getRandomDescription()
		s, err := scenario.NewUncreatedScenario(name, desc)
		// avoid flaky test with infrequent collisions in random names
		if err == nil {
			require.Nil(t, s.Create())
			return name, desc
		}
	}
}

type dummyScenario struct {
	name    string
	desc    string
	enabled bool
}

func newRandomDummyScenario(t *testing.T) *dummyScenario {
	name, desc := createRandomScenario(t)
	enabled := rand.Intn(2) == 0
	s, err := scenario.NewCreatedScenario(name)
	require.Nil(t, err)
	if enabled {
		err = s.Enable()
		require.Nil(t, err)
	}
	return &dummyScenario{
		name:    name,
		desc:    desc,
		enabled: enabled,
	}
}

// dummyScenarios - implement sort interface
type dummyScenarios []*dummyScenario

func (d dummyScenarios) Len() int {
	return len(d)
}

func (d dummyScenarios) Less(i, j int) bool {
	return strings.ToLower(d[i].name) < strings.ToLower(d[j].name)
}

func (d dummyScenarios) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
