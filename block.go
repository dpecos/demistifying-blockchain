package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Block struct {
	ID       int
	Data     string
	Nonce    int
	Hash     string
	Previous string
}

func (block *Block) CalculateHash() {
	hash := ""
	block.Nonce = -1

	for done := false; !done; done = strings.HasPrefix(hash, "0000") {
		block.Nonce = block.Nonce + 1
		// fmt.Printf("  Trying with nonce %d\n", block.Nonce)
		data := fmt.Sprintf("%d-%s-%d-%s", block.ID, block.Data, block.Nonce, block.Previous)
		// fmt.Println(data)
		hasher := sha1.New()
		hasher.Write([]byte(data))
		hash = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	}

	p := message.NewPrinter(language.English)
	p.Printf("Valid hash found! %s (nonce %d)\n", hash, block.Nonce)

	block.Hash = hash
}
