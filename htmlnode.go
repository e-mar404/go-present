package main

import (
	"fmt"
	"strings"
)

type HTMLNode struct {
	tag string
	value string
	children []HTMLNode
	props map[string]string
}

func (hn *HTMLNode) to_html() (string, error) {
	return "", fmt.Errorf("NotImplemented")
}

func (hn *HTMLNode) props_to_html() string {
	str := strings.Builder{}

	for key, value := range(hn.props) {
		str.WriteString(" " + key + "=\"" + value + "\"")
	}

	return str.String()
}

func (hn HTMLNode) String() string {
	str := strings.Builder{}

	str.WriteString("Tag: " + hn.tag + "\n")
	str.WriteString("value: " + hn.value+ "\n")
	str.WriteString("Children: ")
	for _, child := range(hn.children) {
		str.WriteString(child.String())
	}
	str.WriteString("\n")
	str.WriteString("Props: [")
	for value, key := range(hn.props) {
		str.WriteString(value + ":" + key+ ",")
	}
	str.WriteString("]\n")

	return str.String()
}
