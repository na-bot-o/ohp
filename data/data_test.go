package data

import (
	"os"
	"testing"
)

var HOMEDIR = os.Getenv("HOME")
var TESTFILENAME = "./ohp_test"
var TESTFILEPATH = "../.ohp_test"
var ARCHIVEFILE = ".ohp_old_test"
var ARCHIVEFILEPATH = "../.ohp_old_test"

func TestGetDataFilePath(t *testing.T) {
	datafile := New(".ohp")

	if datafile.Path != HOMEDIR+"/.ohp" {
		t.Fatalf("filepath is wrong")
	}

}
func TestGetPages(t *testing.T) {
	expectedInputs := []struct {
		Name string
		Tag  string
		Url  string
	}{
		{"Amazon", "Ecommerce", "https://amazon.com"},
		{"Google", "SearchEngine", "https://google.com"},
		{"Twitter", "SNS", "https://twitter.com"},
		{"Yahoo", "SearchEngine", "https://yahoo.co.jp"},
	}

	dataStub := Data{TESTFILENAME, TESTFILEPATH}

	pages, err := dataStub.GetPages()

	if err != nil {
		t.Error(err)
		return
	}

	for _, expectedInput := range expectedInputs {
		for index, page := range pages {
			if expectedInput == page {
				break
			}
			if index == len(pages)-1 {
				t.Errorf("not found expected value : %+v", page)
			}
		}
	}

}
