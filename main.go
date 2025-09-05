package main

import (
	"fmt"
)

func main() {
	tn1 := TextNode {
		text: "This is some anchor text",		
		textType: Link,
		url: "https://www.boot.dev",
	}

	fmt.Println(tn1)
}
