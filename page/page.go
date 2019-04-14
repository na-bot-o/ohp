package page

import "os"

type Page struct {
	Name string
	Tag  string
	Url  string
}

func New(name string, tag string, url string) Page {
	return Page{name, tag, url}
}

func (p *Page) WrittenIn(file *os.File) error {

	insert_format := p.Name + "," + p.Tag + "," + p.Url + "\n"
	_, err := file.Write(([]byte)(insert_format))

	return err
}
