package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {

	fname := flag.String("f", "", "File to process")
	oFname := flag.String("o", "", "Output file")
	oldW := flag.String("ow", "", "Old word")
	newW := flag.String("nw", "", "New word")
	cSensitive := flag.Bool("cs", false, "Case sensitive, default false")
	flag.Parse()

	switch {
	case *fname != "":
		err := fileReplace(*fname, *oFname, *oldW, *newW, *cSensitive)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	default:
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(replace(text, *oldW, *newW, *cSensitive))
	}

}

func replace(text, old, new string, cSensitive bool) string {
	pattern := `\b` + old + `\b`
	if !cSensitive {
		pattern =`(?i)\b` + old + `\b`
	}
	r := regexp.MustCompile(pattern)
	return r.ReplaceAllString(text, new)
}

func fileReplace(fname, oFname, old, new string, cSensitive bool) error {
		fmt.Println("debug")

	file, err := os.Open(fname)
	
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	outFile, err := os.OpenFile(oFname, os.O_WRONLY|os.O_CREATE, 0644)
	if err!=nil {
		return err
	}
	
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		result := replace(line, old, new, cSensitive)
		fmt.Fprintln(writer, result)
	}
	return nil
}
