package main

import (
	"fmt"
	"github.com/tlarsen7572/Golang-Public/helpers"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"os"
	"path/filepath"
	"strings"
)

var validExts = helpers.StringArray{`.yxmd`, `.yxmc`, `.yxwz`}

func cleanFolder(folder string) {
	_ = filepath.Walk(folder, _cleanFile)
}

func _cleanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return _print(err)
	}

	ext := strings.ToLower(filepath.Ext(info.Name()))
	if !validExts.Contains(ext) {
		return nil
	}

	doc, err := ryxdoc.ReadFile(path)
	if err != nil {
		return _print(err)
	}

	err = doc.Save(path)
	if err != nil {
		return _print(err)
	}
	return nil
}

func _print(err error) error {
	print(fmt.Sprintf(`error cleaning folder: %v`, err.Error()))
	return err
}
