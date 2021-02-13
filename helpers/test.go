package helpers

import (
	"fmt"
	"math/rand"
	"os"
	"path"
)

const (
	MinDigit = 10000
	MaxDigit = 20000
)

// TmpDir - generate a tmpDir, avoiding very rare collisions
func TmpDir() string {
	for {
		res := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", RandomDigit()))
		_, err := os.Stat(res)
		if err != nil {
			return res
		}
	}
}

func RandomDigit() int {
	return rand.Intn(MinDigit) + MaxDigit - MinDigit
}
