package main

type config struct {
	pgConnStr string
	servPort  string
}

func newConfig() (config, error) {
	return config{
		pgConnStr: "asd",
		servPort:  ":444",
	}, nil
}

