package main

import (
	"fmt"
	"github.com/tlarsen7572/Golang-Public/helpers"
	"github.com/tlarsen7572/Golang-Public/txml"
	"io/ioutil"
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

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return _print(err)
	}

	xmlNode, err := txml.Parse(string(content))
	if err != nil {
		return _print(err)
	}

	_cleanDoc(xmlNode)
	xmlStr, err := xmlNode.ToXml(`  `)
	if err != nil {
		return _print(err)
	}
	xmlBytes := []byte(xmlStr)
	err = ioutil.WriteFile(path, xmlBytes, os.ModePerm)
	if err != nil {
		return _print(err)
	}
	return nil
}

func _cleanDoc(xml *txml.Node) {
	nodes := xml.First(`Nodes`)
	_cleanNodes(nodes)
}

func _cleanNodes(nodes *txml.Node) {
	for _, node := range nodes.Nodes {
		properties := node.First(`Properties`)
		properties.RemoveAll(`MetaInfo`)
		childNodes := node.First(`ChildNodes`)
		if len(childNodes.Nodes) > 0 {
			_cleanNodes(childNodes)
		}
	}
}

func _print(err error) error {
	print(fmt.Sprintf(`error cleaning folder: %v`, err.Error()))
	return err
}
