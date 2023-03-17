package config

import (
	"github.com/BurntSushi/toml"
	oConf "web3_practice/common/config"
)

var conf *Config

type Config struct {
	*oConf.Config
}

func Init(file string) error {
	conf = &Config{
		Config: oConf.Default(),
	}

	if _, err := toml.DecodeFile(file, conf); err != nil {
		return err
	}

	if err := conf.Config.Init(); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return conf
}
