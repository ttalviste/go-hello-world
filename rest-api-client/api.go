package nftport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type ContractNFTsResponse struct {
	Response string `json:"response"`
	Total    int    `json:"total"`
	NFTS     []NFT  `json:"nfts"`
}

// A Pokemon Struct to map every pokemon to.
type NFT struct {
	Contract int `json:"contract_address"`
	TokenId  int `json:"token_id"`
}

func FetchSandboxTokens() {
	LoadNFTs("0x50f5474724e0ee42d9a4e711ccfb275809fd6d4a")
	LoadNFTs("0x5cc5b05a8a13e3fbdb0bb9fccd98d38e50f90c38")
}

func LoadNFTs(contract string) {

	url := "https://api.nftport.xyz/v0/nfts/" + contract + "?chain=ethereum"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(body))

	var nftsResponseObject ContractNFTsResponse
	json.Unmarshal(body, &nftsResponseObject)

	fmt.Println(nftsResponseObject.Response)
	fmt.Println(len(nftsResponseObject.NFTS))
	fmt.Println(nftsResponseObject.Total)

}
