package playbook_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

var allStages = []playbook.Stage{
	playbook.Bootstrap,
	playbook.Chroot,
	playbook.Master,
}

func TestBuilderCmd(t *testing.T) {
	cases := allStages
	for _, tc := range cases {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()
			b := playbook.NewBuilder(tc, helpers.TmpDir())
			cmd := b.Cmd()

			// verify args
			require.Subset(t, cmd.Args, b.Args())
			require.Contains(t, cmd.Args, b.PlaybookPath)
			require.Contains(t, cmd.Args, playbook.AnsibleBinary)
			require.Contains(t, cmd.Args, playbook.BecomeFlag)

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
			expected: playbook.NoFlag,
		},
		{
			name:     "verbose flag",
			logLevel: log.DebugLevel,
			expected: playbook.VerboseFlag,
		},
		{
			name:     "extra verbose flag",
			logLevel: log.TraceLevel,
			expected: playbook.ExtraVerboseFlag,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			log.SetLevel(tc.logLevel)
			assert.Equal(t, playbook.GetAnsibleVerboseFlag(), tc.expected)
		})
	}
}

func TestCredCacheDisabled(t *testing.T) {
	for _, tc := range allStages {
		tc := tc
		t.Run(tc.String(), func(t *testing.T) {
			t.Parallel()
			b := playbook.NewBuilder(tc, helpers.TmpDir())
			require.False(t, b.IsCredCacheEnabled())
			require.NotContains(t, b.Env(), fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", b.VaultPasswordFile))
			require.Contains(t, b.Args(), playbook.BecomeFlag)
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
			b := playbook.NewBuilder(tc, tmpDir)
			setupCredCache(t, b)
			defer func() {
				_ = os.RemoveAll(tmpDir)
			}()
			require.True(t, b.IsCredCacheEnabled())
			require.NotContains(t, b.Args(), playbook.BecomeFlag)
			require.Contains(t, b.Env(), fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", b.VaultPasswordFile))
		})
	}
}

func setupCredCache(t *testing.T, b *playbook.Builder) {
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
			b := playbook.NewBuilder(tc, helpers.TmpDir())
			require.False(t, b.IsRoot())
		})
	}
}
