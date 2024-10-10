package contracts

import "github.com/ethereum/go-ethereum/common"

const (
	ETH_CHAIN_ID   uint = 1
	OP_CHAIN_ID    uint = 10
	BSC_CHAIN_ID   uint = 56
	MATIC_CHAIN_ID uint = 137
	ARB_CHAIN_ID   uint = 42161
)

var QuoterVersion = map[uint]int{
	ETH_CHAIN_ID:   1,
	OP_CHAIN_ID:    1,
	BSC_CHAIN_ID:   3,
	MATIC_CHAIN_ID: 1,
	ARB_CHAIN_ID:   1,
}

var QuoterAddr = map[uint]map[string]common.Address{
	ETH_CHAIN_ID:   {"univ3": UniV3QuoterV2Addr},
	OP_CHAIN_ID:    {"univ3": UniV3QuoterV2Addr},
	BSC_CHAIN_ID:   {"pancakev3": PancakeV3QuoterAddr},
	MATIC_CHAIN_ID: {"univ3": UniV3QuoterV2Addr, "quickv3": QuickSwapV3QuoterAddr},
	ARB_CHAIN_ID:   {"univ3": UniV3QuoterV2Addr, "camelotv3": CamelotV3QuoterAddr, "woo": WooRouterAddr},
}

var (
	UniV3QuoterV1Addr     = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
	UniV3QuoterV2Addr     = common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e")
	UniV3QuoterV2BSCAddr  = common.HexToAddress("0x78D78E420Da98ad378D7799bE8f4AF69033EB077")
	PancakeV3QuoterAddr   = common.HexToAddress("0xB048Bbc1Ee6b733FFfCFb9e9CeF7375518e25997")
	CamelotV3QuoterAddr   = common.HexToAddress("0x0Fc73040b26E9bC8514fA028D998E73A254Fa76E")
	WooRouterAddr         = common.HexToAddress("0x9aEd3A8896A85FE9a8CAc52C9B402D092B629a30")
	QuickSwapV3QuoterAddr = common.HexToAddress("0xa15F0D7377B2A0C0c10db057f641beD21028FC89")
)

var TokenAddrMap = map[uint]map[string]common.Address{
	ETH_CHAIN_ID: {
		"WETH": common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		"USDC": common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		"USDT": common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		"WBTC": common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
		"UMA":  common.HexToAddress("0x04Fa0d235C4abf4BcF4787aF4CF447DE572eF828"),
	},
	OP_CHAIN_ID: {
		"WETH": common.HexToAddress("0x4200000000000000000000000000000000000006"),
		"USDC": common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
	},
	BSC_CHAIN_ID: {
		"WBNB": common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		"BUSD": common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"),
		"USDC": common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
		"USDT": common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
		"WETH": common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
		"WBTC": common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c"),
		"CAKE": common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"),
	},
	ARB_CHAIN_ID: {
		"WETH":  common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"),
		"USDCe": common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"),
		"WBTC":  common.HexToAddress("0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f"),
		"GMX":   common.HexToAddress("0xfc5A1A6EB076a2C7aD06eD22C90d7E710E35ad0a"),
		"ARB":   common.HexToAddress("0x912CE59144191C1204E64559FE8253a0e49E6548"),
		"USDT":  common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
		"USDC":  common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
	},
	MATIC_CHAIN_ID: {
		"USDCe":  common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
		"USDC":   common.HexToAddress("0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"),
		"WETH":   common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"),
		"WMATIC": common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"),
	},
}
