package nelapi

//get block count
type BlockCountValue struct {
	Blockcount string
}
type BlockCountResp struct {
	Jsonrpc string
	Id      int
	Result  []BlockCountValue
}

//Nep5TransferReq get block nep5 transactions
type nep5TransferReq struct {
	Jsonrpc string
	Method  string
	Params  []int
	ID      int
}
type Transaction struct {
	Blockindex int
	Txid       string
	N          int
	Asset      string
	From       string
	To         string
	Value      string
}
type Nep5TransferResp struct {
	Jsonrpc string
	Id      int
	Result  []Transaction
}
