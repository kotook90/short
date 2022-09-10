package rout

import (
	"fmt"
	"testing"
)

func TestValidateData(t *testing.T) {

	a := "."
	b := "www"
	c := "http"
	d := "https"
	test := []string{".", "www", "http", "https"}
	for _, v := range test {
		if v != a || v != b || v != c || v != d {
			fmt.Println("test failed")
			return
		}
		fmt.Printf("res: %s\n", v)
	}

}
