package scenario

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrScenarioAlreadyExists = errors.New("scenario is already created")
	ErrScenarioDoesNotExist  = errors.New("scenario does not exist")
)

type Scenario struct {
	Name           string
	Desc           string
	Templates      map[string]AnsibleTemplate
	EnabledSymlink string
	Enabled        bool
}

// New - writes scenario templates to filesystem
func New(baseDir, name, desc string) (*Scenario, error) {
	s := base(baseDir, name, desc)
	for _, t := range s.Templates {
		if t.Exists() {
			return nil, ErrScenarioAlreadyExists
		}
		if err := t.Write(); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func base(baseDir, name, desc string) *Scenario {
	return &Scenario{
		Name: name,
		Desc: desc,
		Templates: map[string]AnsibleTemplate{
			"vars":  NewVarsTemplate(baseDir, name),
			"tasks": NewTasksTemplate(baseDir, name, desc),
		},
		EnabledSymlink: path.Join(EnabledDir(baseDir), fmt.Sprintf("%s.yml", name)),
		Enabled:        false,
	}
}

// Read - initialize a scenario struct from filesystem
func Read(baseDir, name string) (*Scenario, error) {
	s := base(baseDir, name, "")
	for _, t := range s.Templates {
		if !t.Exists() {
			return nil, ErrScenarioDoesNotExist
		}
	}
	s.readDesc()
	s.setEnabled()
	return s, nil
}

func (s *Scenario) readDesc() {
	taskFile, err := os.Open(s.Templates["tasks"].GetDst())
	if err != nil {
		return
	}
	defer taskFile.Close()

	desc, err := bufio.NewReader(taskFile).ReadString('\n')
	if err != nil || !strings.HasPrefix(desc, "#") {
		return
	}
	s.Desc = strings.Trim(desc, "# \n")
}

func (s *Scenario) setEnabled() {
	if _, err := os.Lstat(s.EnabledSymlink); err == nil {
		s.Enabled = true
	}
}

// Rm - deletes a CreatedScenario
func (s *Scenario) Rm() error {
	if s.Enabled {
		if err := s.Disable(); err != nil {
			return err
		}
	}
	for _, t := range s.Templates {
		if err := t.Rm(); err != nil {
			return err
		}
	}
	return nil
}

// Enable - enables the scenario by writing the EnabledSymlink
func (s *Scenario) Enable() error {
	if s.Enabled {
		return nil
	}
	err := os.Symlink(s.Templates["tasks"].GetDst(), s.EnabledSymlink)
	if err != nil {
		return err
	}
	s.Enabled = true
	return nil
}

// Disable - disables the scenario by removing the EnabledSymlink
func (s *Scenario) Disable() error {
	if !s.Enabled {
		return nil
	}
	if err := os.Remove(s.EnabledSymlink); err != nil {
		return err
	}
	s.Enabled = false
	return nil
}

func (s *Scenario) IsEnabled() bool {
	return s.Enabled
}

func (s *Scenario) Exists() bool {
	for _, t := range s.Templates {
		if !t.Exists() {
			return false
		}
	}
	return true
}
