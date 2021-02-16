package ansible

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	DryRunFlag = "--list-tasks"
)

type Playbook struct {
	Builder *Builder
}

func (p *Playbook) Run() error {
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

func (p *Playbook) DryRun() error {
	cmd := p.Builder.Cmd()
	cmd.Args = append(cmd.Args, DryRunFlag)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p *Playbook) Name() string {
	return p.Builder.Stage.String()
}

// NewBootstrap - returns a bootstrap playbook implementing AnsiblePlaybook
func NewBootstrapPlaybook(baseDir string) *Playbook {
	return &Playbook{Builder: NewBuilder(baseDir, BootstrapStage)}
}

// NewChroot - returns bootstrap playbook implementing Playbook interface
func NewChrootPlaybook(baseDir string) *Playbook {
	return &Playbook{Builder: NewBuilder(baseDir, ChrootStage)}
}

// NewMaster - returns master playbook implementing Playbook interface
func NewMasterPlaybook(baseDir string) *Playbook {
	return &Playbook{Builder: NewBuilder(baseDir, MasterStage)}
}
