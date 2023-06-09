package router

import (
	"github.com/gin-gonic/gin"
	"web3_practice/common"
	ctrl "web3_practice/common/control"
	"web3_practice/common/middleware"
	"web3_practice/gateway/logic/control"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	return common.GetTraceId(ctx)
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)
	app.Use(middleware.Trace)

	blockchainGroup := app.Group("/blockchain")
	{
		blockchainGroup.GET("/transaction/detail/:tx", control.BlockChain.GetTransactionDetail)
		blockchainGroup.GET("/account/detail", control.BlockChain.GetAccountDetail)
		blockchainGroup.GET("/gas/suggest", control.BlockChain.GetEthSuggestGas)
	}

	walletGroup := app.Group("/wallet")
	{
		walletGroup.GET("/key/create", control.Wallet.GetKeyPair)
		walletGroup.GET("/mnemonic/create", control.Wallet.GetMnemonic)
	}
}
