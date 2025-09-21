package main
import (
	"strings"
	"testing"
)

func TestPrevNextSlide(t *testing.T) {
	tt := []struct {
		path           string
		prevSlideCount int
		nextSlideCount int
		expectedSlide  int
		expectedOut    string
	}{
		{
			path:           "./testdata/TestPrevNextSlide/Empty",
			prevSlideCount: 0,
			nextSlideCount: 0,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			path:           "./testdata/TestPrevNextSlide/Empty",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			path:           "./testdata/TestPrevNextSlide/OneFile",
			prevSlideCount: 0,
			nextSlideCount: 0,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			path:           "./testdata/TestPrevNextSlide/OneFile",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			path:           "./testdata/TestPrevNextSlide/ThreeFiles",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  1,
			expectedOut:    "two",
		},
		{
			path: "./testdata/TestPrevNextSlide/Empty",
			nextSlideCount: 0,
			prevSlideCount: 0,
			expectedSlide: -1,
			expectedOut: "No content to show",
		},
		{
			path: "./testdata/TestPrevNextSlide/Empty",
			nextSlideCount: 0,
			prevSlideCount: 1,
			expectedSlide: -1,
			expectedOut: "No content to show",
		},
		{
			path: "./testdata/TestPrevNextSlide/OneFile",
			nextSlideCount: 0,
			prevSlideCount: 1,
			expectedSlide: 0,
			expectedOut: "one",
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

		for range tc.nextSlideCount {
			p.PrevSlide()
		}

		if p.curSlide != tc.expectedSlide {
			t.Errorf("Slide pointer is not correct. Got: %v, Expected: %v\n", p.curSlide, tc.expectedSlide)
		}

		if !strings.Contains(p.viewport.View(), tc.expectedOut) {
			t.Errorf("Viewport did not contain expected output. Got:\n%v\n, Expected to contain: %v\n", p.viewport.View(), tc.expectedOut)
		}
	}
}
