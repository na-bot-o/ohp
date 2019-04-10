package page

import (
	"os"
)

type Page struct {
	Name string
	Tag  string
	Url  string
}

func New(name string, tag string, url string) Page {
	return Page{name, tag, url}
}

func Write(file *os.File, output string) error {
	_, err := file.Write(([]byte)(output))

	return err
}
