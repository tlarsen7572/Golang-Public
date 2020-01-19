package folders

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func InitializeFolderController(visiblePaths ...string) *FolderController {
	return &FolderController{visiblePaths: visiblePaths}
}

type FolderController struct {
	visiblePaths []string
}

func (controller *FolderController) ReadFolder(folder string) ([]string, error) {
	if folder == `` {
		return controller.visiblePaths, nil
	}
	if !filepath.IsAbs(folder) {
		return nil, errors.New(`only absolute paths are supported`)
	}
	if !controller.IsFolderVisible(folder) {
		return nil, errors.New(fmt.Sprintf(`'%v' does not exist or is not visible to the application`, folder))
	}
	return readFolderContent(folder)
}

func (controller *FolderController) IsFolderVisible(folder string) bool {
	isVisible := false
	for _, visiblePath := range controller.visiblePaths {
		rel, err := filepath.Rel(folder, visiblePath)
		if err != nil {
			continue
		}
		if rel == `.` {
			isVisible = true
			break
		}
		if len(rel) > 1 && rel[0:2] == `..` {
			isVisible = true
			break
		}
	}
	return isVisible
}

func readFolderContent(folder string) ([]string, error) {
	items, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	folderContent := make([]string, 0)
	for _, item := range items {
		if item.IsDir() {
			folderContent = append(folderContent, filepath.Join(folder, item.Name()))
		}
	}
	return folderContent, nil
}
