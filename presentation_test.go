package main

import (
	"strings"
	"testing"
)

func TestPrevNextSlide(t *testing.T) {
	tt := []struct {
		title          string
		path           string
		prevSlideCount int
		nextSlideCount int
		expectedSlide  int
		expectedOut    string
	}{
		{
			title:          "No movement on empty dir",
			path:           "./testdata/TestPrevNextSlide/Empty",
			prevSlideCount: 0,
			nextSlideCount: 0,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			title:          "NextSlide on empty dir",
			path:           "./testdata/TestPrevNextSlide/Empty",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			title:          "No movement on dir with 1 file",
			path:           "./testdata/TestPrevNextSlide/OneFile",
			prevSlideCount: 0,
			nextSlideCount: 0,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			title:          "NextSlide on dir with 1 file",
			path:           "./testdata/TestPrevNextSlide/OneFile",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			title:          "NextSlide when there are multiple files",
			path:           "./testdata/TestPrevNextSlide/ThreeFiles",
			prevSlideCount: 0,
			nextSlideCount: 1,
			expectedSlide:  1,
			expectedOut:    "two",
		},
		{
			title:          "PrevSlide on empty dir",
			path:           "./testdata/TestPrevNextSlide/Empty",
			nextSlideCount: 0,
			prevSlideCount: 1,
			expectedSlide:  -1,
			expectedOut:    "No content to show",
		},
		{
			title:          "PrevSlide on dir with 1 file",
			path:           "./testdata/TestPrevNextSlide/OneFile",
			nextSlideCount: 0,
			prevSlideCount: 1,
			expectedSlide:  0,
			expectedOut:    "one",
		},
		{
			title:          "PrevSlide on dir with multiple files",
			path:           "./testdata/TestPrevNextSlide/ThreeFiles",
			nextSlideCount: 2,
			prevSlideCount: 1,
			expectedOut:    "two",
			expectedSlide:  1,
		},
	}

	for _, tc := range tt {
		p, err := NewPresentation(tc.path)
		if err != nil {
			t.Errorf("%v\nunable to create presentation with path: %v, err: %v\n", tc.title, tc.path, err)
		}

		for range tc.nextSlideCount {
			p.NextSlide()
		}

		for range tc.prevSlideCount {
			p.PrevSlide()
		}

		if p.curSlide != tc.expectedSlide {
			t.Errorf("%v\nSlide pointer is not correct. Got: %v, Expected: %v\n", tc.title, p.curSlide, tc.expectedSlide)
		}

		if !strings.Contains(p.viewport.View(), tc.expectedOut) {
			t.Errorf("%v\nViewport did not contain expected output. Got:\n%v\n, Expected to contain: %v\n", tc.title, p.viewport.View(), tc.expectedOut)
		}
	}
}
