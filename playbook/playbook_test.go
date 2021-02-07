package playbook_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func TestBootstrapDryRun(t *testing.T) {
	t.Parallel()
	p := playbook.NewBootstrapPlaybook()
	err := p.DryRun()
	require.Nil(t, err)
}
func TestChrootDryRun(t *testing.T) {
	t.Parallel()
	p := playbook.NewChrootPlaybook()
	err := p.DryRun()
	require.Nil(t, err)
}

func TestMasterDryRun(t *testing.T) {
	t.Parallel()
	p := playbook.NewMasterPlaybook()
	err := p.DryRun()
	require.Nil(t, err)
}
