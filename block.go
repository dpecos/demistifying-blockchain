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

func (block *Block) Mine() {
	block.Hash = ""
	block.Nonce = -1

	for done := false; !done; done = strings.HasPrefix(block.Hash, "0000") {
		block.Nonce = block.Nonce + 1
		data := fmt.Sprintf("%d-%s-%d-%s", block.ID, block.Data, block.Nonce, block.Previous)
		block.Hash = calculateHash(data)
	}

	p := message.NewPrinter(language.English)
	p.Printf("Valid hash found! %s (nonce %d)\n", block.Hash, block.Nonce)
}

func calculateHash(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
