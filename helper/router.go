package helper

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"go-contracts/contracts/erc20"
	"go-contracts/contracts/univ3/router"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	GasLimit     = 900000
	MoreGasPrice = 100000
	UniswapFee   = 100 // Uniswap V3 的 fee tier (0.01%)
)

// ExactInputSingle 两个加密货币执行Uniswap V3 的兑换操作
func ExactInputSingle(ctx context.Context, client *ethclient.Client, privateKey, routerAddress, tokenIn, tokenOut string, amountIn *big.Int) (string, error) {
	chainID, _ := client.ChainID(ctx)

	_, err := GetTokenTransferApproval(ctx, client, privateKey, tokenIn, routerAddress, amountIn)
	if err != nil {
		return "", err
	}

	// 将私钥导入
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	// 创建授权的交易签署者
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	gasPrice = gasPrice.Add(gasPrice, big.NewInt(MoreGasPrice))
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	// 设置交易选项
	txOpts, _ := bind.NewKeyedTransactorWithChainID(key, chainID)
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = uint64(GasLimit) // 设置 gas 限额
	txOpts.GasPrice = gasPrice

	// 与 Uniswap V3 路由器交互执行兑换
	route, err := router.NewRouterTransactor(common.HexToAddress(routerAddress), client)
	if err != nil {
		return "", err
	}

	params := router.ISwapRouterExactInputSingleParams{
		TokenIn:           common.HexToAddress(tokenIn),
		TokenOut:          common.HexToAddress(tokenOut),
		Fee:               big.NewInt(UniswapFee), // Uniswap V3 的 fee
		Recipient:         fromAddress,
		Deadline:          big.NewInt(time.Now().Add(time.Minute * 15).Unix()), // 15 分钟超时
		AmountIn:          amountIn,
		AmountOutMinimum:  big.NewInt(0), // 设置最小接受输出值
		SqrtPriceLimitX96: big.NewInt(0), // 不限制价格滑点
	}

	tx, err := route.ExactInputSingle(
		txOpts,
		params,
	)
	if err != nil {
		return "", err
	}

	fmt.Printf("Transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

// ExactInput 两个加密货币通过path执行Uniswap V3的兑换操作
func ExactInput(ctx context.Context, client *ethclient.Client, privateKey, routerAddress, tokenIn, tokenOut string, amountIn *big.Int) (string, error) {
	chainID, _ := client.ChainID(ctx)

	// 获取WETH
	midToken, err := GetMidToken(client, routerAddress)
	if err != nil {
		return "", err
	}

	// 获取报价和路径
	quote, path, err := QuoteExactInput(ctx, client, QuoterV2Address, tokenIn, midToken, tokenOut, amountIn)
	if err != nil {
		return "", err
	}
	if len(path) == 0 {
		return "", fmt.Errorf("path is nil")
	}

	quoteAmountD := decimal.NewFromBigInt(quote, -ARBDecimals)
	fmt.Println("quote:", quoteAmountD.String())
	fmt.Println("path:", common.Bytes2Hex(path))

	// 授权代币交易
	_, err = GetTokenTransferApproval(ctx, client, privateKey, tokenIn, routerAddress, amountIn)
	if err != nil {
		return "", err
	}

	// 将私钥导入
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	// 创建授权的交易签署者
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	gasPrice = gasPrice.Add(gasPrice, big.NewInt(MoreGasPrice))
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	// 设置交易选项
	txOpts, _ := bind.NewKeyedTransactorWithChainID(key, chainID)
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = uint64(GasLimit) // 设置 gas 限额
	txOpts.GasPrice = gasPrice

	// 与 Uniswap V3 路由器交互执行兑换
	route, err := router.NewRouterTransactor(common.HexToAddress(routerAddress), client)
	if err != nil {
		return "", err
	}

	params := router.ISwapRouterExactInputParams{
		Path:             path,
		Recipient:        fromAddress,
		Deadline:         big.NewInt(time.Now().Add(time.Minute * 15).Unix()), // 15 分钟超时
		AmountIn:         amountIn,
		AmountOutMinimum: big.NewInt(0), // 设置最小接受输出值
	}

	tx, err := route.ExactInput(
		txOpts,
		params,
	)
	if err != nil {
		return "", err
	}

	fmt.Println("Transaction sent: ", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func GetMidToken(client *ethclient.Client, routerAddress string) (string, error) {
	route, err := router.NewRouterCaller(common.HexToAddress(routerAddress), client)
	if err != nil {
		return "", err
	}

	weth, err := route.WETH9(&bind.CallOpts{})
	if err != nil {
		return "", err
	}

	return weth.String(), nil
}

// GetTokenTransferApproval 封装 ERC20 代币的 approve 操作
func GetTokenTransferApproval(ctx context.Context, client *ethclient.Client, privateKey string, tokenAddress string, spenderAddress string, amount *big.Int) (string, error) {
	chainID, _ := client.ChainID(ctx)
	// 将私钥转化为 ECDSA 格式
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	// 获取钱包地址
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	// 获取当前 nonce 值
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	// 获取建议的 gas 价格
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	gasPrice = gasPrice.Add(gasPrice, big.NewInt(MoreGasPrice))
	// 创建授权交易签名者
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // 代币授权操作不需要 ETH 直接支付
	auth.GasLimit = uint64(GasLimit) // 设置 gas 限制
	auth.GasPrice = gasPrice

	// 初始化 ERC20 合约实例
	tokenContract, err := erc20.NewErc20(common.HexToAddress(tokenAddress), client)
	if err != nil {
		return "", err
	}

	// 调用 ERC20 合约的 approve 函数
	spender := common.HexToAddress(spenderAddress)
	tx, err := tokenContract.Approve(auth, spender, amount)
	if err != nil {
		return "", err
	}

	fmt.Println("approvalHash:", tx.Hash().Hex())
	// 返回交易哈希
	return tx.Hash().Hex(), nil
}
