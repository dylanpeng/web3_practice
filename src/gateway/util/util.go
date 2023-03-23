package util

import (
	"github.com/dylanpeng/golib/web3_eth"
	"web3_practice/common"
	"web3_practice/gateway/config"
)

var EthClient *web3_eth.Client

func InitEthClient() (err error) {
	ethConfig := config.GetConfig().Eth

	EthClient, err = web3_eth.NewClient(ethConfig)

	if err != nil {
		common.Logger.Errorf("InitEthClient fail. | err: %s", err)
		return err
	}
	return nil
}
