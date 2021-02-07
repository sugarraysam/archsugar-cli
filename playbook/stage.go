package playbook

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"

	log "github.com/sirupsen/logrus"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

const (
	AnsibleBinary    = "ansible-playbook"
	BecomeFlag       = "-K"
	VerboseFlag      = "-v"
	ExtraVerboseFlag = "-vvv"
	NoFlag           = ""
)

var (
	DefaultArgs = []string{"-e target=localhost"}
	DefaultEnv  = []string{
		"ANSIBLE_FORCE_COLOR=true",
		"ANSIBLE_NOCOWS=1",
	}
	ChrootBareEnv = []string{
		"ANSIBLE_CHROOT_EXE=/usr/bin/arch-chroot",
		"ANSIBLE_EXECUTABLE=/usr/bin/bash",
	}
	ChrootBareArgs    = []string{"-c", "chroot", "-e target=all", "-i", "/mnt,"}
	PlaybookPath      = path.Join(helpers.BaseDir, "playbook.yml")
	VaultPasswordFile = path.Join(helpers.BaseDir, "vault_password_file.sh")
	GroupVarsDir      = path.Join(helpers.BaseDir, "group_vars")
)

// Stage - type representing each of the three archsugar stages
type Stage int

const (
	Bootstrap Stage = iota
	Chroot
	Master
)

func (s Stage) String() string {
	return []string{"bootstrap", "chroot", "master"}[s]
}

func (s Stage) ToEnv() string {
	return fmt.Sprintf("SUGAR_STAGE=%s", s.String())
}

// Cmd - returns cmd for given stage
func (s Stage) Cmd() *exec.Cmd {
	cmd := exec.Command(AnsibleBinary, s.Args()...)
	cmd.Env = append(cmd.Env, s.Env()...)
	return cmd
}

// Args - returns cmd args for a specific stage
func (s Stage) Args() []string {
	res := DefaultArgs
	// Bare installation
	if s == Chroot && !isVagrantInstall() {
		res = ChrootBareArgs
	}
	if GetAnsibleVerboseFlag() != NoFlag {
		res = append(res, GetAnsibleVerboseFlag())
	}
	if !isRoot() && !isCredCacheEnabled() {
		res = append(res, BecomeFlag)
	}
	return append(res, PlaybookPath)
}

// Env - returns cmd env for a specific stage
func (s Stage) Env() []string {
	res := DefaultEnv
	if s == Chroot && !isVagrantInstall() {
		res = ChrootBareEnv
	}
	res = append(res, os.Environ()...)
	res = append(res, s.ToEnv())

	if isCredCacheEnabled() && !isRoot() {
		res = append(res, fmt.Sprintf("ANSIBLE_VAULT_PASSWORD_FILE=%s", VaultPasswordFile))
	}
	return res
}

func isVagrantInstall() bool {
	return os.Getenv("SUGAR_VAGRANT") == "true"
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

// isRoot - detects if user running archsugar is root
func isRoot() bool {
	user, err := user.Current()
	if err != nil {
		// Should not happen - crash program if it does
		panic(err)
	}
	return user.Uid == "0"
}

// detect if both `vault_password_file.sh` and `group_vars/all.yml` exist
func isCredCacheEnabled() bool {
	files := []string{
		VaultPasswordFile,
		path.Join(GroupVarsDir, "all.yml"),
	}
	for _, filename := range files {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return false
		}
	}
	return true
}
