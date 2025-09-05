package main

import "strings"

type TextType string

const (
	Text TextType = "Text"
	Bold TextType = "Bold"
	Italic TextType = "Italic"
	Code TextType = "Code"
	Link TextType = "Link"
	Image TextType = "Image"
)

type TextNode struct {
	text string
	textType TextType
	url string
}

func (tn TextNode) String() string {
	str := strings.Builder{}

	str.WriteString("{\n")
	str.WriteString("  Text: " + tn.text + ",\n")
	str.WriteString("  TextType: " + string(tn.textType) + ",\n")
	str.WriteString("  URL: " + tn.url + ",\n")
	str.WriteString("}\n")

	return str.String()
}
