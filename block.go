package main
import (
	"bytes" // need to convert data into byte in order to be sent on the network
	"crypto/sha256" // this will help us to create the hash of the block
	"strconv" // this will help us to convert int to string and vice versa
	"time" // this will help us to get the current time(timestamp)
)

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10)) // Get the current time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{ // concatenate all the block data
		timestamp,
		block.PreviousBlockHash,
		block.AllData,
	}, []byte{})

	// Create a new SHA-256 hash
	hash := sha256.Sum256(headers)

	// Set the block's hash to the computed hash
	block.MyBlockHash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	// Create a new block with the given previous block hash and data
	block := &Block{
		Timestamp: time.Now().Unix(), // Get the current time in seconds since the epoch
		PreviousBlockHash: prevBlockHash,
		AllData: []byte(data), // Convert the data into bytes
	}
	block.SetHash() // Set the hash of the block
	return block // Return the new block
}

func NewGenesisBlock() *Block {
	// Create the first block in the blockchain (the genesis block)
	return NewBlock("Genesis Block", []byte{}) // The previous block hash is empty for the genesis block
}