package scenario_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestTemplates(t *testing.T) {
	name := helpers.RandomScenarioName()
	desc := helpers.RandomScenarioDesc()

	varsTmpl := scenario.NewVarsTemplate(TmpDir, name)
	tasksTmpl := scenario.NewTasksTemplate(TmpDir, name, desc)

	for _, tmpl := range []scenario.AnsibleTemplate{varsTmpl, tasksTmpl} {
		require.False(t, tmpl.Exists())
		require.Nil(t, tmpl.Write())
		require.True(t, tmpl.Exists())
		require.Nil(t, tmpl.Rm())
		require.False(t, tmpl.Exists())
	}
}
