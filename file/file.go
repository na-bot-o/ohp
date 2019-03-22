package file

import homedir "github.com/mitchellh/go-homedir"

type DataFile struct {
	name string
}

func (df *DataFile) GetPath() (string, error) {

	home, err := homedir.Dir()
	filepath := home + df.name

	if err != nil {
		return "", err
	}

	return filepath, nil

}
func New(name string) DataFile {
	return DataFile{name}
}
