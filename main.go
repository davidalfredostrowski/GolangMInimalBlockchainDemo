package main

import (
	"fmt"
	"time"
	"crypto/sha256"
)


type Block struct {
	timestamp time.Time
	transactions []string
	prevhash     []byte
   	Hash         []byte
}

func main() {
	abc := []string{" A sent 50 coins to BC"}
	xyz := Blocks(abc, []byte{})
	fmt.Println("this is our first Block")
	Print(xyz)

	pqrs := []string{" PQ sent 230 coins to RS"}
	klmn := Blocks(pqrs, xyz.Hash)
	fmt.Println("This is our second block")
	Print(klmn)
 }




func Blocks(transactions []string, prevhash []byte) *Block {
	currentTime := time.Now() 
	return &Block{
		timestamp: currentTime,
		transactions: transactions,
		prevhash: prevhash,
		Hash: NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions[]string, prevhash[]byte) []byte{
	input := append(prevhash, time.String()...)
	for transactions := range transactions {
		input = append(input, string(rune(transactions))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block){
	fmt.Printf("\ttime: %s\n",block.timestamp.String())
	fmt.Printf("\tprevhash: %x\n", block.prevhash)
	fmt.Printf("\thash: %x\n",block.Hash)
	Transaction(block)
}

func Transaction(block *Block){
	fmt.Println("\tTransactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n",i, transaction)
	}
}

