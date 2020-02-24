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

func (ryxProject *RyxProject) RenameFiles(fromFiles []string, toFiles []string) ([]string, error) {
	return ryxProject._renameFiles(fromFiles, toFiles)
}

func (ryxProject *RyxProject) MoveFiles(files []string, moveTo string) ([]string, error) {
	newFiles := []string{}
	for _, file := range files {
		_, name := filepath.Split(file)
		newPath := filepath.Join(moveTo, name)
		newFiles = append(newFiles, newPath)
	}
	return ryxProject._renameFiles(files, newFiles)
}

func (ryxProject *RyxProject) MakeAllFilesAbsolute() int {
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
			_ = doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeFilesAbsolute(macroAbsPath []string) int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		var changed int
		if StringsContain(macroAbsPath, path) {
			changed = doc.MakeAllMacrosAbsolute(macroPaths...)
		} else {
			changed = doc.MakeMacrosAbsolute(macroAbsPath, macroPaths...)
		}
		if changed > 0 {
			docsChanged++
			_ = doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeAllFilesRelative() int {
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
			_ = doc.Save(path)
		}
	}
	return docsChanged
}

func (ryxProject *RyxProject) MakeFilesRelative(macroAbsPath []string) int {
	docs, err := ryxProject.Docs()
	if err != nil {
		return 0
	}
	docsChanged := 0
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		var changed int
		if StringsContain(macroAbsPath, path) {
			changed = doc.MakeAllMacrosRelative(folder, macroPaths...)
		} else {
			changed = doc.MakeMacrosRelative(macroAbsPath, folder, macroPaths...)
		}
		if changed > 0 {
			docsChanged++
			_ = doc.Save(path)
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

func (ryxProject *RyxProject) _renameFiles(oldPaths []string, newPaths []string) ([]string, error) {
	if len(oldPaths) != len(newPaths) {
		return nil, errors.New(`the lists of From and To files were not the same length`)
	}

	oldPathsFailed := []string{}
	oldPathsSuccess := []string{}
	newPathsSuccess := []string{}

	//Save old files to new location
	for index := range oldPaths {
		oldPath := oldPaths[index]
		newPath := newPaths[index]
		doc, err := ryxdoc.ReadFile(oldPath)
		if err != nil {
			oldPathsFailed = append(oldPathsFailed, oldPath)
			continue
		}
		macroPaths := ryxProject.generateMacroPaths(filepath.Dir(oldPath))
		doc.MakeAllMacrosAbsolute(macroPaths...)
		renameErr := doc.Save(newPath)
		if renameErr != nil {
			oldPathsFailed = append(oldPathsFailed, oldPath)
			continue
		}
		oldPathsSuccess = append(oldPathsSuccess, oldPath)
		newPathsSuccess = append(newPathsSuccess, newPath)
	}

	//Get docs in the project
	docs, err := ryxProject.Docs()
	if err != nil {
		for _, path := range newPathsSuccess {
			_ = os.Remove(path)
		}
		return nil, err
	}

	//Fix macro paths
	for path, doc := range docs {
		folder := filepath.Dir(path)
		macroPaths := ryxProject.generateMacroPaths(folder)
		renamed := doc.RenameMacroNodes(oldPathsSuccess, newPathsSuccess, macroPaths...)
		if renamed > 0 {
			_ = doc.Save(path)
		}
	}

	//Delete old files
	for _, path := range oldPathsSuccess {
		_ = os.Remove(path)
	}

	return oldPathsFailed, nil
}

func StringsContain(strings []string, value string) bool {
	for _, item := range strings {
		if item == value {
			return true
		}
	}
	return false
}
