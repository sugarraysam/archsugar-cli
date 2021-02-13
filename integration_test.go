package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

var (
	tmpDirBootstrap = helpers.TmpDir()
	tmpDirChroot    = helpers.TmpDir()
	tmpDirMaster    = helpers.TmpDir()
)

func TestIntegration(t *testing.T) {
	cases := []struct {
		tmpDir string
		pb     playbook.AnsiblePlaybook
	}{
		{
			tmpDir: tmpDirBootstrap,
			pb:     playbook.NewBootstrap(tmpDirBootstrap),
		},
		{

			tmpDir: tmpDirChroot,
			pb:     playbook.NewChroot(tmpDirChroot),
		},
		{
			tmpDir: tmpDirMaster,
			pb:     playbook.NewMaster(tmpDirMaster),
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.pb.Name(), func(t *testing.T) {
			t.Parallel()

			// setup & cleanup
			setup(t, tc.tmpDir)
			defer func() {
				_ = os.RemoveAll(tc.tmpDir)
			}()
			require.Nil(t, tc.pb.DryRun())
		})
	}
}

func setup(t *testing.T, tmpDir string) {
	repo, err := dotfiles.NewRepo(tmpDir, dotfiles.DefaultURL, "dev")
	require.Nil(t, err)
	require.Nil(t, repo.Clone())
}
