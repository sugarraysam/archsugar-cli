package playbook_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMain(m *testing.M) {
	// Setup - create new directory structure under /tmp
	tmpDir := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", helpers.GetRandomDigit()))
	helpers.BaseDir = tmpDir
	playbook.VaultPasswordFile = path.Join(tmpDir, "vault_password_file.sh")
	playbook.GroupVarsDir = path.Join(tmpDir, "group_vars")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		os.Exit(1)
	}

	// Run
	rc := m.Run()

	// Teardown
	_ = os.RemoveAll(tmpDir)
	os.Exit(rc)
}
