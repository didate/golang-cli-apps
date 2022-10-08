package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

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

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer, skipPreview bool) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	// create temp file and check for erros
	temp, err := os.CreateTemp("/tmp", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outFname := temp.Name()
	fmt.Fprintln(out, outFname)
	if err:= saveHTML(outFname, htmlData); err!=nil {
		return err
	}
	if skipPreview {
		return nil
	}
	defer os.Remove(outFname)
	return preview(outFname)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// creating html bloc
	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)
	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}

func preview(fname string) error{
	cName :=""
	cParams:= []string{}
	switch runtime.GOOS {
	case "linux":
		cName="xdg-open"
	case "windows":
		cName="cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName="open"
	default:
		return fmt.Errorf("OS not supported")
	}
	// append file to param slice
	cParams = append(cParams, fname)

	// locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err!=nil {
		return err
	}
	err = exec.Command(cPath, cParams...).Run()
	time.Sleep(2*time.Second)
	return err
}