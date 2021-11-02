package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nftmarket",
	Short: "NFTMarket is a command to get info from NFT Market",
	Long:  `NFTMarket support to get total sell amount from NFT market`,
}

func Execute() error {
	return rootCmd.Execute()
}
