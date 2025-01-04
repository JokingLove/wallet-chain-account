package ethereum

type Eip1559DynamicFeeTx struct {
	ChainId              string `json:"chain_id"`
	Nonce                uint64 `json:"nonce"`
	FromAddress          string `json:"from_address"`
	ToAddress            string `json:"to_address"`
	GasLimit             uint64 `json:"gas_limit"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	Amount               string `json:"amount"`
	ContractAddress      string `json:"contract_address"`
}

type TxStructure struct {
	ChainId         string `json:"chain_id"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Amount          int64  `json:"amount"`
	ContractAddress string `json:"contract_address"`
	GasLimit        uint64 `json:"gas_limit"`
	FeeAmount       int64  `json:"fee_amount"`
	Memo            string `json:"memo"`
	Decimal         int    `json:"decimal"`
	Sequence        uint64 `json:"sequence"`
	AccountNumber   uint64 `json:"account_number"`
	PubKey          string `json:"pub_key"`
}
