package main 

import "testing"

func TestTextnodeEquality(t *testing.T) {
	tn1 := TextNode {
		text: "This is a text node",
		textType: Bold,
	}

	tn2 := TextNode {
		text: "This is a text node",
		textType: Bold,
	}

	if tn1 != tn2 {
		t.Errorf("text nodes %v and %v should be equal", tn1, tn2)
	}
}
