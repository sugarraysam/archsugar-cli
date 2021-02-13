package ansible

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"

	log "github.com/sirupsen/logrus"
)

const (
	AnsibleBinary    = "ansible-playbook"
	BecomeFlag       = "-K"
	VerboseFlag      = "-v"
	ExtraVerboseFlag = "-vvv"
	NoFlag           = ""
)

type Builder struct {
	Stage             Stage
	PlaybookPath      string
	VaultPasswordFile string
	GroupVarsDir      string
}

func NewBuilder(stage Stage, basePath string) *Builder {
	return &Builder{
		Stage:             stage,
		PlaybookPath:      path.Join(basePath, "playbook.yml"),
		VaultPasswordFile: path.Join(basePath, "vault_password_file.sh"),
		GroupVarsDir:      path.Join(basePath, "group_vars"),
	}
}

// Cmd - returns cmd for given stage
func (b *Builder) Cmd() *exec.Cmd {
	cmd := exec.Command(AnsibleBinary, b.Args()...)
	cmd.Env = append(cmd.Env, b.Env()...)
	return cmd
}

// Args - returns cmd args for a specific stage
func (b *Builder) Args() []string {
	res := b.Stage.DefaultArgs()
	if GetAnsibleVerboseFlag() != NoFlag {
		res = append(res, GetAnsibleVerboseFlag())
	}
	if !b.IsRoot() && !b.IsCredCacheEnabled() {
		res = append(res, BecomeFlag)
	}
	return append(res, b.PlaybookPath)
}

// Env - returns cmd env for a specific stage
func (b *Builder) Env() []string {
	res := b.Stage.DefaultEnv()
	res = append(res, os.Environ()...)
	if !b.IsRoot() && b.IsCredCacheEnabled() {
		res = append(res, fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", b.VaultPasswordFile))
	}
	return res
}

// GetAnsibleVerboseFlag - translate logLevel to a valid ansible flag
func GetAnsibleVerboseFlag() string {
	switch log.GetLevel() {
	case log.DebugLevel:
		return VerboseFlag
	case log.TraceLevel:
		return ExtraVerboseFlag
	default:
		return NoFlag
	}
}

// IsRoot - detects if user running archsugar is root
func (b *Builder) IsRoot() bool {
	user, err := user.Current()
	if err != nil {
		// Should not happen - crash program if it does
		panic(err)
	}
	return user.Uid == "0"
}

// detect if both `vault_password_file.sh` and `group_vars/all.yml` exist
func (b *Builder) IsCredCacheEnabled() bool {
	files := []string{
		b.VaultPasswordFile,
		path.Join(b.GroupVarsDir, "all.yml"),
	}
	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			return false
		}
	}
	return true
}
