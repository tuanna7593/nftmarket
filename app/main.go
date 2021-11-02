package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	"github.com/tuanna7593/nftmarketplace/client"
	"github.com/tuanna7593/nftmarketplace/cmd"
	"github.com/tuanna7593/nftmarketplace/config"
)

type MarketItemSold struct {
	Seller common.Address
	Owner  common.Address
	Price  *big.Int
}

func main() {
	// read config
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.InitConfig(&cfg)

	// init ether client
	err = initEthClient()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() (cfg config.ClientConfig, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

func initEthClient() error {
	cfg := config.GetConfig()
	if cfg.RPCEndpoint == "" {
		return fmt.Errorf("not found rpc endpoint")
	}

	cli, err := ethclient.Dial(cfg.RPCEndpoint)
	if err != nil {
		return err
	}

	client.SetEtheClient(cli)
	return nil
}
