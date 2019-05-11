package configs

import (
	"os"
)

// Configs contains all the environment varialbes.
type Configs struct {
	Port string
}

// New returns the environment varialbes.
func New() Configs {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	return Configs{
		Port: port,
	}
}
