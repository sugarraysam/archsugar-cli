package playbook

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	DryRunFlag = "--list-tasks"
)

// AnsiblePlaybook - represents a playbook that will be executed by the CLI
type AnsiblePlaybook interface {
	Run() error
	DryRun() error
}

type playbook struct {
	Stage Stage
}

func (p *playbook) Run() error {
	cmd := p.Stage.Cmd()
	log.WithFields(log.Fields{
		"stage": p.Stage.String(),
		"cmd":   cmd.Args,
		"env":   cmd.Env,
	}).Debug()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p *playbook) DryRun() error {
	cmd := p.Stage.Cmd()
	cmd.Args = append(cmd.Args, DryRunFlag)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// NewBootstrapPlaybook - returns a bootstrap playbook implementing AnsiblePlaybook
func NewBootstrapPlaybook() AnsiblePlaybook {
	return &playbook{Stage: Bootstrap}
}

// NewChrootPlaybook - returns bootstrap playbook implementing Playbook interface
func NewChrootPlaybook() AnsiblePlaybook {
	return &playbook{Stage: Chroot}
}

// NewMasterPlaybook - returns master playbook implementing Playbook interface
func NewMasterPlaybook() AnsiblePlaybook {
	return &playbook{Stage: Master}
}
