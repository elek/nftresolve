package main

import (
	"encoding/json"
	"fmt"
	erc1155meta "github.com/elek/nftresolve/erc1155"
	erc721meta "github.com/elek/nftresolve/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

var chains = map[string]string{
	"avax":  "https://api.avax.network/ext/bc/C/rpc",
	"bsc":   "https://bsc-dataseed.binance.org/",
	"matic": "https://rpc-mainnet.maticvigil.com/",
}

func main() {

	app := &cli.App{
		Name:      "nftresolve",
		Usage:     "Display metadata of NFT token instances",
		ArgsUsage: "<contract address(hex)> <token id (dec)>",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "chain",
				Usage:   "eth-rpc compatible API endpoint or chain alias (avac, bsc, matic)",
				EnvVars: []string{"NFTRESOLVE_CHAIN"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return fmt.Errorf("contract id (hex) and token id (dec) are required as arguments ")
			}
			return resolve(c.String("chain"), c.Args().Get(0), c.Args().Get(1))
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func resolveChain(chain string) string {
	if url, found := chains[chain]; found {
		return url
	}
	return chain
}
func resolve(chain string, contract string, nftId string) error {
	client, err := ethclient.Dial(resolveChain(chain))

	if err != nil {
		return errors.Wrap(err, "Couldn't connect to the eth-rpc")
	}

	id := big.NewInt(0)
	id, ok := id.SetString(nftId, 10)
	if !ok {
		return fmt.Errorf("Couldn't set bigint")
	}

	address := common.HexToAddress(contract)
	uri, err := erc721GetUri(address, client, id)
	if err != nil {
		uri, err = erc1155GetUri(address, client, id)
		if err != nil {
			return errors.Wrap(err, "Neither erc721 nor erc1155 contract couldn't be called")
		}
	}

	fmt.Printf("URI: %s\n", uri)
	resolvedUrl := strings.ReplaceAll(uri, "0x{id}", "0x"+id.Text(16))
	resolvedUrl = strings.ReplaceAll(resolvedUrl, "{id}", nftId)
	if uri != resolvedUrl {
		fmt.Printf("Resolved URI: %s\n", resolvedUrl)
	}
	if strings.HasPrefix(resolvedUrl, "http") {
		req, err := http.NewRequest("GET", resolvedUrl, nil)
		if err != nil {
			return errors.Wrap(err, "Couldn't create HTTP get request for "+resolvedUrl)
		}
		req.Header.Set("Accept", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "Couldn't retrieve url "+resolvedUrl)
		}
		content, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		parsedJson := make(map[string]interface{})
		err = json.Unmarshal(content, &parsedJson)
		if err != nil {
			fmt.Println(string(content))
		} else {
			idented, err := json.MarshalIndent(parsedJson, "", "   ")
			if err == nil {
				fmt.Println(string(idented))
			} else {
				fmt.Println(idented)
			}
		}
	}
	return nil
}

func erc721GetUri(address common.Address, client *ethclient.Client, id *big.Int) (string, error) {
	erc721, err := erc721meta.NewErc721meta(address, client)
	if err != nil {
		errors.Wrap(err, "Couldn't initialize contract")
	}

	uri, err := erc721.TokenURI(&bind.CallOpts{}, id)
	if err != nil {
		return "", err
	}
	return uri, nil
}

func erc1155GetUri(address common.Address, client *ethclient.Client, id *big.Int) (string, error) {
	contract, err := erc1155meta.NewErc1155meta(address, client)
	if err != nil {
		errors.Wrap(err, "Couldn't initialize contract")
	}

	uri, err := contract.Uri(&bind.CallOpts{}, id)
	if err != nil {
		return "", err
	}
	return uri, nil
}
