package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

const (
	srcFile            = "./testdata/test.txt"
	expectedResultFile = "./testdata/expected.txt"
)

func TestReplace(t *testing.T) {

	testCases := []struct {
		origin   string
		old      string
		new      string
		expected string
	}{
		{"Go est un langage de gopher", "Go", "PHP", "PHP est un langage de gopher"},
		{"Go est un langage de Gopher go", "Go", "PHP", "PHP est un langage de Gopher PHP"},
	}

	for _, test := range testCases {
		result := replace(test.origin, "Go", "PHP", false)
		fmt.Println(result)
		if result != test.expected {
			t.Errorf("Expected : %s\n Got : %s", test.expected, result)
		}
	}

}

func TestFileReplace(t *testing.T) {

	tmp, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating tmp file %v", err)
	}
	defer os.Remove(tmp.Name())

	err = fileReplace(srcFile, tmp.Name(), "Go", "Python", false)
	if err != nil {
		t.Fatal(err)
	}

	rbyte, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatalf("Error when reading tmp file %v", err)
	}

	ebyte, err := os.ReadFile(expectedResultFile)
	if err != nil {
		t.Fatalf("Error when reading expected file")
	}

	if !bytes.Equal(rbyte, ebyte) {
		t.Fatalf("\n Expected : %s \n Got : %s", ebyte, rbyte)
	}

}
