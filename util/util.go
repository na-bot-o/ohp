package util

type Env struct {
	FileName    string
	ArchiveName string
}

func LoadEnv() Env {
	//hard coding file and archive name
	return Env{".ohp", ".ohp_old"}
}
