package file

import homedir "github.com/mitchellh/go-homedir"

type Datafile struct {
	name string
}

func (df *Datafile) GetPath() (string, error) {

	home, err := homedir.Dir()
	filepath := home + df.name

	if err != nil {
		return "", err
	}

	return filepath, nil

}
func New(name string) Datafile {
	return Datafile{name}
}
