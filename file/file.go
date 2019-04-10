package data

import (
	"fmt"
	"io"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

type Data struct {
	Name string
	Path string
}

// func (df *DataFile) GetPath() (string, error) {

// 	home, err := homedir.Dir()
// 	filepath := home + df.Name

// 	if err != nil {
// 		return "", err
// 	}

// 	return filepath, nil

// }

//Archive .ohp file for recovering
func (d *Data) CopyTo(archive Data) {

	old_file, err := os.Create(archive.Path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer old_file.Close()

	file, err := os.OpenFile(d.Path, os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(old_file, file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func New(name string) Data {

	home, err := homedir.Dir()
	filepath := home + name

	if err != nil {
		fmt.Println("can't get homedir")
		os.Exit(1)
	}

	return Data{name, filepath}
}
