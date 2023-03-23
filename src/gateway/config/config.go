package config

import (
	"github.com/BurntSushi/toml"
	"github.com/dylanpeng/golib/web3_eth"
	oConf "web3_practice/common/config"
)

var conf *Config

type Config struct {
	*oConf.Config
	Eth *web3_eth.Config `toml:"web3_eth" json:"web3_eth"`
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
