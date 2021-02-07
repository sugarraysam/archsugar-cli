package helpers_test

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/sugarraysam/archsugar-cli/helpers"

	"github.com/stretchr/testify/assert"
)

func TestBaseDir(t *testing.T) {
	cwd, _ := filepath.Abs(".")
	want := path.Dir(cwd)
	assert.Equal(t, want, helpers.BaseDir)
}
