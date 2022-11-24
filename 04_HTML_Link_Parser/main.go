package main

import (
	"fmt"
	"io"
	"strings"
)

var exampleHtml = `
<html>
<body>
	<h1>Olá enfermeira!</h1>
	<a href="/karuta-club">Link pra página do Karuta Club</a>
	<a href="/animeshadow">Link para a página "The Anime Shadow", do Mateus e do Pedro.</a>
</body>
</html>
`

func main() {
	var r io.Reader = strings.NewReader(exampleHtml)

	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nPrinting links:%v\n", links)
}
