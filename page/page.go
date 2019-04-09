package page

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Page struct {
	Name string
	Tag  string
	Url  string
}

func New(name string, tag string, url string) Page {
	return Page{name, tag, url}
}

func GetRows(filepath string) (rows []Page, err error) {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4096)

	for {
		row, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fields := strings.Split(string(row), ",")

		page := New(fields[0], fields[1], fields[2])

		rows = append(rows, page)
	}
	return rows, nil
}

func Write(file *os.File, output string) error {
	_, err := file.Write(([]byte)(output))

	return err
}
