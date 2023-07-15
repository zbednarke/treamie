package package1

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileInfo represents information about a file or directory.
type FileInfo struct {
	Name     string
	Size     int64
	IsDir    bool
	Depth    int
	Parent   string
	Children []FileInfo
}

// TraverseDirectory traverses the directory structure and gathers information about files and directories.
func TraverseDirectory(rootPath string, ignoreList []string) ([]FileInfo, error) {
	rootDirName := filepath.Base(rootPath)
	fileInfos := []FileInfo{{
		Name:     rootDirName,
		IsDir:    true,
		Depth:    0,
		Children: []FileInfo{},
	},
	}

	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			log.Println(redString("Error getting relative path:" + err.Error()))
			return nil
		}

		fname := info.Name()
		if filepath.Dir(relPath) != "." {
			fname = filepath.Dir(relPath) + "/" + info.Name()
		}
		fileInfo := FileInfo{
			Name:     fname,
			IsDir:    info.IsDir(),
			Parent:   filepath.Dir(relPath),
			Children: []FileInfo{},
		}

		// log.Println(fileInfo)

		// Calculate the depth based on the number of directory separators
		fileInfo.Depth = strings.Count(relPath, string(os.PathSeparator))

		// If the file or directory is in the ignore list, skip it
		if isIgnored(relPath, ignoreList) {
			if fileInfo.IsDir {
				return filepath.SkipDir // Skip the entire directory
			}
			return nil // Skip the file
		}

		// Add the file info to its parent directory
		if fileInfo.Parent == "." {
			fileInfo.Parent = rootDirName
		}
		parentDir := getParentDirectory(fileInfos, fileInfo)
		if parentDir != nil {
			parentDir.Children = append(parentDir.Children, fileInfo)
		} else {
			log.Println(redString("Parent directory not found for file: " + fileInfo.Name))
		}

		if fileInfo.IsDir {
			fileInfos = append(fileInfos, fileInfo)
			return nil
		} else {
			// Calculate file size for non-directory files
			fsFileInfo, err := info.Info()
			if err != nil {
				log.Println(redString("Error getting file info:" + err.Error()))
				return nil
			}
			fileInfo.Size = fsFileInfo.Size()
			log.Println("size: " + strconv.FormatInt(fileInfo.Size, 10))
			return nil
		}
	})

	if err != nil {
		return nil, fmt.Errorf("error traversing directory: %v", err)
	}

	return fileInfos, nil
}

func redString(str string) string {
	return "\033[31m" + str + "\033[0m"
}

func greenString(str string) string {
	return "\033[32m" + str + "\033[0m"
}

// getParentDirectory finds the parent directory in the file info slice based on the parent path.
func getParentDirectory(fileInfos []FileInfo, fileInfo FileInfo) *FileInfo {
	parentPath := fileInfo.Parent

	log.Println(greenString(fileInfo.Name + " PPath: " + parentPath + " isDir: " + strconv.FormatBool(fileInfo.IsDir)))

	for i := range fileInfos {
		if fileInfos[i].IsDir && fileInfos[i].Name == parentPath {
			return &fileInfos[i]
		}
	}
	return nil
}

// isIgnored checks if a file or directory should be ignored based on the ignore list.
func isIgnored(path string, ignoreList []string) bool {
	for _, ignore := range ignoreList {
		if strings.Contains(path, ignore) {
			return true
		}
	}
	return false
}
