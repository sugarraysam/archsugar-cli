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
	EnabledBasedir          = path.Join(TasksBasedir, "enabled")
	ErrScenarioDoesNotExist = errors.New("scenario does not exist")
)

type CreatedScenario interface {
	Rm() error
	Enable() error
	Disable() error
	GetName() string
	GetDesc() string
	IsEnabled() bool
}

type cscenario struct {
	Name           string
	Desc           string
	Templates      map[string]AnsibleTemplate
	EnabledSymlink string
	Enabled        bool
}

// NewCreatedScenario - returns a CreatedScenario and validates if it is created
func NewCreatedScenario(name string) (CreatedScenario, error) {
	s := &cscenario{
		Name: name,
		Desc: "",
		Templates: map[string]AnsibleTemplate{
			"vars":  NewVarsTemplate(name),
			"tasks": NewTasksTemplate(name, ""),
		},
		EnabledSymlink: path.Join(EnabledBasedir, fmt.Sprintf("%s.yml", name)),
		Enabled:        false,
	}
	for _, t := range s.Templates {
		if !t.Exists() {
			return nil, ErrScenarioDoesNotExist
		}
	}
	// reads in the description from the existing ansible task file
	s.setDesc()
	if _, err := os.Lstat(s.EnabledSymlink); err == nil {
		s.Enabled = true
	}
	return s, nil
}

func (s *cscenario) setDesc() {
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

// Rm - deletes a CreatedScenario
func (s *cscenario) Rm() error {
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
func (s *cscenario) Enable() error {
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
func (s *cscenario) Disable() error {
	if !s.Enabled {
		return nil
	}
	if err := os.Remove(s.EnabledSymlink); err != nil {
		return err
	}
	s.Enabled = false
	return nil
}

// GetName - return name of scenario
func (s *cscenario) GetName() string {
	return s.Name
}

// GetDesc - return description of scenario
func (s *cscenario) GetDesc() string {
	return s.Desc
}

// IsEnabled - returns wheter scenario is enabled or not
func (s *cscenario) IsEnabled() bool {
	return s.Enabled
}
