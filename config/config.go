package config

import (
	"github.com/BurntSushi/toml"
)

// Config is struct for config
type Config struct {
	DB DBConfig `toml:"database"`
}

// NewConfig is create Config from toml file
func NewConfig(path string, env string) (Config, error) {
	var conf Config

	confPath := path + env + ".toml"

	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}
