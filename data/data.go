package data

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/na-bot-o/ohp/page"
)

//Data struct is object to manage data file
type Data struct {
	Name string
	Path string
}

//CopyTo func archives .ohp file for recovering
func (d *Data) CopyTo(archive Data) {

	oldFile, err := os.Create(archive.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer oldFile.Close()

	file, err := os.OpenFile(d.Path, os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(oldFile, file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

//GetPages gets page info in date file
func (d *Data) GetPages() (pages []page.Page, err error) {
	file, err := os.OpenFile(d.Path, os.O_RDONLY, 0755)
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

		page := page.New(fields[0], fields[1], fields[2])

		pages = append(pages, page)
	}
	return pages, nil
}

//New func is create data object
func New(name string) Data {

	home, err := homedir.Dir()
	filepath := home + name

	if err != nil {
		fmt.Println("can't get homedir")
		os.Exit(1)
	}

	return Data{name, filepath}
}
