package config_test

import (
	c "github.com/tlarsen7572/Golang-Public/ryx/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := c.LoadConfig()
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if config.InstallPath != `C:\Program Files\Alteryx` {
		t.Fatalf(`expected install path 'C:\Program Files\Alteryx' but got '%v'`, config.InstallPath)
	}
	if config.ProgramDataPath != `C:\ProgramData\Alteryx` {
		t.Fatalf(`expected programdata path 'C:\ProgramData\Alteryx' but got '%v'`, config.ProgramDataPath)
	}
	if len(config.UserFolders) != 1 {
		t.Fatalf(`expected 1 user folder but got %v`, len(config.UserFolders))
	}
	if config.UserFolders[0] != `C:\Users\tlarsen` {
		t.Fatalf(`expected user folder 'tlarsen' but got '%v'`, config.UserFolders[0])
	}
	if config.Address != `localhost:35012` {
		t.Fatalf(`expected address 'localhost:35012' but got '%v'`, config.Address)
	}
	if count := len(config.BrowseFolderRoots); count != 2 {
		t.Fatalf(`expected 2 browse folder roots but got %v`, count)
	}
	if config.LogPath != `.\log.txt` {
		t.Fatalf(`expected install path '.\log.txt' but got '%v'`, config.InstallPath)
	}
}

func TestMacroPaths(t *testing.T) {
	config, _ := c.LoadConfig()
	macroPaths := config.MacroPaths()
	if count := len(macroPaths); count == 0 {
		t.Fatalf(`expected at least 1 macro path but got 0`)
	}
}
