// scripts/upgrade-box.js
require('dotenv').config();
const { ethers, upgrades } = require("hardhat");

const nftMarketAddress = process.env.NFT_MARKET_ADDRESS

async function main() {
  if (nftMarketAddress == "") {
    console.log("not found the NFTMarket address to upgrade")
    return
  }
  
  const NFTMarketV3Fac = await ethers.getContractFactory("NFTMarketV3");
  await upgrades.upgradeProxy(nftMarketAddress, NFTMarketV3Fac);
  console.log("NFTMarket upgraded");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
