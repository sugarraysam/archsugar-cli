package ansible_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	ansible "github.com/sugarraysam/archsugar-cli/ansible"
)

func TestStages(t *testing.T) {
	cases := []struct {
		stage          ansible.Stage
		expectedString string
		expectedArgs   []string
		expectedEnv    []string
	}{
		{
			stage:          ansible.BootstrapStage,
			expectedString: "bootstrap",
			expectedArgs:   ansible.DefaultArgs,
			expectedEnv:    ansible.DefaultEnv,
		},
		{
			stage:          ansible.ChrootStage,
			expectedString: "chroot",
			expectedArgs:   ansible.ChrootArgs,
			expectedEnv:    append(ansible.DefaultEnv, ansible.ChrootExtraEnv...),
		},
		{
			stage:          ansible.MasterStage,
			expectedString: "master",
			expectedArgs:   ansible.DefaultArgs,
			expectedEnv:    ansible.DefaultEnv,
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
