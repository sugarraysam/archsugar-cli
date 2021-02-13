package ansible

import (
	"fmt"
)

var (
	DefaultArgs = []string{"-e target=localhost"}
	ChrootArgs  = []string{"-c", "chroot", "-e target=all", "-i", "/mnt,"}
	DefaultEnv  = []string{
		"ANSIBLE_FORCE_COLOR=true",
		"ANSIBLE_NOCOWS=1",
	}
	ChrootExtraEnv = []string{
		"ANSIBLE_CHROOT_EXE=/usr/bin/arch-chroot",
		"ANSIBLE_EXECUTABLE=/usr/bin/bash",
	}
)

// Stage - type representing each of the three archsugar stages
type Stage int

const (
	BootstrapStage Stage = iota
	ChrootStage
	MasterStage
)

func (s Stage) String() string {
	return []string{"bootstrap", "chroot", "master"}[s]
}

func (s Stage) DefaultEnv() []string {
	res := DefaultEnv
	res = append(res, fmt.Sprintf("SUGAR_STAGE=%s", s.String()))
	if s == ChrootStage {
		res = append(res, ChrootExtraEnv...)
	}
	return res
}

func (s Stage) DefaultArgs() []string {
	if s == ChrootStage {
		return ChrootArgs
	}
	return DefaultArgs
}
