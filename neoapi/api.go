package neoapi

import (
	"encoding/json"
	"fmt"
	"log"
	"nep5-balance-compare/http"
	"nep5-balance-compare/utils"
	"strings"
)

//GetAddressNep5Balance get the balance of the asset
func GetAddressNep5Balance(server string, address string, assetID string) (Nep5BalanceResp, error) {
	log.Println(address, assetID, server)
	hash, err := utils.ToHash160(address)
	if err != nil {
		fmt.Println(err.Error())
		return Nep5BalanceResp{}, err
	}
	reader := strings.NewReader(`{
		"jsonrpc": "2.0",
		"method": "invokefunction",
		"params": [
		  "` + assetID + `",
		  "balanceOf",
		  [
			{
			  "type": "Hash160",
			  "value": "0x` + hash + `"
			}
		  ]
		],
		"id": 3
	  }`)

	respBytes, err := http.Request("POST", server, reader)
	if err != nil {
		fmt.Print(err.Error())
		return Nep5BalanceResp{}, err
	}
	var resp Nep5BalanceResp
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		fmt.Println(err)
		return Nep5BalanceResp{}, err
	}
	return resp, nil
}
