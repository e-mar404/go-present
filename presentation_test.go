package gopresent

import (
	"testing"
)

func TestNextSlide(t *testing.T) {
	tt := []struct {
		path           string
		nextSlideCount int
		expected       string
	}{
		{
			// no next slide count when the dir is empty
			path:           "./testdata/TestNextSlide/Empty",
			nextSlideCount: 0,
			expected:       "No content to show",
		},
		// {
		// 	// call next once when the dir is empty
		// },
		// {
		// 	// call next when the dir only has one file
		// },
		// {
		// 	// call next 3 times when the dir has 2 files
		// },
		// {
		// 	// call next once when the dir has 2 file
		// },
	}

	for _, tc := range tt {
		_, err := NewPresentation(tc.path)
		if err != nil {
			t.Errorf("unable to create presentation with path: %v, err: %v\n", tc.path, err)
		}
	}
}
