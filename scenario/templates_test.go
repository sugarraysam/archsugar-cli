package scenario_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func TestTemplateFilesExist(t *testing.T) {
	files := []string{
		path.Join(scenario.TemplatesBasedir, scenario.VarsBasename),
		path.Join(scenario.TemplatesBasedir, scenario.TasksBasename),
	}
	for _, f := range files {
		_, err := os.Stat(f)
		require.Nil(t, err)
	}
}

func TestVarsTemplate(t *testing.T) {
	name := getRandomName()
	expectedDst := path.Join(scenario.VarsBasedir, fmt.Sprintf("%s.yml", name))

	// GetDst
	tmpl := scenario.NewVarsTemplate(name)
	assert.Equal(t, tmpl.GetDst(), expectedDst)

	// Write + Exists
	require.False(t, tmpl.Exists())
	err := tmpl.Write()
	require.Nil(t, err)
	require.True(t, tmpl.Exists())

	// Rm + Exists
	err = tmpl.Rm()
	require.Nil(t, err)
	require.False(t, tmpl.Exists())
}

func TestTasksTemplate(t *testing.T) {
	name := getRandomName()
	desc := "random description"
	expectedDst := path.Join(scenario.TasksBasedir, fmt.Sprintf("%s.yml", name))

	// GetDst
	tmpl := scenario.NewTasksTemplate(name, desc)
	assert.Equal(t, tmpl.GetDst(), expectedDst)

	// Write + Exists
	require.False(t, tmpl.Exists())
	err := tmpl.Write()
	require.Nil(t, err)
	require.True(t, tmpl.Exists())

	// Rm + Exists
	err = tmpl.Rm()
	require.Nil(t, err)
	require.False(t, tmpl.Exists())
}
