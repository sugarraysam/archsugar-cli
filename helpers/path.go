package helpers

import (
	"path"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

const ArchsugarDefaultPath = ".archsugar"

var (
	BaseDir string = func() string {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("homdir.Dir(), got %v", err)
		}
		return path.Join(home, ArchsugarDefaultPath)
	}()
)
