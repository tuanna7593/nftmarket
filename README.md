# Prerequisite
- Copy `.env.dev` to `.env`
- Update the `RPC_ENDPOINT` and `SECRET` value in `.env` file

# Deploy contracts
- Run `$npx hardhat run scripts/deploy.js --network ropsten`
- Update `NFT_MARKET_ADDRESS` and `NFT_ADDRESS` value in `.env` file with the value response from the above command

# Verify contracts
- Run `npx hardhat verify --network ropsten {{CONTRACT_ADDRESS}} "{{PARAMETERS}}"`

# Upgrade contracts
- Run `$npx hardhat run scripts/upgrade.js --network ropsten`

# Build CLI app
- Run `make build-cli`

# Commands in CLI
- Run `./nftmarket totalSell` to get total sell amount of the nft market.
- Run `./nftmarket listingPrice` to get the price for listing an nft in the market.
