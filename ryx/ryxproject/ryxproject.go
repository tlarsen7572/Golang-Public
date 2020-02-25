package ryxproject

import (
	"errors"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxdoc"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxfolder"
	"github.com/tlarsen7572/Golang-Public/ryx/ryxnode"
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

func (ryxProject *RyxProject) RenameFolder(from string, to string) error {
	parent := filepath.Dir(from)
	toPath := filepath.Join(parent, to)
	oldPaths := make([]string, 0)
	newPaths := make([]string, 0)
	folder, err := ryxfolder.Build(from)
	if err != nil {
		return err
	}
	for _, file := range folder.AllFiles() {
		trimmed := strings.TrimPrefix(file, from)
		newPath := filepath.Join(toPath, trimmed)
		oldPaths = append(oldPaths, file)
		newPaths = append(newPaths, newPath)
	}
	organizer, err := ryxProject._collectAffectedNodes(oldPaths, newPaths)
	if err != nil {
		return err
	}

	err = os.Rename(from, toPath)
	if err != nil {
		return err
	}

	for _, tracker := range organizer.trackers {
		for _, node := range tracker.nodes {
			node.SetMacro(tracker.newPath)
		}
	}
	for index, oldPath := range oldPaths {
		if doc, ok := organizer.affectedDocs[oldPath]; ok {
			organizer.affectedDocs[newPaths[index]] = doc
			delete(organizer.affectedDocs, oldPath)
		}
	}

	for path, doc := range organizer.affectedDocs {
		_ = doc.Save(path)
	}
	return nil
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

	changeOrganizer, err := ryxProject._collectAffectedNodes(oldPaths, newPaths)
	if err != nil {
		return nil, err
	}

	//Save old files to new location
	for index := range oldPaths {
		oldPath := oldPaths[index]
		newPath := newPaths[index]
		doc, ok := changeOrganizer.allDocs[oldPath]
		if !ok {
			oldPathsFailed = append(oldPathsFailed, oldPath)
			delete(changeOrganizer.trackers, oldPath)
			continue
		}
		macroPaths := ryxProject.generateMacroPaths(filepath.Dir(oldPath))
		doc.MakeAllMacrosAbsolute(macroPaths...)
		renameErr := doc.Save(newPath)
		if renameErr != nil {
			oldPathsFailed = append(oldPathsFailed, oldPath)
			delete(changeOrganizer.trackers, oldPath)
			continue
		}
		if _, ok := changeOrganizer.affectedDocs[oldPath]; ok {
			changeOrganizer.affectedDocs[newPath] = doc
		}
		changeOrganizer.allDocs[newPath] = doc
		delete(changeOrganizer.affectedDocs, oldPath)
		delete(changeOrganizer.allDocs, oldPath)
		oldPathsSuccess = append(oldPathsSuccess, oldPath)
	}

	//Redirect paths where the renamed files are used
	for _, tracker := range changeOrganizer.trackers {
		for _, node := range tracker.nodes {
			node.SetMacro(tracker.newPath)
		}
	}
	for path, doc := range changeOrganizer.affectedDocs {
		_ = doc.Save(path)
	}

	//Delete old files
	for _, path := range oldPathsSuccess {
		_ = os.Remove(path)
	}

	return oldPathsFailed, nil
}

type RenameOrganizer struct {
	trackers     map[string]*RenameTracker
	affectedDocs map[string]*ryxdoc.RyxDoc
	allDocs      map[string]*ryxdoc.RyxDoc
}

type RenameTracker struct {
	path    string
	newPath string
	nodes   []*ryxnode.RyxNode
}

func (ryxProject *RyxProject) _collectAffectedNodes(oldPaths []string, newPaths []string) (*RenameOrganizer, error) {
	docs, err := ryxProject.Docs()
	if err != nil {
		return nil, err
	}
	organizer := &RenameOrganizer{
		trackers:     make(map[string]*RenameTracker, len(oldPaths)),
		affectedDocs: make(map[string]*ryxdoc.RyxDoc),
		allDocs:      docs,
	}
	for index := range oldPaths {
		organizer.trackers[oldPaths[index]] = &RenameTracker{
			path:    oldPaths[index],
			newPath: newPaths[index],
			nodes:   nil,
		}
	}
	for path, doc := range docs {
		folder := filepath.Dir(path)
		affectedMacros := 0
		for _, node := range doc.Nodes {
			if node.ReadCategory() != ryxnode.Macro {
				continue
			}
			macroPaths := append(ryxProject.macroPaths, folder)
			macro := node.ReadMacro(macroPaths...)
			for _, oldPath := range oldPaths {
				if macro.FoundPath == oldPath {
					organizer.trackers[oldPath].nodes = append(organizer.trackers[oldPath].nodes, node)
					affectedMacros++
					continue
				}
			}
		}
		if affectedMacros > 0 {
			organizer.affectedDocs[path] = doc
		}
	}
	return organizer, nil
}

func StringsContain(strings []string, value string) bool {
	for _, item := range strings {
		if item == value {
			return true
		}
	}
	return false
}
