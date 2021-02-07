package playbook_test

import (
	"os"
	"path"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

var allStages = []playbook.Stage{
	playbook.Bootstrap,
	playbook.Chroot,
	playbook.Master,
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
func TestArgs(t *testing.T) {
	cases := []struct {
		stage       playbook.Stage
		defaultArgs []string
	}{
		{
			stage:       playbook.Bootstrap,
			defaultArgs: playbook.DefaultArgs,
		},
		{
			stage:       playbook.Chroot,
			defaultArgs: playbook.ChrootBareArgs,
		},
		{
			stage:       playbook.Master,
			defaultArgs: playbook.DefaultArgs,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.stage.String(), func(t *testing.T) {
			// check args start with expected default args
			args := tc.stage.Args()
			require.Equal(t, tc.defaultArgs, args[:len(tc.defaultArgs)])

			// find becomeFlag and playbookPath
			becomeFlagFound := false
			playbookPathFound := false
			for _, a := range args {
				if a == playbook.BecomeFlag {
					becomeFlagFound = true
				}
				if a == playbook.PlaybookPath {
					playbookPathFound = true
				}
			}
			require.True(t, becomeFlagFound)
			require.True(t, playbookPathFound)
		})
	}
}
func TestEnv(t *testing.T) {
	cases := []struct {
		stage      playbook.Stage
		defaultEnv []string
	}{
		{
			stage:      playbook.Bootstrap,
			defaultEnv: playbook.DefaultEnv,
		},
		{
			stage:      playbook.Chroot,
			defaultEnv: playbook.ChrootBareEnv,
		},
		{
			stage:      playbook.Master,
			defaultEnv: playbook.DefaultEnv,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.stage.String(), func(t *testing.T) {
			// check env start with expected defaultEnv
			env := tc.stage.Env()
			require.Equal(t, tc.defaultEnv, env[:len(tc.defaultEnv)])

			// make sure SUGAR_STAGE == s.String()
			for _, e := range env {
				if strings.HasPrefix(e, "SUGAR_STAGE") {
					require.Equal(t, e, tc.stage.ToEnv())
				}
			}
		})
	}
}

func TestCredCacheDisabled(t *testing.T) {
	for _, s := range allStages {
		s := s
		t.Run(s.String(), func(t *testing.T) {
			// env should NOT contain ANSIBLE_VAULT_PASSWORD_FILE
			for _, env := range s.Env() {
				require.False(t, strings.HasPrefix(env, "ANSIBLE_VAULT_PASSWORD_FILE"))
			}
			// args should contain BecomeFlag
			found := false
			for _, arg := range s.Args() {
				if arg == playbook.BecomeFlag {
					found = true
				}
			}
			require.True(t, found)
		})
	}
}

func TestCredCacheEnabled(t *testing.T) {
	// setup
	_, err := os.Create(playbook.VaultPasswordFile)
	require.Nil(t, err)
	err = os.Mkdir(playbook.GroupVarsDir, 0755)
	require.Nil(t, err)
	_, err = os.Create(path.Join(playbook.GroupVarsDir, "all.yml"))
	require.Nil(t, err)

	// cleanup
	defer func() {
		_ = os.Remove(playbook.VaultPasswordFile)
		_ = os.RemoveAll(playbook.GroupVarsDir)
	}()

	for _, s := range allStages {
		s := s
		t.Run(s.String(), func(t *testing.T) {
			// args should NOT contain BecomeFlag
			for _, arg := range s.Args() {
				require.NotEqual(t, arg, playbook.BecomeFlag)
			}
			// env should contain ANSIBLE_VAULT_PASSWORD_FILE
			found := false
			for _, env := range s.Env() {
				if strings.HasPrefix(env, "ANSIBLE_VAULT_PASSWORD_FILE") {
					found = true
				}
			}
			require.True(t, found)
		})
	}
}

func TestVagrantInstall(t *testing.T) {
	// setup && cleanup
	os.Setenv("SUGAR_VAGRANT", "true")
	defer func() { os.Unsetenv("SUGAR_VAGRANT") }()

	s := playbook.Chroot

	// make sure using DefaultArgs and not ChrootBareArgs
	args := s.Args()
	require.Equal(t, playbook.DefaultArgs, args[:len(playbook.DefaultArgs)])

	// make sure using DefaultEnv and not ChrootBareEnv
	env := s.Env()
	require.Equal(t, playbook.DefaultEnv, env[:len(playbook.DefaultEnv)])
}
