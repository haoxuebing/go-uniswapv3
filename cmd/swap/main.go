package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"go-contracts/helper"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const (
	USDTAddress = "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9" // USDT 合约地址
	USDCAddress = "0xaf88d065e77c8cC2239327C5EDb3A432268e5831" // USDC 合约地址
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 读取环境变量
	privateKey := os.Getenv("PRIVATE_KEY")
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	rawURL := fmt.Sprintf("https://arbitrum-mainnet.infura.io/v3/%s", infuraProjectID)
	client, err := ethclient.Dial(rawURL)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// 设置交易值
	amountIn := big.NewInt(100000) // 假设兑换 0.1 USDT（USDT 精度为6位）
	// _, err = helper.ExactInputSingle(ctx, client, privateKey, helper.UniswapV3RouterAddress, helper.USDT, helper.USDC, amountIn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err = helper.ExactInput(ctx, client, privateKey, helper.UniswapV3RouterAddress, helper.USDC, helper.ARB, amountIn)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Transaction hash:", txHash)
}
