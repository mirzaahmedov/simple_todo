package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	BindingAddress string `toml:"binding_address"`
	DatabaseURL    string `toml:"database_url"`
}

func Load(filePath string) (*Config, error) {
	c := &Config{
		BindingAddress: ":8080",
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
