package main

import (
	"fmt"
	"testing"
)

func TestPropsToHTML(t *testing.T) {
	tests := []struct{
		node *HTMLNode
		expected string
	}{
		{
			node: &HTMLNode{},
			expected: "",
		},
		{
			node: &HTMLNode{
				props: map[string]string{
					"href": "https://www.google.com",
				},
			},
			expected: " href=\"https://www.google.com\"",
		},
		{
			node: &HTMLNode{
				props: map[string]string{
					"href": "https://www.google.com",
					"target": "_blank",
				},
			},
			expected: " href=\"https://www.google.com\" target=\"_blank\"",
		},
	}

	for _, tc := range tests {
		got := tc.node.props_to_html()

		if tc.expected != got {
			t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
		}
	}
}
