package control

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	hdwallet "github.com/stevelacy/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
	oCommon "web3_practice/common"
	"web3_practice/common/consts"
	ctrl "web3_practice/common/control"
	"web3_practice/common/exception"
	"web3_practice/lib/proto/blockchain"
)

var Wallet = &walletCtrl{}

type walletCtrl struct{}

func (c *walletCtrl) GetKeyPair(ctx *gin.Context) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		oCommon.Logger.Warningf("GetKeyPair GenerateKey fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		oCommon.Logger.Warningf("GetKeyPair publicKey fail. | err: %s", err)
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)[4:]
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	rsp := &blockchain.GetKeyPairRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.KeyPair{
			PrivateKey: privateKeyStr,
			PublicKey:  publicKeyStr,
			Address:    address,
			Mnemonic:   "",
		},
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *walletCtrl) GetMnemonic(ctx *gin.Context) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	address := account.Address.Hex()
	privateKey, _ := wallet.PrivateKeyHex(account)
	publicKey, _ := wallet.PublicKeyHex(account)

	rsp := &blockchain.GetMnemonicRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &blockchain.KeyPair{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
			Address:    address,
			Mnemonic:   mnemonic,
		},
	}

	ctrl.SendRsp(ctx, rsp)
}
