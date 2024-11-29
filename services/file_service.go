package services

import (
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name     string
	IsDir    bool
	SubFiles []FileInfo // For nested files and directories
}

// ListFiles returns a list of files in the specified directory
func ListFiles(directory string) ([]FileInfo, error) {
	var files []FileInfo

	// Read the directory contents
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip the base directory itself
		if path != directory {
			// Determine whether the entry is a file or directory
			file := FileInfo{
				Name:  info.Name(),
				IsDir: info.IsDir(),
			}
			if info.IsDir() {
				// If it's a directory, find its contents recursively
				subFiles, err := ListFiles(path)
				if err != nil {
					return err
				}
				file.SubFiles = subFiles
			}
			files = append(files, file)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
