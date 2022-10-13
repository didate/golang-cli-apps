package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {

	b := bytes.NewBufferString("Ceci est une phrase")
	exp := 4
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("Ceci est une phrase\n voici une autre phrase\n une derniere phrase")
	exp := 3
	res := count(b, true, false)
	if exp != res {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}

func TestBytesCount(t *testing.T) {
	b := bytes.NewBufferString("Ceci est une phrase")

	exp := 19
	res := count(b, false, true)
	if exp != res {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}
