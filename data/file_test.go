package data

import (
	"os"
	"testing"
)

func TestGetDataFilePath(t *testing.T) {
	filepath, err := GetDataFilePath()

	if err != nil {
		t.Fatalf("can't get data file path")
	}

	home := os.Getenv("HOME")

	if filepath != home+"/.ohp" {
		t.Fatalf("filepath is wrong")
	}

}
