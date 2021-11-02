package contracts

import "github.com/ethereum/go-ethereum/crypto"

var (
	MarketItemSoldSignature     = []byte("MarketItemSold(uint256,address,uint256,address,address,uint256)")
	MarketItemSoldSignatureHash = crypto.Keccak256Hash(MarketItemSoldSignature)
)
