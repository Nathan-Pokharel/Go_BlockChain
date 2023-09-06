package main

import (
	"fmt"
	"strconv"
    "os"
    "bufio"
	blockchain "github.com/Nathan-Pokharel/Blockchain/Blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
    var input string
    var again string
    scanner := bufio.NewScanner(os.Stdin)
    for{
        fmt.Println("Enter Data To Add To The Blockchain")
        scanner.Scan()
        input = scanner.Text()
        chain.AddBlock(input)
        fmt.Print("Add Another? y/n : ")
        fmt.Scanln(&again)
        if again != "y"{
            break
        }
    }
	for _, block := range chain.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
        pow := blockchain.NewProof(block)
        fmt.Printf("Proof Of Work: %s\n",strconv.FormatBool(pow.Validate()))
        fmt.Println()

	}
}
