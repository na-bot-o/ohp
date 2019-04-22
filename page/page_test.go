package page

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

var WRITTENFILENAME = "./ohp_written_test"
var WRITTENFILEPATH = "../.ohp_written_test"

func TestWrittenData(t *testing.T) {

	type pages []*Page

	writtenFile, err := os.Create(WRITTENFILEPATH)

	if err != nil {
		t.Error("can't create file")
	}
	defer os.Remove(WRITTENFILEPATH)

	inputs := []struct {
		Name string
		Tag  string
		Url  string
	}{
		{"Amazon", "Ecommerce", "https://amazon.com"},
		{"Google", "SearchEngine", "https://google.com"},
		{"Twitter", "SNS", "https://twitter.com"},
		{"Yahoo", "SearchEngine", "https://yahoo.co.jp"},
	}

	for _, input := range inputs {
		pageData := New(input.Name, input.Tag, input.Url)
		pageData.WrittenIn(writtenFile)
	}

	writtenFile.Close()

	readFile, err := os.OpenFile(WRITTENFILEPATH, os.O_RDONLY, 0666)
	if err != nil {
		t.Error("failed to open file")
	}

	reader := bufio.NewReaderSize(readFile, 4096)

	for _, input := range inputs {
		row, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fields := strings.Split(string(row), ",")

		if input.Name != fields[0] || input.Tag != fields[1] || input.Url != fields[2] {
			t.Errorf("inputted wrong data expected: %+v, inputted: %s", input, string(row))
		}

	}

}
