package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/ansible"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

var (
	tmpDirBootstrap = helpers.TmpDir()
	tmpDirChroot    = helpers.TmpDir()
	tmpDirMaster    = helpers.TmpDir()
)

func TestIntegration(t *testing.T) {
	cases := []struct {
		tmpDir string
		p      *ansible.Playbook
	}{
		{
			tmpDir: tmpDirBootstrap,
			p:      ansible.NewBootstrapPlaybook(tmpDirBootstrap),
		},
		{

			tmpDir: tmpDirChroot,
			p:      ansible.NewChrootPlaybook(tmpDirChroot),
		},
		{
			tmpDir: tmpDirMaster,
			p:      ansible.NewMasterPlaybook(tmpDirMaster),
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.p.Name(), func(t *testing.T) {
			t.Parallel()

			// setup & cleanup
			setup(t, tc.tmpDir)
			defer func() {
				_ = os.RemoveAll(tc.tmpDir)
			}()
			require.Nil(t, tc.p.DryRun())
		})
	}
}

func setup(t *testing.T, tmpDir string) {
	repo, err := dotfiles.NewRepo(tmpDir, dotfiles.DefaultURL, "dev")
	require.Nil(t, err)
	require.Nil(t, repo.Clone())
}
