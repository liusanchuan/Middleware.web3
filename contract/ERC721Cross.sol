// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

interface ERC721Cross is IERC721 {
    // Emitted when `id` token is sync minted from `from` to `to`.
    event SyncMinted(address from, address to, uint256 id);
    // Emitted when `id` token is transfered by cross-chain from `from` to `to` on `receiveChainID` chain.
    event CrossTransfered(
        address from,
        address to,
        uint256 id,
        uint256 receiveChainID
    );
    // Emitted when `id` token is Received by cross-chain from `from` to `to` on `receiveChainID` chain.
    event CrossReceived(
        address from,
        address to,
        uint256 id,
        uint256 senderChainID
    );

    function syncMint(
        address to,
        uint256 id,
        bytes memory evidence
    ) external;

    function crossTransfer(
        address to,
        uint256 id,
        uint256 receiveChainID,
        bytes memory evidence
    ) external;

    function crossReceive(
        address from,
        uint256 id,
        uint256 senderChainID,
        bytes memory evidence
    ) external;
}

contract OmniOneNFT is ERC721Cross, ERC721, Ownable {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;
    address public proxyOwner;

    constructor(address _proxyOwner) ERC721("OmniOneNFT", "OONFT") {
        proxyOwner = _proxyOwner;
    }

    function syncMint(
        address to,
        uint256 id,
        bytes memory evidence
    ) external override {
        bytes32 msgHash = keccak256(abi.encodePacked(msg.sender, to, id));
        // Check ECDSA signature
        require(
            recoverSigner(msgHash, evidence) == proxyOwner,
            "ERR: Your evidence not valid!"
        );
        
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);

        emit SyncMinted(msg.sender, to, id);
    }

    function crossTransfer(
        address to,
        uint256 id,
        uint256 receiveChainID,
        bytes memory evidence
    ) external override {}

    function crossReceive(
        address from,
        uint256 id,
        uint256 senderChainID,
        bytes memory evidence
    ) external override {}

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
    }

    function recoverSigner(bytes32 msgHash, bytes memory sign)
        public
        pure
        returns (address)
    {
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHashMessage = keccak256(
            abi.encodePacked(prefix, msgHash)
        );
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(sign);
        return ecrecover(prefixedHashMessage, v, r, s);
    }

    function splitSignature(bytes memory sig)
        public
        pure
        returns (
            bytes32 r,
            bytes32 s,
            uint8 v
        )
    {
        require(sig.length == 65, "invalid signature length");

        assembly {
            /*
            First 32 bytes stores the length of the signature

            add(sig, 32) = pointer of sig + 32
            effectively, skips first 32 bytes of signature

            mload(p) loads next 32 bytes starting at the memory address p into memory
            */

            // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
            // second 32 bytes
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        // implicitly return (r, s, v)
    }
}
