package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"flag"
)

func main(){

	countLines := flag.Bool("l", false, "Count Line")
	countBytes := flag.Bool("b", false, "Count Line")
	flag.Parse()
	fmt.Println(count(os.Stdin, *countLines, *countBytes))
}

func count(r io.Reader, clines bool, cbytes bool) int{

	scanner := bufio.NewScanner(r)
	if(cbytes){
		scanner.Split(bufio.ScanBytes)
	}else if !clines {
		scanner.Split(bufio.ScanWords)
	} 

	wc :=0
	for scanner.Scan() {
		wc++
	}
	return wc;
}