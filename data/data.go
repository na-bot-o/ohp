package data

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Data struct {
	Name string
	Tag  string
	Url  string
}

func New(name string, tag string, url string) Data {
	return Data{name, tag, url}
}

func GetRows(filepath string) (rows []string, err error) {
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

		rows = append(rows, string(row))
	}
	return rows, nil
}

//Archive .ohp file for recovering
func ArchiveFile(filepath string, old_filepath string) {

	old_file, err := os.Create(old_filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer old_file.Close()

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)

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

func Write(file *os.File, output string) error {
	_, err := file.Write(([]byte)(output))

	return err
}
