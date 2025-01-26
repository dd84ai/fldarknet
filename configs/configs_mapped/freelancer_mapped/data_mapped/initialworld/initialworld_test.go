package initialworld

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-darkstat/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-darkstat/configs/tests"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	fileref := tests.FixtureFileFind().GetFile(FILENAME)

	config := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(config.LockedGates), 0, "expected finding some elements")
	fmt.Println(config.LockedGates)
}
