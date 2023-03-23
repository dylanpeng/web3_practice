package control

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	oCommon "web3_practice/common"
	"web3_practice/common/consts"
	ctrl "web3_practice/common/control"
	"web3_practice/common/exception"
	"web3_practice/gateway/util"
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

	// 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	account := common.HexToAddress("0x5B7f33E9f0B24465cBD575d3cb354D286a9DF576")
	balance, err := util.EthClient.GetClient().BalanceAt(context.Background(), account, nil)
	if err != nil {
		oCommon.Logger.Warningf("GetTransactionDetail BalanceAt fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	oCommon.Logger.Infof("balance: %s", balance)

	rsp := &blockchain.GetTransactionRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.Transaction{
			Tx:       "aaa",
			From:     "bbb",
			To:       "ccc",
			Amount:   balance.String(),
			GasFee:   "0.001564548",
			GasPrice: "0.0000024564",
		},
	}

	ctrl.SendRsp(ctx, rsp)
}
