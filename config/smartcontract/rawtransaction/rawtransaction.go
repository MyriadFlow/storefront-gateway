package rawtransaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract"
	"github.com/MyriadFlow/storefront-gateway/generated/smartcontract/signatureSeries"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/ethwallet"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/sirupsen/logrus"
)

func SendRawTransaction(abiS string, method string, args ...interface{}) (*types.Transaction, error) {

	abiP, err := abi.JSON(strings.NewReader(abiS))
	if err != nil {
		logwrapper.Errorf("failed to parse JSON abi, error %v", err)
		return nil, err
	}
	client, err := smartcontract.GetClient()
	if err != nil {
		return nil, err
	}
	mnemonic := envconfig.EnvVars.MNEMONIC
	privateKey, publicKey, _, err := ethwallet.HdWallet(mnemonic) // Verify: https://iancoleman.io/bip39/
	if err != nil {
		logwrapper.Errorf("failed to get private and public key from mnemonic, error %v", err.Error())
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}
	envContractAddress := envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS

	toAddress := common.HexToAddress(envContractAddress)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to call client.NetworkID, error: %v", err.Error())
		return nil, err
	}

	bytesData, err := abiP.Pack(method, args...)
	if err != nil {
		logwrapper.Errorf("failed to pack trasaction of method %v, error: %v", method, err)
		return nil, err
	}

	logwrapper.Infof("nonce is %v", nonce)

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to suggestGasTipCap, error %v", err)
		return nil, err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(80001),
	}
	bn, _ := client.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	big2 := big.NewInt(2)
	mulRes := big.NewInt(0).Mul(baseFee, big2)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       310000,
		To:        &toAddress,
		Data:      bytesData,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		logwrapper.Errorf("failed to sign trasaction %v, error: %v", tx, err.Error())
		return nil, err
	}

	err = client.SendTransaction(context.TODO(), signedTx)
	if err != nil {
		logwrapper.Error("failed to send trasaction, error: ", err)
		return nil, err
	}
	return signedTx, nil
}

func SendRawTransactionDelegateSignature(abiS string, method string, address string, creator common.Address, metdata string, royalty *big.Int) (*types.Transaction, error) {

	client, err := smartcontract.GetClient()
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(envconfig.EnvVars.WALLET_PRIVATE_KEY)
	if err != nil {
		logwrapper.Errorf("failed to parse Private Key, error %v", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logwrapper.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey, error %v", err)
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		if err != nil {
			logwrapper.Warnf("failed to get gas price")
			return nil, err
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		logwrapper.Warnf("failed to get transactops")
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	addressVar := common.HexToAddress(address)
	instance, err := signatureSeries.NewSignatureSeries(addressVar, client)
	if err != nil {
		logwrapper.Warnf("failed to create instance")
		return nil, err
	}
	tx, err := instance.DelegateAssetCreation(auth, creator, metdata, royalty)
	if err != nil {
		logwrapper.Warnf("failed to execute transaction")
		return nil, err
	}
	logrus.Info("transaction send", tx.Hash().String())
	return tx, nil
}

func SendRawTransactionCreateAssetSignature(abiS string, address string, metdata string, royalty *big.Int) (*types.Transaction, error) {

	client, err := smartcontract.GetClient()
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(envconfig.EnvVars.WALLET_PRIVATE_KEY)
	if err != nil {
		logwrapper.Errorf("failed to parse Private Key, error %v", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logwrapper.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey, error %v", err)
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		if err != nil {
			logwrapper.Warnf("failed to get gas price")
			return nil, err
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		logwrapper.Warnf("failed to get transactops")
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	addressVar := common.HexToAddress(address)
	instance, err := signatureSeries.NewSignatureSeries(addressVar, client)
	if err != nil {
		logwrapper.Warnf("failed to create instance")
		return nil, err
	}
	tx, err := instance.CreateAsset(auth, metdata, royalty)
	if err != nil {
		logwrapper.Warnf("failed to execute transaction")
		return nil, err
	}
	logrus.Info("transaction send", tx.Hash().String())
	return tx, nil
}

func SendRawTransactionTransferAssetSignature(abiS string, address string, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {

	client, err := smartcontract.GetClient()
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(envconfig.EnvVars.WALLET_PRIVATE_KEY)
	if err != nil {
		logwrapper.Errorf("failed to parse Private Key, error %v", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logwrapper.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey, error %v", err)
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		if err != nil {
			logwrapper.Warnf("failed to get gas price")
			return nil, err
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		logwrapper.Warnf("failed to get transactops")
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	addressVar := common.HexToAddress(address)
	instance, err := signatureSeries.NewSignatureSeries(addressVar, client)
	if err != nil {
		logwrapper.Warnf("failed to create instance")
		return nil, err
	}
	tx, err := instance.SafeTransferFrom(auth, from, to, tokenId)
	if err != nil {
		logwrapper.Warnf("failed to execute transaction")
		return nil, err
	}
	logrus.Info("transaction send", tx.Hash().String())
	return tx, nil
}
