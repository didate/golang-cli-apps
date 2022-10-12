package main

import (
	"fmt"
	"testing"
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
