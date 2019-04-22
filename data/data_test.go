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
	// check that match expected inputs
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

	//check that non-match unexpexted inputs

	unExpectedInputs := []struct {
		Name string
		Tag  string
		Url  string
	}{
		{"Facebook", "SNS", "https://facebook.com"},
		{"Slack", "Chat", "https://slack.com"},
	}

	for _, unExpectedInput := range unExpectedInputs {
		for _, page := range pages {
			if unExpectedInput == page {
				t.Errorf("found unexpected value : %+v", page)
			}

		}
	}

}

//check whether execute copyto function correctly
func TestCheckSameDataInOhpAndArchive(t *testing.T) {

	dataStub := Data{TESTFILENAME, TESTFILEPATH}
	archiveStub := Data{ARCHIVEFILE, ARCHIVEFILEPATH}

	_, err := os.Create(archiveStub.Path)

	if err != nil {
		t.Error(err)
		return
	}

	defer os.Remove(ARCHIVEFILEPATH)

	dataStub.CopyTo(archiveStub)

	dataPages, err := dataStub.GetPages()

	if err != nil {
		t.Error(err)
		return
	}

	archivePages, err := archiveStub.GetPages()

	if err != nil {
		t.Error(err)
		return
	}

	pageRows := len(dataPages)
	index := 0

	for {
		if index == pageRows-1 {
			break
		}

		if dataPages[index] != archivePages[index] {
			t.Errorf("data isn't copied, data is %+v, archive is %+v", dataPages[index], archivePages[index])
			return
		}
		index++
	}

}
