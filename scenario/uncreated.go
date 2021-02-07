package scenario

import (
	"github.com/pkg/errors"
)

var (
	ErrScenarioAlreadyExists = errors.New("scenario is already created")
)

// UncreatedScenario - interface representing functionality of a new scenario pending creation
type UncreatedScenario interface {
	Create() error
}

type uscenario struct {
	Name      string
	Desc      string
	Templates []AnsibleTemplate
}

// NewUncreatedScenario - returns an UncreatedScenario and validates it does not yet exist
func NewUncreatedScenario(name, desc string) (UncreatedScenario, error) {
	s := &uscenario{
		Name: name,
		Desc: desc,
		Templates: []AnsibleTemplate{
			NewVarsTemplate(name),
			NewTasksTemplate(name, desc),
		},
	}
	for _, t := range s.Templates {
		if t.Exists() {
			return nil, ErrScenarioAlreadyExists
		}
	}
	return s, nil
}

// Create - writes an UncreatedScenario's templates to the filesystem
func (s *uscenario) Create() error {
	for _, t := range s.Templates {
		if err := t.Write(); err != nil {
			return err
		}
	}
	return nil
}
