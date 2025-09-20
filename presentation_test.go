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
			path:           "./testdata/TestNextSlide/Empty",
			nextSlideCount: 1,
			expected:       "No content to show",
		},
		// {
		// 	path: "./testdata/TestNextSlide/Empty",
		// 	nextSlideCount: 1,
		// 	expected: "No content to show",
		// },
		// {
		// 	// call next when the dir only has one file
		// 	path: "./testdata/TestNextSlide/OneFile",
		// 	nextSlideCount: 1,
		// 	expected: "one",
		// },
		// {
		// 	// call next 3 times when the dir has 2 files
		// },
		// {
		// 	// call next once when the dir has 2 file
		// },
	}

	for _, tc := range tt {
		p, err := NewPresentation(tc.path)
		if err != nil {
			t.Errorf("unable to create presentation with path: %v, err: %v\n", tc.path, err)
		}

		var got string
		for range tc.nextSlideCount {
			got, _ = p.NextSlide()
		}

		if got != tc.expected {
			t.Errorf("Expected content from NextSlide is not what is expected. Got: %v, Expected: %v\n", got, tc.expected)
		}
	}
}
