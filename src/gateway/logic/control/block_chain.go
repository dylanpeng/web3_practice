package control

import (
	"github.com/gin-gonic/gin"
	"web3_practice/common/consts"
	ctrl "web3_practice/common/control"
	"web3_practice/lib/proto/blockchain"
)

var BlockChain = &blockChainCtrl{}

type blockChainCtrl struct{}

func (c *blockChainCtrl) GetTransactionDetail(ctx *gin.Context) {
	req := &blockchain.GetTransactionReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Tx == "") {
		return
	}

	rsp := &blockchain.GetTransactionRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.Transaction{
			Tx:       "aaa",
			From:     "bbb",
			To:       "ccc",
			Amount:   "0.1154",
			GasFee:   "0.001564548",
			GasPrice: "0.0000024564",
		},
	}

	ctrl.SendRsp(ctx, rsp)
}
