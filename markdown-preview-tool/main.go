package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const defaultTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>{{.Body}}</body>
</html>`

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename, tFname string, out io.Writer, skipPreview bool) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, tFname)

	if err != nil {
		return err
	}
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
	if err := saveHTML(outFname, htmlData); err != nil {
		return err
	}
	if skipPreview {
		return nil
	}
	defer os.Remove(outFname)
	return preview(outFname)
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}
	// creating html bloc
	var buffer bytes.Buffer

	// execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}
	// append file to param slice
	cParams = append(cParams, fname)

	// locate executable in PATH
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}
	err = exec.Command(cPath, cParams...).Run()
	time.Sleep(2 * time.Second)
	return err
}
