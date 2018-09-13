package main

import (
	"errors"
	"log"
	"nep5-balance-compare/nelapi"
	"nep5-balance-compare/neoapi"
	"nep5-balance-compare/utils"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const (
	BlocksPerGroup   = 20
	DefaultCaseCount = 1000
)

type testCase struct {
	addr  string
	asset string
}

var neoSrv1 string
var neoSrv2 string

var addressAssets = make(map[string]testCase)
var testCaseCnt int
var caseCount int
var blockCount int
var unequalCases = []testCase{}

func handleNep5TransferResp(resp nelapi.Nep5TransferResp) {
	for _, tx := range resp.Result {
		addAddress(tx)
	}
}
func addAddress(tx nelapi.Transaction) {
	key := tx.To + tx.Asset
	if _, ok := addressAssets[key]; !ok {
		addressAssets[key] = testCase{
			addr:  tx.To,
			asset: tx.Asset,
		}
		caseCount++
		log.Printf("get case %d: address: %s, assetid: %s\n", caseCount, tx.To, tx.Asset)
	}
}

func getAddressAndAssetID() {
	i := 0
	for i < blockCount {
		params := []int{}
		for j := 0; j < BlocksPerGroup; j++ {
			params = append(params, blockCount-j-10*i)
		}
		resp, err := nelapi.GetNep5TransferByBlockIndex(params)
		if err != nil {
			log.Println(err.Error())
			return
		}
		handleNep5TransferResp(resp)
		i += 10
		if testCaseCnt < caseCount {
			return
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "nep5-compare"
	app.Version = "1.0.0"
	app.Usage = "compare nep5 balance between two version-different neo-cli"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "base, b",
			Usage:       "The base server, usually 2.7.6",
			Value:       "http://47.254.44.88:10332",
			Destination: &neoSrv1,
		},
		cli.StringFlag{
			Name:        "compare, c",
			Usage:       "The compare server, usually new version neo-cli. no default, must input",
			Value:       "",
			Destination: &neoSrv2,
		},
		cli.IntFlag{
			Name:        "number, n",
			Usage:       "How many address-assetID pairs to test",
			Value:       DefaultCaseCount,
			Destination: &testCaseCnt,
		},
	}
	app.Action = func(c *cli.Context) error {
		if neoSrv2 == "" {
			return errors.New("no compare server, please input the new to-compare server")
		}
		if testCaseCnt <= 0 {
			return errors.New("please input a valid number")
		}
		blockCount = nelapi.GetBlockCount()
		if blockCount < 0 {
			return errors.New("cannt get current block height")
		}
		log.Printf("current block count: %d\n", blockCount)
		log.Printf("start to get at least %d cases.\n", testCaseCnt)
		tStart := time.Now()
		getAddressAndAssetID()
		tGetaddrs := time.Now()
		log.Printf("got %+v cases, start to compare...\n", caseCount)
		for _, c := range addressAssets {
			addr := c.addr
			asset := c.asset
			ch1 := make(chan neoapi.Nep5BalanceResp)
			ch2 := make(chan neoapi.Nep5BalanceResp)

			go func() {
				resp, err := neoapi.GetAddressNep5Balance(neoSrv1, addr, asset)
				if err != nil {
					log.Println(err)
					close(ch1)
					return
				}
				ch1 <- resp
			}()
			go func() {
				resp, err := neoapi.GetAddressNep5Balance(neoSrv2, addr, asset)
				if err != nil {
					log.Println(err)
					close(ch2)
					return
				}
				ch2 <- resp
			}()
			resp1, ok1 := <-ch1
			resp2, ok2 := <-ch2
			if ok1 && ok2 && !strings.Contains(resp1.Result.State, "FAULT") && !strings.Contains(resp2.Result.State, "FAULT") {
				value1 := resp1.Result.Stack[0].Value
				value2 := resp2.Result.Stack[0].Value
				var result string
				if value1 == value2 {
					result = "Equal"
				} else {
					result = "Unequal"
					unequalCases = append(unequalCases, testCase{addr: addr, asset: asset})
				}
				log.Printf("[%s] addr: %s, asset: %s, neo3.0: %s neo2.7.6: %s\n", result, addr, asset, value1, value2)
			}

		}
		tFinish := time.Now()
		log.Printf("total test: %d, unequal: %d, get cases cost: %v, compare cost: %v\n", caseCount, len(unequalCases), tGetaddrs.Sub(tStart), tFinish.Sub(tGetaddrs))
		for _, value := range unequalCases {
			hash160, err := utils.ToHash160(value.addr)
			if err != nil {
				continue
			}
			log.Printf("[unequal]%+v， address-hash160: %s\n", value, hash160)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
