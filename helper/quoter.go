package helper

import (
	"bytes"
	"context"
	"encoding/hex"
	"go-contracts/contracts/univ3/quoterV2"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QuoteExactInputSingle(ctx context.Context, client *ethclient.Client, quoterV2Address, inToken, outToken string, amountIn *big.Int) (amountOut *big.Int, err error) {
	quoter, err := quoterV2.NewQuoterV2Caller(common.HexToAddress(quoterV2Address), client)
	if err != nil {
		log.Fatal(err)
	}

	params := quoterV2.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           common.HexToAddress(inToken),
		TokenOut:          common.HexToAddress(outToken),
		AmountIn:          amountIn,
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	return quoter.QuoteExactInputSingle(&bind.CallOpts{}, params)
}

func QuoteExactOutputSingle(ctx context.Context, client *ethclient.Client, quoterV2Address, inToken, outToken string, amountOut *big.Int) (amountIn *big.Int, err error) {
	quoter, err := quoterV2.NewQuoterV2Caller(common.HexToAddress(quoterV2Address), client)
	if err != nil {
		log.Fatal(err)
	}

	params := quoterV2.IQuoterV2QuoteExactOutputSingleParams{
		TokenIn:           common.HexToAddress(inToken),
		TokenOut:          common.HexToAddress(outToken),
		Amount:            amountOut,
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	return quoter.QuoteExactOutputSingle(&bind.CallOpts{}, params)
}

func QuoteExactInput(ctx context.Context, client *ethclient.Client, quoterV2Address, inToken, midToken, outToken string, amountIn *big.Int) (amountOut *big.Int, path []byte, err error) {
	contractAbi, err := abi.JSON(strings.NewReader(quoterV2.QuoterV2MetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	path = formatQuoterPath(inToken, midToken, outToken)
	method := contractAbi.Methods["quoteExactInput"]
	calldata, err := method.Inputs.Pack(path, amountIn)
	if err != nil {
		log.Fatal(err)
	}
	to := common.HexToAddress(quoterV2Address)
	calldata = append(method.ID, calldata...)
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &to,
		Data: calldata,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	resM := map[string]interface{}{}
	err = method.Outputs.UnpackIntoMap(resM, result)
	if err != nil {
		log.Fatal(err)
	}
	amountOutV, ok := resM["amountOut"]
	if !ok {
		log.Fatal("amountOut not found in result")
	}
	amountOut, ok = amountOutV.(*big.Int)
	if !ok {
		log.Fatal("amountOut is not big.Int")
	}
	return
}

func QuoteExactOutput(ctx context.Context, client *ethclient.Client, quoterV2Address, inToken, midToken, outToken string, amountOut *big.Int) (amountIn *big.Int, path []byte, err error) {
	quoter, err := quoterV2.NewQuoterV2Caller(common.HexToAddress(quoterV2Address), client)
	if err != nil {
		log.Fatal(err)
	}
	path = formatQuoterPath(inToken, midToken, outToken)

	amountIn, err = quoter.QuoteExactOutput(&bind.CallOpts{}, path, amountOut)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func formatQuoterPath(inToken, midToken, outToken string) (path []byte) {
	path = buildPath([]string{inToken, midToken, outToken}, []*big.Int{big.NewInt(500), big.NewInt(500)})
	return
}

// hexToBytes converts a hex string to a byte slice
func hexToBytes(str string) []byte {
	data, _ := hex.DecodeString(str[2:]) // Remove "0x" prefix
	return data
}

// encodeFee converts a fee into a 3-byte representation
func encodeFee(fee *big.Int) []byte {
	feeBytes := make([]byte, 3)
	feeBytes[0] = byte(fee.Uint64() >> 16)
	feeBytes[1] = byte(fee.Uint64() >> 8)
	feeBytes[2] = byte(fee.Uint64())
	return feeBytes
}

func buildPath(tokens []string, fees []*big.Int) []byte {
	var path bytes.Buffer
	for i := 0; i < len(tokens)-1; i++ {
		// Append token address (20 bytes)
		path.Write(hexToBytes(tokens[i]))
		// Append fee (3 bytes)
		path.Write(encodeFee(fees[i]))
	}
	// Append the last token address (20 bytes)
	path.Write(hexToBytes(tokens[len(tokens)-1]))

	return path.Bytes()
}
