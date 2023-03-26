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
	req := &blockchain.GetTransactionReq{
		Tx: ctx.Param("tx"),
	}

	if !ctrl.ParamAssert(ctx, req, req.Tx == "") {
		return
	}

	txHash := common.HexToHash(req.Tx)
	tx, ispending, err := util.EthClient.GetClient().TransactionByHash(context.Background(), txHash)
	if err != nil {
		oCommon.Logger.Warningf("GetTransactionDetail TransactionByHash fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	receipt, err := util.EthClient.GetClient().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		oCommon.Logger.Warningf("GetTransactionDetail TransactionReceipt fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	sender, err := util.EthClient.GetClient().TransactionSender(context.Background(), tx, receipt.BlockHash, receipt.TransactionIndex)
	if err != nil {
		oCommon.Logger.Warningf("GetTransactionDetail TransactionSender fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	rsp := &blockchain.GetTransactionRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.Transaction{
			Tx:                req.Tx,
			From:              sender.String(),
			To:                tx.To().String(),
			Amount:            tx.Value().String(),
			Gas:               tx.Gas(),
			GasPrice:          tx.GasPrice().String(),
			GasTip:            tx.GasTipCap().String(),
			GasFee:            tx.GasFeeCap().String(),
			Nonce:             tx.Nonce(),
			Cost:              tx.Cost().String(),
			ChainId:           tx.ChainId().String(),
			IsPending:         ispending,
			BlockHash:         receipt.BlockHash.String(),
			BlockNumber:       receipt.BlockNumber.String(),
			TransactionIndex:  uint64(receipt.TransactionIndex),
			Type:              uint32(receipt.Type),
			Status:            receipt.Status,
			GasUsed:           receipt.GasUsed,
			EffectiveGasPrice: receipt.EffectiveGasPrice.String(),
		},
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *blockChainCtrl) GetAccountDetail(ctx *gin.Context) {
	req := &blockchain.GetAccountDetailReq{
		Address: ctx.Query("account"),
	}

	if !ctrl.ParamAssert(ctx, req, req.Address == "") {
		return
	}

	// 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	account := common.HexToAddress(req.Address)
	balance, err := util.EthClient.GetClient().BalanceAt(context.Background(), account, nil)
	if err != nil {
		oCommon.Logger.Warningf("GetAccountDetail BalanceAt fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	nonce, err := util.EthClient.GetClient().PendingNonceAt(context.Background(), account)
	if err != nil {
		oCommon.Logger.Warningf("GetAccountDetail NonceAt fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	rsp := &blockchain.GetAccountDetailRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.EthAccount{
			Address: req.Address,
			Balance: balance.String(),
			Nonce:   nonce,
		},
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *blockChainCtrl) GetEthSuggestGas(ctx *gin.Context) {
	gasPrice, err := util.EthClient.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		oCommon.Logger.Warningf("GetEthSuggestGas SuggestGasPrice fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	gasTipCap, err := util.EthClient.GetClient().SuggestGasTipCap(context.Background())
	if err != nil {
		oCommon.Logger.Warningf("GetAccountDetail NonceAt fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	rsp := &blockchain.GetEthSuggestGasRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.GasDetail{
			GasPriceSuggest:  gasPrice.String(),
			GasTipCapSuggest: gasTipCap.String(),
		},
	}

	ctrl.SendRsp(ctx, rsp)
}
