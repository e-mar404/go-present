package gopresent

import (
	"os"
)

func getMdFilesFrom(path string) ([]os.DirEntry, error) {
	entries, err :=	os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry)
		}
	}
	return files, nil
}

// todo: getFilesBy... then some sort of filter
