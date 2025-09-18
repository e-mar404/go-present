package gopresent 

import (
	"testing"
)

func TestDefaultSlideRenderer(t *testing.T) {
	content := "# hello"
	renderer, err := NewSlideRenderer(
		WithGlamourDefault(),
	)
	if err != nil {
		t.Errorf("could not create default glamour slide renderer: %v\n", err)
	}

	_, err = renderer.Render(content)
	if err != nil {
		t.Errorf("could not render content[%v], got error: %v\n", content, err)
	}
}

func TestConsecutiveSlideRenders(t *testing.T) {
	content := "# hello"
	renderer, err := NewSlideRenderer(
		WithGlamourDefault(),
	)
	if err != nil {
		t.Errorf("could not create default glamour slide renderer: %v\n", err)
	}

	_, err = renderer.Render(content)
	_, err = renderer.Render(content)
	if err != nil {
		t.Errorf("could not render content [%v] consecutively, got error: %v\n", content, err)
	}
}
