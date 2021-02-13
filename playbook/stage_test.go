package playbook_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func TestStages(t *testing.T) {
	cases := []struct {
		stage          playbook.Stage
		expectedString string
		expectedArgs   []string
		expectedEnv    []string
	}{
		{
			stage:          playbook.Bootstrap,
			expectedString: "bootstrap",
			expectedArgs:   playbook.DefaultArgs,
			expectedEnv:    playbook.DefaultEnv,
		},
		{
			stage:          playbook.Chroot,
			expectedString: "chroot",
			expectedArgs:   playbook.ChrootArgs,
			expectedEnv:    append(playbook.DefaultEnv, playbook.ChrootExtraEnv...),
		},
		{
			stage:          playbook.Master,
			expectedString: "master",
			expectedArgs:   playbook.DefaultArgs,
			expectedEnv:    playbook.DefaultEnv,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.expectedString, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.stage.String(), tc.expectedString)
			require.ElementsMatch(t, tc.stage.DefaultArgs(), tc.expectedArgs)
			require.Subset(t, tc.stage.DefaultEnv(), tc.expectedEnv)
			require.Contains(t, tc.stage.DefaultEnv(), fmt.Sprintf("SUGAR_STAGE=%s", tc.expectedString))
		})
	}
}
