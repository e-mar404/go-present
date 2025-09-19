package gopresent

import (
	"errors"
	"os"
	"testing"
)

func fileNames(d []os.DirEntry) []string {
	fileNames := make([]string, len(d))
	for _, f := range d {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func TestGetMdFilesFromEmpty(t *testing.T) {
	tt := []struct {
		path     string
		expected []string
		err      error
	}{
		{
			path: "./testdata/TestGetFiles/NonExistent",
			err:  os.ErrNotExist,
		},
		{
			path:     "./testdata/TestGetFiles/Empty",
			expected: []string{},
		},
		{
			path:     "./testdata/TestGetFiles/OnlyMdFiles",
			expected: []string{"001.md", "002.md"},
		},
		{
			path:     "./testdata/TestGetFiles/OnlyNonMdFiles",
			expected: []string{},
		},
		{
			path:     "./testdata/TestGetFiles/MixedFiles",
			expected: []string{"001.md", "002.md"},
		},
	}

	for _, tc := range tt {
		got, err := getMdFilesFrom(tc.path)
		if err != nil && !errors.Is(err, tc.err) {
			t.Errorf("Unexpected error received. Got: %v, Expected: %v\n", err, tc.err)
		}

		if len(got) != len(tc.expected) {
			t.Errorf("List of files is not what is expected. Got: %v, Expected:%v", fileNames(got), tc.expected)
		}

		for i, f := range got {
			if f.Name() != tc.expected[i] {
				t.Errorf("List of files is not what is expected. Got: %v, Expected:%v", fileNames(got), tc.expected)
			}
		}
	}
}
