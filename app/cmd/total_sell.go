package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	ethClient "github.com/tuanna7593/nftmarketplace/client"
	"github.com/tuanna7593/nftmarketplace/config"
	"github.com/tuanna7593/nftmarketplace/contracts"
	"github.com/tuanna7593/nftmarketplace/utils"
)

func init() {
	rootCmd.AddCommand(totalSellCmd)
}

var totalSellCmd = &cobra.Command{
	Use:   "totalSell",
	Short: "Print the total sell amount in NFT Market",
	Long:  `Print the total sell amount in NFT Market`,
	Run: func(cmd *cobra.Command, args []string) {
		runTotalSellCommand()
	},
}

type MarketItemSold struct {
	Seller common.Address
	Owner  common.Address
	Price  *big.Int
}

func runTotalSellCommand() {
	cfg := config.GetConfig()
	client := ethClient.GetEtheClient()

	marketContractAbi, err := abi.JSON(strings.NewReader(string(contracts.NFTMarketABI)))
	if err != nil {
		log.Fatal(err)
	}

	marketAddress := common.HexToAddress(cfg.NFTMarketAddress)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{
			marketAddress,
		},
		Topics: [][]common.Hash{
			{contracts.MarketItemSoldSignatureHash},
		},
	})
	if err != nil {
		log.Fatalf("failed to get logs from nft markets:%v", err)
	}

	totalSoldAmount := big.NewInt(0)
	for i := range logs {
		var itemSold MarketItemSold

		err := marketContractAbi.UnpackIntoInterface(&itemSold, "MarketItemSold", logs[i].Data)
		if err != nil {
			log.Fatal(err)
		}

		totalSoldAmount.Add(totalSoldAmount, itemSold.Price)
	}

	fmt.Printf("Total Sold Amount: %s eth\n", utils.ConvertWeiToEth(totalSoldAmount).FloatString(2))
}
