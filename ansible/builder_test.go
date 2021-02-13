package ansible_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/ansible"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

var allStages = []ansible.Stage{
	ansible.BootstrapStage,
	ansible.ChrootStage,
	ansible.MasterStage,
}

func TestBuilderCmd(t *testing.T) {
	cases := allStages
	for _, tc := range cases {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()
			b := ansible.NewBuilder(helpers.TmpDir(), tc)
			cmd := b.Cmd()

			// verify args
			require.Subset(t, cmd.Args, b.Args())
			require.Contains(t, cmd.Args, b.PlaybookPath)
			require.Contains(t, cmd.Args, ansible.AnsibleBinary)
			require.Contains(t, cmd.Args, ansible.BecomeFlag)

			// verify env
			require.Subset(t, cmd.Env, b.Env())
		})
	}
}
func TestGetAnsibleVerboseFlag(t *testing.T) {
	cases := []struct {
		name     string
		logLevel log.Level
		expected string
	}{
		{
			name:     "empty flag",
			logLevel: log.InfoLevel,
			expected: ansible.NoFlag,
		},
		{
			name:     "verbose flag",
			logLevel: log.DebugLevel,
			expected: ansible.VerboseFlag,
		},
		{
			name:     "extra verbose flag",
			logLevel: log.TraceLevel,
			expected: ansible.ExtraVerboseFlag,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			log.SetLevel(tc.logLevel)
			assert.Equal(t, ansible.GetAnsibleVerboseFlag(), tc.expected)
		})
	}
}

func TestCredCacheDisabled(t *testing.T) {
	for _, tc := range allStages {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()
			b := ansible.NewBuilder(helpers.TmpDir(), tc)
			require.False(t, b.IsCredCacheEnabled())
			require.NotContains(t, b.Env(), fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", b.VaultPasswordFile))
			require.Contains(t, b.Args(), ansible.BecomeFlag)
		})
	}
}

func TestCredCacheEnabled(t *testing.T) {
	for _, tc := range allStages {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()

			// setup & cleanup
			tmpDir := helpers.TmpDir()
			b := ansible.NewBuilder(tmpDir, tc)
			setupCredCache(t, b)
			defer func() {
				_ = os.RemoveAll(tmpDir)
			}()
			require.True(t, b.IsCredCacheEnabled())
			require.NotContains(t, b.Args(), ansible.BecomeFlag)
			require.Contains(t, b.Env(), fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", b.VaultPasswordFile))
		})
	}
}

func setupCredCache(t *testing.T, b *ansible.Builder) {
	require.Nil(t, os.MkdirAll(b.GroupVarsDir, 0755))
	_, err := os.Create(b.VaultPasswordFile)
	require.Nil(t, err)
	_, err = os.Create(path.Join(b.GroupVarsDir, "all.yml"))
	require.Nil(t, err)
}

func TestIsRoot(t *testing.T) {
	for _, tc := range allStages {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()
			b := ansible.NewBuilder(helpers.TmpDir(), tc)
			require.False(t, b.IsRoot())
		})
	}
}
