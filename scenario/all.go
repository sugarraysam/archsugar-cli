package scenario

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	MaxLineLen    = 100
	EnabledToken  = "   X   "
	DisabledToken = " "
	MainTaskFile  = "main.yml"
)

var (
	Header    = FormatLine("NAME", "ENABLED", "DESCRIPTION")
	Separator = strings.Repeat("-", MaxLineLen)
)

type AllScenarios interface {
	List(w io.Writer) error
	DisableAll() error
	EnableAll() error
	GetScenarios() []CreatedScenario
}

type scenarios []CreatedScenario

// GetAllScenarios - returns all created scenarios, sorted by name
func GetAllScenarios() (AllScenarios, error) {
	var xs scenarios
	files, err := ioutil.ReadDir(TasksBasedir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() || f.Name() == MainTaskFile {
			continue
		}
		// no concurrency to keep scenarios in sorted order (from ioutil.ReadDir)
		// also, plus there should not be more than 30 scenarios at any point
		name := strings.TrimSuffix(f.Name(), ".yml")
		scenario, err := NewCreatedScenario(name)
		if err != nil {
			return nil, err
		}
		xs = append(xs, scenario)
	}
	return xs, nil
}

// List - pretty print created scenarios to writer w
func (xs scenarios) List(w io.Writer) error {
	if err := printHeaderAndSeparator(w); err != nil {
		return err
	}
	for _, s := range xs {
		line := FormatLine(s.GetName(), DisabledToken, s.GetDesc())
		if s.IsEnabled() {
			line = FormatLine(s.GetName(), EnabledToken, s.GetDesc())
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func printHeaderAndSeparator(w io.Writer) error {
	for _, line := range []string{Header, Separator} {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func FormatLine(name, enabled, desc string) string {
	nameLen := 20
	enabledLen := len("ENABLED")
	format := "%-*s | %*s | %-s"
	line := fmt.Sprintf(format, nameLen, name, enabledLen, enabled, desc)
	if len(line) > MaxLineLen {
		return shrink(line)
	}
	return line
}

func shrink(line string) string {
	shrinkChars := "..."
	return strings.Join([]string{
		line[:MaxLineLen-len(shrinkChars)],
		shrinkChars,
	}, "")
}

// Disable - disable all master stage scenarios
func (xs scenarios) DisableAll() error {
	for _, s := range xs {
		if err := s.Disable(); err != nil {
			return err
		}
	}
	return nil
}

// EnableAll - enable all master scenarios
func (xs scenarios) EnableAll() error {
	for _, s := range xs {
		if err := s.Enable(); err != nil {
			return err
		}
	}
	return nil
}

// GetScenarios - returns underlying scenarios
func (xs scenarios) GetScenarios() []CreatedScenario {
	return xs
}
