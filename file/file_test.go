package file

import (
	"os"
	"testing"
)

func TestGetDataFilePath(t *testing.T) {
	datafile := New("./ohp")

	filepath, err := datafile.GetPath()

	if err != nil {
		t.Fatalf("can't get data file path")
	}

	home := os.Getenv("HOME")

	if filepath != home+"/.ohp" {
		t.Fatalf("filepath is wrong")
	}

}
