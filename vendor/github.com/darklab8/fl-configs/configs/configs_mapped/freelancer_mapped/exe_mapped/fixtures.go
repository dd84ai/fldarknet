package exe_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/tests"
)

func FixtureFLINIConfig() *Config {
	fileref := tests.FixtureFileFind().GetFile(FILENAME_FL_INI)
	config := Read(iniload.NewLoader(fileref).Scan())
	return config
}
