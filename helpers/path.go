package helpers

import (
	"fmt"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var (
	BaseDir string = func() string {
		_, filename, linenum, ok := runtime.Caller(0)
		if !ok {
			log.Fatal(nil, fmt.Sprintf("runtime.Caller(0) failed: %v : %v", filename, linenum))
		}
		return path.Join(path.Dir(filename), "..")
	}()
)
