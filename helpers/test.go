package helpers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/sugarraysam/archsugar-cli/scenario"
)

const (
	MinDigit = 10000
	MaxDigit = 20000
)

// TmpDir - generate a tmpDir, avoiding very rare collisions
func TmpDir() string {
	for {
		res := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", RandomDigit()))
		_, err := os.Stat(res)
		if err != nil {
			return res
		}
	}
}

func RandomDigit() int {
	return rand.Intn(MinDigit) + MaxDigit - MinDigit
}

func RandomScenarioName() string {
	return fmt.Sprintf("scenario-%d", RandomDigit())
}

var descriptions = []string{
	"who do you be?",
	"how much does it look like?",
	"I will code you the moon",
	"where is your horse?",
	"",
	"how much would could a programmer chop chop chop?, how much would could a programmer chop chop chop?",
}

func RandomScenarioDesc() string {
	i := rand.Intn(len(descriptions))
	return descriptions[i]
}

func CreateRandomScenario(t *testing.T, baseDir string) (*scenario.Scenario, error) {
	maxAttempts := 5
	try := 0
	for {
		name := RandomScenarioName()
		desc := RandomScenarioDesc()
		s, err := scenario.New(baseDir, name, desc)
		// avoid flaky test with infrequent collisions in random names
		if err == nil {
			return s, nil
		}
		try += 1
		if try >= maxAttempts {
			return nil, errors.New("could not create random scenario")
		}
	}
}
