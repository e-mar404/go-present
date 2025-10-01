package main

import (
	"os"
	"path/filepath"
	"strings"
)

func mdFilesFrom(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []os.DirEntry
	for _, entry := range entries {
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if !entry.IsDir() && ext == ".md" {
			files = append(files, entry)
		}
	}
	return files, nil
}

// todo: getFilesBy ... then some sort of filter
