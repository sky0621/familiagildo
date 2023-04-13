package setup

type Env string

func (e Env) IsGCP() bool {
	return e == "gcp"
}

func (e Env) IsLocal() bool {
	return e == "local"
}
