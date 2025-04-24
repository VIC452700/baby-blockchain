package main

// Create the method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // Get the last block in the blockchain
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash) // Create a new block with the previous block's hash and the new data
	blockchain.Blocks = append(blockchain.Blocks, newBlock) // Append the new block to the blockchain
}

// Create the function that returns the whole blockchain and add the genesis to it first. 
func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock() // Create the genesis block
	return &Blockchain{Blocks: []*Block{genesisBlock}} // Return a new blockchain with the genesis block
}