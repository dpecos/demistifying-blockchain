package main

import "fmt"

// https://www.youtube.com/watch?v=_160oMzblY8

func main() {
	blockchain := BlockChain{
		Challenge: "0000",
		Blocks:    make(map[string]*Block),
	}

	fmt.Printf("Blockchain initialized with challenge '%s'\n\n", blockchain.Challenge)

	blockchain.Append("First block")
	blockchain.Append("Second block")
	blockchain.Append("Third block")

	fmt.Printf("\nResulting blockchain (in reversed order):\n\n")

	blockchain.Print()
}
