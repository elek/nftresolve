# Resolve NFT

Most NFTs are just a link (or some metadata). This simple application calls the smart contract based on the available NFT standards (ERC721/ERC1155) to retrieve the metadata from the URL.

The only thing what you need 

 * contract id
 * id of the NFT token
 * (chain rpc endpoint)

Example:

```
nftresolve --chain avax 0xd17189a6a70D33c733d745748d4352bf856C953F 5

nftresolve --chain avax 0xd17189a6a70D33c733d745748d4352bf856C953F 5
URI: https://ipfs.infura.io/ipfs/QmSTJsbqFeqVrfGnNLjuJ5vpxBrUdgBJU2w6ZbTxTgHijr
{
   "attributes": [],
   "description": "This collection includes water colored map of capitals all around the word. The name of the cities are intentionally missing..."
   "image": "https://ipfs.infura.io/ipfs/QmVraHy82d2JVcG6BDN32tzwocZMFKZYA7uyCPY9SGqdTr",
   "name": "The ballad of far cities #5"
}
```
