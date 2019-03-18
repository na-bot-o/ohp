package data

import "testing"

func TestGetDataFilePath(t *testing.T) {
	_, err := GetDataFilePath()

	if err != nil {
		t.Fatalf("can't get data file path")
	}
}
