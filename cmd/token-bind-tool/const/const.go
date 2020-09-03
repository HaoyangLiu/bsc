package _const

const (
	Passwd              = "12345678"
	NetworkType         = "network-type"
	KeystorePath        = "keystore-path"
	ConfigPath          = "config-path"
	Operation           = "operation"
	BEP20ContractAddr   = "bep20-contract-addr"
	BEP2Symbol          = "bep2-symbol"
	Recipient           = "recipient"
	PeggyAmount         = "peggy-amount"
	LedgerAccountIndex  = "ledger-account-index"

	InitKey               = "initKey"
	RefundRestBNB         = "refundRestBNB"
	DeployContract        = "deployContract"
	ApproveBind           = "approveBindAndTransferOwnership"
	DeployTransferRefund  = "deploy_transferTokenAndOwnership_refund"
	ApproveBindFromLedger = "approveBindFromLedger"

	Mainnet = "mainnet"
	TestNet = "testnet"

	BindKeystore = "bind_keystore"

	TestnetRPC     = "https://data-seed-prebsc-1-s1.binance.org:8545"
	TestnetChainID = 97

	MainnnetRPC    = "https://bsc-dataseed1.binance.org:443"
	MainnetChainID = 56

	OneBNB          = 1e18
	DefaultGasPrice = 20000000000
	DefaultGasLimit = 4700000

	MainnetExplorerTxUrl = "%s: https://bscscan.com/tx/%s"
	TestnetExplorerTxUrl = "%s: https://testnet.bscscan.com/tx/%s"

	MainnetExplorerAddressUrl = "%s: https://bscscan.com/address/%s"
	TestnetExplorerAddressUrl = "%s: https://testnet.bscscan.com/address/%s"

	BSCAddrLength = 42
)
