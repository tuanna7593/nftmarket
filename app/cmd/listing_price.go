package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	ethClient "github.com/tuanna7593/nftmarketplace/client"
	"github.com/tuanna7593/nftmarketplace/config"
	"github.com/tuanna7593/nftmarketplace/contracts"
	"github.com/tuanna7593/nftmarketplace/utils"
)

func init() {
	rootCmd.AddCommand(listingPriceCmd)
}

var listingPriceCmd = &cobra.Command{
	Use:   "listingPrice",
	Short: "Print the price for listing an nft in the market",
	Long:  `Print the price for listing an nft in the market`,
	Run: func(cmd *cobra.Command, args []string) {
		runListingPriceCommand()
	},
}

func runListingPriceCommand() {
	cfg := config.GetConfig()
	client := ethClient.GetEtheClient()

	marketAddress := common.HexToAddress(cfg.NFTMarketAddress)
	nftMarket, err := contracts.NewNFTMarket(marketAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	listingPrice, err := nftMarket.GetListingPrice(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listing Price: %s eth \n", utils.ConvertWeiToEth(listingPrice).FloatString(2))
}
