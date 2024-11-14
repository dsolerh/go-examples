package configs

import "flag"

// GlobalConfigs contains the configuration for the Api
type GlobalConfigs struct {
	*RepositoryConfigs
	*APIConfigs
	Port string
}

type APIConfigs struct {
}

type RepositoryConfigs struct {
}

func New() *GlobalConfigs {
	c := new(GlobalConfigs)
	repo := new(RepositoryConfigs)
	api := new(APIConfigs)
	c.RepositoryConfigs = repo
	c.APIConfigs = api

	flag.StringVar(&c.Port, "port", "8080", "the port for the server to listen")

	flag.Parse()
	return c
}
