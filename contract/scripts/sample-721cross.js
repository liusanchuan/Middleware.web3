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
let owner, user1, user2, users;

async function applySyncMint(userAddr) {

	let id = 0;
	let random = ethers.utils.randomBytes(32);
	let messageHash = ethers.utils.solidityKeccak256(
		['address', 'address', "uint256", "bytes32"],
		[user1.address, user1.address, id, random]);

	console.log("bytesMsg", messageHash);
	let flatSigned = await owner.signMessage(ethers.utils.arrayify(messageHash));
	return [id, random, flatSigned]
}
async function applyCrossTransfer(sender, receiver, id, receiveChainId){

	let random = ethers.utils.randomBytes(32);
	let messageHash = ethers.utils.solidityKeccak256(
		['address', 'address', "uint256", "uint256", "bytes32"],
		[sender, receiver, id, receiveChainId, random]);
	let flatSigned = await owner.signMessage(ethers.utils.arrayify(messageHash));
	return [flatSigned,random]
}

async function applyCrossReceive(sender, receiver, id, senderChainId){

	let random = ethers.utils.randomBytes(32);
	let messageHash = ethers.utils.solidityKeccak256(
		['address', 'address', "uint256", "uint256", "bytes32"],
		[sender, receiver, id, senderChainId, random]);
	let flatSigned = await owner.signMessage(ethers.utils.arrayify(messageHash));
	return [flatSigned,random]
}

async function main() {

	[owner, user1, user2, ...users] = await hre.ethers.getSigners()

	// We get the contract to deploy
	const OmniOneNFT = await hre.ethers.getContractFactory("OmniOneNFT");
	const ooNFT = await OmniOneNFT.deploy(owner.address);

	await ooNFT.deployed();
	console.log("OmniOneNFT deployed to:", ooNFT.address);
	//  event SyncMinted(address from, address to, uint256 id);
	// ooNFT.on("SyncMinted", (setter, from, to, id, event) => {
	// 	console.log("sync minted id", id, " from ", from, ' to ', to,);
	// 	console.log(event);
	// })
	// // event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
	// ooNFT.on("Transfer", (from, to, id) => {
	// 	console.log("Transder id", id, " from ", from, ' to ', to,);
	// })

	// User0 apply to get an NFT
	{
		[id, random, flatSigned] = await applySyncMint(user1.address);
		console.log( id, random, flatSigned);

		let tx = await ooNFT.connect(user1).syncMint(user1.address, id, random, flatSigned);
		const receipt = await tx.wait()

		for (const event of receipt.events) {
			console.log(`Event ${event.event} with args ${event.args}`);
		}
	}
	// Cross transfer
	{
		let tokenID=0;
		let receiveChainId = 5 
		let sender =user1.address;
		let receiver = users[0].address
		let [flatSigned,random] = await applyCrossTransfer(sender,receiver, tokenID,receiveChainId)
		let tx = await ooNFT.connect(user1).crossTransfer(receiver, id, receiveChainId, random, flatSigned);
		const receipt = await tx.wait();
		// Listen event
		for (const event of receipt.events) {
			console.log(`Event ${event.event} with args ${event.args}`);
		}

	}
	// Cross receiver
	{
			
		let tokenID=0;
		let senderChainId = 5 
		let sender = users[0].address
		let receiver =user2.address;
		let [flatSigned,random] =await applyCrossReceive(sender, receiver,tokenID, senderChainId)
		let tx = await ooNFT.connect(user2).crossReceive(sender, tokenID, senderChainId, random, flatSigned);
		const receipt = await tx.wait();

		for (const event of receipt.events) {
			console.log(`Event ${event.event} with args ${event.args}`);
		}
	}

	// ===== get logs ==== //

	const logs = await hre.ethers.provider.getLogs({
		address:ooNFT.address,
		fromBlock:0,
        toBlock: 'latest'
	})

	// parse event by ooNFT
	for (const log of logs) {
		const result = ooNFT.interface.parseLog(log)
		console.log(result)
	}

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
