package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	countLines := flag.Bool("l", false, "Count Line")
	countBytes := flag.Bool("b", false, "Count bytes")
	filename := flag.String("f", "", "File to process")
	flag.Parse()
	switch {
	case *filename != "":
		file, err := os.Open(*filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(count(file, *countLines, *countBytes))
	default:
		fmt.Println()
		fmt.Println(count(os.Stdin, *countLines, *countBytes))

	}
}

func count(r io.Reader, clines bool, cbytes bool) int {

	scanner := bufio.NewScanner(r)
	if cbytes {
		scanner.Split(bufio.ScanBytes)
	} else if !clines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
