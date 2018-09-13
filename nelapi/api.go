package nelapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"nep5-balance-compare/http"
	"strconv"
	"strings"
)

//GetBlockCount get block count
func GetBlockCount() int {
	reader := strings.NewReader(`{
		"jsonrpc": "2.0",
		"method": "getblockcount",
		"params": [],
		"id": "1"
	}`)

	respBytes, err := http.Request("POST", "https://api.nel.group/api/mainnet", reader)
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}
	var resp BlockCountResp
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	result, err := strconv.Atoi(resp.Result[0].Blockcount)
	if err != nil {
		return -1
	}
	return result
}

//GetNep5TransferByBlockIndex get nep5 transfers by block index
func GetNep5TransferByBlockIndex(blocks []int) (Nep5TransferResp, error) {
	body := nep5TransferReq{
		Jsonrpc: "2.0",
		Method:  "getnep5transferbyblockindex",
		Params:  blocks,
		ID:      1,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return Nep5TransferResp{}, err
	}
	reader := bytes.NewReader(data)

	respBytes, err := http.Request("POST", "https://api.nel.group/api/mainnet", reader)

	if err != nil {
		fmt.Println(err.Error())
		return Nep5TransferResp{}, err
	}
	var resp Nep5TransferResp
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		fmt.Println(err.Error())
		return Nep5TransferResp{}, err
	}
	return resp, nil
}
