package main

import (
	"context"

	"go-contracts/helper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var (
	WETH = "0x82af49447d8a07e3bd95bd0d56f35241523fbab1"
	USDC = "0xaf88d065e77c8cC2239327C5EDb3A432268e5831"
	ARB  = "0x912CE59144191C1204E64559FE8253a0e49E6548"
)

func main() {
	rpcUrl := "https://arb1.arbitrum.io/rpc"
	client, _ := ethclient.Dial(rpcUrl)

	ctx := context.Background()
	amountInParam := decimal.NewFromInt(1)
	amountIn := amountInParam.Mul(decimal.New(1, helper.USDCDecimals)).BigInt()
	amountOutParam := decimal.NewFromInt(1)
	amountOut := amountOutParam.Mul(decimal.New(1, helper.ARBDecimals)).BigInt()

	// QuoteExactInputSingle
	amountQuote, err := helper.QuoteExactInputSingle(ctx, client, helper.QuoterV2Address, USDC, ARB, amountIn)
	if err != nil {
		panic(err)
	}
	quoteAmountD := decimal.NewFromBigInt(amountQuote, -helper.ARBDecimals)
	println(quoteAmountD.String())

	// QuoteExactOutputSingle
	amountQuote2, err := helper.QuoteExactOutputSingle(ctx, client, helper.QuoterV2Address, USDC, ARB, amountOut)
	if err != nil {
		panic(err)
	}
	quoteAmountD2 := decimal.NewFromBigInt(amountQuote2, -helper.USDCDecimals)
	println(quoteAmountD2.String())

	// QuoteExactInput
	amountQuote3, path, err := helper.QuoteExactInput(ctx, client, helper.QuoterV2Address, USDC, WETH, ARB, amountIn)
	if err != nil {
		panic(err)
	}
	quoteAmountD3 := decimal.NewFromBigInt(amountQuote3, -helper.ARBDecimals)
	println(quoteAmountD3.String())
	println(common.Bytes2Hex(path))

	// QuoteExactOutput
	amountQuote4, path, err := helper.QuoteExactOutput(ctx, client, helper.QuoterV2Address, USDC, WETH, ARB, amountIn)
	if err != nil {
		panic(err)
	}
	quoteAmountD4 := decimal.NewFromBigInt(amountQuote4, -helper.ARBDecimals)
	println(quoteAmountD4.String())
	println(common.Bytes2Hex(path))
}
