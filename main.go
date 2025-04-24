// use the main package
package main

// import the fmt package
import (
	"fmt"
)

func main() {
	fmt.Println("------------- I am building my first blockchain! ---------------")

	newBlockchain := NewBlockchain() // Create a new blockchain
	// Create 2 blocks and add 2 transactions
	newBlockchain.AddBlock("Transaction 1") // first block containing one transaction
	newBlockchain.AddBlock("Transaction 2") // second block containing one transaction
	// Print all the blocks and their contents
	for i, block := range newBlockchain.Blocks { // Loop through all the blocks in the blockchain
		fmt.Printf("Block ID : %d\n", i) // Print the block ID
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i)) // print the timestamp of the block, to make them different, we just add a value i
		fmt.Printf("Block Hash %x\n", block.MyBlockHash) // Print the hash of the block
		fmt.Printf("Previous Hash: %x\n", block.PreviousBlockHash) // Print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData) // Print the data of the block
	}
}