package main_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMain(m *testing.M) {
	// Setup - create new directory structure under /tmp
	tmpDir := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", helpers.GetRandomDigit()))
	helpers.BaseDir = tmpDir

	// Run
	rc := m.Run()

	// Teardown
	// _ = os.RemoveAll(tmpDir)
	os.Exit(rc)
}

func TestIntegration(t *testing.T) {
	// clone dotfiles to helpers.BaseDir
	repo, err := dotfiles.NewRepo(dotfiles.DefaultURL, "dev")
	require.Nil(t, err)
	require.Nil(t, repo.Clone())

	playbooks := map[string]playbook.AnsiblePlaybook{
		"bootstrap": playbook.NewBootstrapPlaybook(),
		"chroot":    playbook.NewChrootPlaybook(),
		"master":    playbook.NewMasterPlaybook(),
	}

	for name, playbook := range playbooks {
		playbook := playbook
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			require.Nil(t, playbook.DryRun())
		})
	}
}
