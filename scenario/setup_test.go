package scenario_test

import (
	"os"
	"testing"

	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

var (
	TmpDir  = helpers.TmpDir()
	AllDirs = []string{
		scenario.TasksDir(TmpDir),
		scenario.VarsDir(TmpDir),
		scenario.EnabledDir(TmpDir),
	}
)

// Use a common directory structure for tests to limit I/O and speedup unit testing
func TestMain(m *testing.M) {
	for _, dir := range AllDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			os.Exit(1)
		}
	}

	// Run
	rc := m.Run()

	// Teardown
	_ = os.RemoveAll(TmpDir)
	os.Exit(rc)
}
