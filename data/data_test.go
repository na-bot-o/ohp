package data

import (
	"os"
	"testing"
)

func TestGetDataFilePath(t *testing.T) {
	datafile := New("./ohp")

	home := os.Getenv("HOME")

	if datafile.Path != home+"/.ohp" {
		t.Fatalf("filepath is wrong")
	}

}
