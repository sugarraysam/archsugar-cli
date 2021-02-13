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
	Name() string
}

type playbook struct {
	Builder *Builder
}

func (p *playbook) Run() error {
	cmd := p.Builder.Cmd()
	log.WithFields(log.Fields{
		"stage": p.Builder.Stage.String(),
		"cmd":   cmd.Args,
		"env":   cmd.Env,
	}).Debug()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p *playbook) DryRun() error {
	cmd := p.Builder.Cmd()
	cmd.Args = append(cmd.Args, DryRunFlag)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p *playbook) Name() string {
	return p.Builder.Stage.String()
}

// NewBootstrap - returns a bootstrap playbook implementing AnsiblePlaybook
func NewBootstrap(basePath string) AnsiblePlaybook {
	return &playbook{Builder: NewBuilder(Bootstrap, basePath)}
}

// NewChroot - returns bootstrap playbook implementing Playbook interface
func NewChroot(basePath string) AnsiblePlaybook {
	return &playbook{Builder: NewBuilder(Chroot, basePath)}
}

// NewMaster - returns master playbook implementing Playbook interface
func NewMaster(basePath string) AnsiblePlaybook {
	return &playbook{Builder: NewBuilder(Master, basePath)}
}
