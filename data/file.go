package data

import homedir "github.com/mitchellh/go-homedir"

func GetDataFilePath() (string, error) {
	home, err := homedir.Dir()
	filepath := home + "/.ohp"

	if err != nil {
		return "", err
	}

	return filepath, nil
}
