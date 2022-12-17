// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const {ethers} = require('hardhat');
const namehash = require('eth-ens-namehash')
async function printTransacHistory(bank, idx) {
  let t = await bank.transArray(idx);
  console.log(t.user, t.trans_type, t.amount, t.timestamp);

}
async function main() {

  let [owner, user1, user2, ...users] = await hre.ethers.getSigners()
  // We get the contract to deploy
  const OmniOneNFT = await hre.ethers.getContractFactory("OmniOneNFT");
  const ooNFT = await OmniOneNFT.deploy(owner.address);

  await ooNFT.deployed();
  console.log("OmniOneNFT deployed to:", ooNFT.address);
  //  event SyncMinted(address from, address to, uint256 id);
  ooNFT.on("SyncMinted", (setter, from, to,id, event) => {
    console.log("sync minted id", id, " from ", from, ' to ', to,);
    console.log(event);
  })
  // event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
  ooNFT.on("Transfer", (from, to, id) => {
    console.log("Transder id", id, " from ", from, ' to ', to,);
  })
  // User0 apply to get an NFT
  let id = 0;
  let random = ethers.utils.randomBytes(32);
  let messageHash = ethers.utils.solidityKeccak256(
    ['address', 'address', "uint256","bytes32"],
    [user1.address, user1.address, id,random]);

  console.log("bytesMsg", messageHash);
  let flatSigned = await owner.signMessage(ethers.utils.arrayify(messageHash));
  let tx = await ooNFT.connect(user1).syncMint(user1.address, id,random, flatSigned);
  const receipt = await tx.wait()

  for (const event of receipt.events) {
    console.log(`Event ${event.event} with args ${event.args}`);
  }
  // Cross transfer


  return
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
