package main

// https://www.youtube.com/watch?v=_160oMzblY8

func main() {
	blockchain := BlockChain{
		LastID: -1,
		Blocks: make(map[string]*Block),
	}

	blockchain.Append("First block")
	blockchain.Append("Second block")
	blockchain.Append("Third block")

	blockchain.Print()
}
