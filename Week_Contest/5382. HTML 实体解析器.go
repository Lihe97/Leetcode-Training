package main

import (
	"fmt"
	"strings"
)

func entityParser(text string) string {

	text = strings.ReplaceAll(text,"&quot;","\"")
	text = strings.ReplaceAll(text,"&apos;","'")
	text = strings.ReplaceAll(text,"&amp;","&")
	text = strings.ReplaceAll(text,"&gt;",">")
	text = strings.ReplaceAll(text,"&lt;","<")
	text = strings.ReplaceAll(text,"&frasl;","/")

	return text
}
func main() {
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	
}
