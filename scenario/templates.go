package scenario

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/sugarraysam/archsugar-cli/helpers"
)

const (
	VarsBasename  = "vars.tmpl"
	TasksBasename = "tasks.tmpl"
)

var (
	VarsBasedir      = path.Join(helpers.BaseDir, "roles/main/vars/master")
	TasksBasedir     = path.Join(helpers.BaseDir, "roles/main/tasks/master")
	TemplatesBasedir = path.Join(helpers.BaseDir, "scenario/templates")
)

// AnsibleTemplate - represent a template to easily create scenarios
type AnsibleTemplate interface {
	Write() error
	Rm() error
	Exists() bool
	GetDst() string
}

type tmpl struct {
	Name string
	Src  string
	Dst  string
	Data map[string]string
}

// Write - write Ansible vars file from template
func (t *tmpl) Write() error {
	// change delims because ansible uses "{{", "}}"
	gotmpl, err := template.New(t.Name).Delims("((", "))").ParseFiles(t.Src)
	if err != nil {
		return err
	}
	outFile, err := os.Create(t.Dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return gotmpl.ExecuteTemplate(outFile, t.Name, t.Data)
}

// Rm - remove ansible vars file
func (t *tmpl) Rm() error {
	return os.Remove(t.Dst)
}

// Exist - check if destination already exists
func (t *tmpl) Exists() bool {
	_, err := os.Stat(t.Dst)
	return err == nil
}

// GetDst - returns template destination
func (t *tmpl) GetDst() string {
	return t.Dst
}

// NewVarsTemplate - used to create an Ansible variable file for a scenario from a go template
func NewVarsTemplate(name string) AnsibleTemplate {
	return &tmpl{
		Name: VarsBasename,
		Src:  path.Join(TemplatesBasedir, VarsBasename),
		Dst:  path.Join(VarsBasedir, fmt.Sprintf("%s.yml", name)),
		Data: map[string]string{
			"Name": name,
		},
	}
}

// NewTasksTemplate - used to create an Ansible tasks file for a scenario from a go template
func NewTasksTemplate(name, desc string) AnsibleTemplate {
	return &tmpl{
		Name: TasksBasename,
		Src:  path.Join(TemplatesBasedir, TasksBasename),
		Dst:  path.Join(TasksBasedir, fmt.Sprintf("%s.yml", name)),
		Data: map[string]string{
			"Name": name,
			"Desc": desc,
		},
	}
}
