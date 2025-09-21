package gopresent

import (
	"strings"
	"testing"
)

func TestNextSlide(t *testing.T) {
	tt := []struct {
		path           string
		nextSlideCount int
		expectedSlide  int
		expectedOut    string
	}{
		{
			path:           "./testdata/TestNextSlide/Empty",
			nextSlideCount: 0,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			path:           "./testdata/TestNextSlide/Empty",
			nextSlideCount: 1,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			path:           "./testdata/TestNextSlide/OneFile",
			nextSlideCount: 0,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			path:           "./testdata/TestNextSlide/OneFile",
			nextSlideCount: 1,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			path:           "./testdata/TestNextSlide/ThreeFiles",
			nextSlideCount: 1,
			expectedSlide:  1,
			expectedOut:    "two",
		},
	}

	for _, tc := range tt {
		p, err := NewPresentation(tc.path)
		if err != nil {
			t.Errorf("unable to create presentation with path: %v, err: %v\n", tc.path, err)
		}

		for range tc.nextSlideCount {
			p.NextSlide()
		}

		if p.curSlide != tc.expectedSlide {
			t.Errorf("Slide pointer is not correct. Got: %v, Expected: %v\n", p.curSlide, tc.expectedSlide)
		}

		if !strings.Contains(p.viewport.View(), tc.expectedOut) {
			t.Errorf("Viewport did not contain expected output. Got:\n%v\n, Expected to contain: %v\n", p.viewport.View(), tc.expectedOut)
		}
	}
}
