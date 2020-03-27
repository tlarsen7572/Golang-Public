package config

import (
	"encoding/json"
	"github.com/tlarsen7572/Golang-Public/ryx/ini_reader"
	"github.com/tlarsen7572/Golang-Public/ryx/tool_data_loader"
	"io/ioutil"
	"path/filepath"
)

func LoadConfig() (*Config, error) {
	content, err := ioutil.ReadFile(`config.json`)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	InstallPath       string
	ProgramDataPath   string
	UserFolders       []string
	Address           string
	BrowseFolderRoots []string
	LogPath           string
	ToolData          []tool_data_loader.ToolData
}

func (config *Config) MacroPaths() []string {
	macroPaths := []string{}
	iniFolder := filepath.Join(config.ProgramDataPath, `DataProducts`, `AddOnData`, `Macros`)
	iniFiles, _ := ioutil.ReadDir(iniFolder)
	for _, fileInfo := range iniFiles {
		if filepath.Ext(fileInfo.Name()) != `.ini` {
			continue
		}
		mapped := ini_reader.LoadIni(filepath.Join(iniFolder, fileInfo.Name()))
		folder, ok := mapped[`Path`]
		if ok {
			macroPaths = append(macroPaths, folder)
		}
	}
	return macroPaths
}
