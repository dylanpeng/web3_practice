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

	apiGroup := app.Group("/blockchain")
	{
		apiGroup.POST("/transaction/detail", control.BlockChain.GetTransactionDetail)
	}
}
