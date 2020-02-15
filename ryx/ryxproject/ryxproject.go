package ryxproject

import (
	"errors"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxfolder"
	"os"
	"path/filepath"
	"strings"
)

type RyxProject struct {
	path       string
	macroPaths []string
}

func Open(path string, macroPaths ...string) (*RyxProject, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New(`cannot open a file; only directories can be opened`)
	}
	return &RyxProject{path: absPath, macroPaths: macroPaths}, nil
}

func (ryxProject *RyxProject) Structure() (*ryxfolder.RyxFolder, error) {
	return ryxfolder.Build(ryxProject.path)
}

func (ryxProject *RyxProject) Docs() (map[string]*ryxdoc.RyxDoc, error) {
	structure, err := ryxProject.Structure()
	if err != nil {
		return nil, err
	}
	docs := docsFromStructure(structure)
	return docs, nil
}

func (ryxProject *RyxProject) ReadPath() string {
	return ryxProject.path
}

func (ryxProject *RyxProject) RenameFile(oldPath string, newPath string) error {
	docs, err := ryxProject.Docs()
	if err != nil {
		return err
	}
	for path, doc := range docs {
		if path == newPath {
			continue
		}
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		renamed := doc.RenameMacroNodes(oldPath, newPath, macroPaths...)
		if renamed > 0 {
			doc.Save(path)
		}
	}
	return os.Rename(oldPath, newPath)
}

func (ryxProject *RyxProject) MakeAllMacrosAbsolute() int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		changed := doc.MakeAllMacrosAbsolute(macroPaths...)
		if changed > 0 {
			docsChanged++
			doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeMacroAbsolute(macroAbsPath string) int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		changed := doc.MakeMacroAbsolute(macroAbsPath, macroPaths...)
		if changed > 0 {
			docsChanged++
			doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeAllMacrosRelative() int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		changed := doc.MakeAllMacrosRelative(folder, macroPaths...)
		if changed > 0 {
			docsChanged++
			doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeMacroRelative(macroAbsPath string) int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		changed := doc.MakeMacroRelative(macroAbsPath, folder, macroPaths...)
		if changed > 0 {
			docsChanged++
			doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) RetrieveDocument(path string) (*ryxdoc.RyxDoc, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	rel, err := filepath.Rel(ryxProject.ReadPath(), absPath)
	if err != nil {
		return nil, err
	}
	if strings.Contains(rel, filepath.Join(`..`, ``)) {
		return nil, errors.New(`path is not a child of the project directory`)
	}
	return ryxdoc.ReadFile(absPath)
}

func (ryxProject *RyxProject) WhereUsed(path string) []string {
	usage := []string{}
	docs, err := ryxProject.Docs()
	if err != nil {
		return usage
	}
	for docPath, doc := range docs {
		folder := filepath.Dir(docPath)
		macroPaths := ryxProject.generateMacroPaths(folder)
		for _, node := range doc.ReadMappedNodes() {
			macro := node.ReadMacro(macroPaths...)
			if macro.FoundPath == path {
				usage = append(usage, docPath)
				break
			}
		}
	}
	return usage
}

func docsFromStructure(structure *ryxfolder.RyxFolder) map[string]*ryxdoc.RyxDoc {
	docs := map[string]*ryxdoc.RyxDoc{}
	for _, file := range structure.AllFiles() {
		doc, err := ryxdoc.ReadFile(file)
		if err == nil {
			docs[file] = doc
		}
	}
	return docs
}

func (ryxProject *RyxProject) generateMacroPaths(additionalPaths ...string) []string {
	return append(additionalPaths, ryxProject.macroPaths...)
}
