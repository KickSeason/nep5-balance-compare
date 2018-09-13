package neoapi

//Nep5BalanceReq the request to get nep5 balance
type Nep5BalanceReq struct {
	ID      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//Nep5BalanceResp nep5 balance response
type Nep5BalanceResp struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		GasConsumed string `json:"gas_consumed"`
		Script      string `json:"script"`
		Stack       []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"stack"`
		State string `json:"state"`
	} `json:"result"`
}

// //StackItem of nep5 balance response
// type StackItem struct {
// 	Type  string
// 	Value string
// }

// //Nep5BalanceResult the rpc result
// type Nep5BalanceResult struct {
// 	Script       string
// 	State        string
// 	Gas_consumed string
// 	Stack        []StackItem
// 	Tx           string
// }

// //Nep5BalanceResp the response
// type Nep5BalanceResp struct {
// 	Jsonrpc string
// 	Method  string
// 	Result  Nep5BalanceResult
// }
