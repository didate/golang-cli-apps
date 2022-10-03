package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Markdown Preview Tool</title>
</head>
<body>`

	footer = `</body>
</html>`
)

func main()  {
	// Parse flags
	filename := flag.String("file","","Markdown file to preview")
	flag.Parse()

	if *filename==""{
		flag.Usage()
		os.Exit(1)
	}
	if err:= run(*filename); err!=nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	outFile := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outFile)
	return saveHTML(outFile, htmlData)
}

func parseContent(input []byte) []byte{
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// creating html bloc
	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)
	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error{
	return os.WriteFile(outFname, data, 0644)
}