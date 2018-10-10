package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Block struct {
	Number    int
	Timestamp time.Time
	Data      string
	Nonce     int
	Hash      string
	Previous  string
}

func (block *Block) String() string {
	return fmt.Sprintf("Block %d: {\n   ts:\t\t%s\n   nonce:\t%d\n   hash:\t%s\n   previous:\t%s\n}", block.Number, block.Timestamp, block.Nonce, block.Hash, block.Previous)
}

func (block *Block) Mine() {

	fmt.Printf("Mining block %d... ", block.Number)

	tsStart := time.Now()

	block.Hash = ""
	block.Nonce = -1

	for done := false; !done; done = strings.HasPrefix(block.Hash, "000") {
		block.Nonce = block.Nonce + 1
		data := fmt.Sprintf("%d-%d-%s-%d-%s", block.Number, block.Timestamp.Nanosecond(), block.Data, block.Nonce, block.Previous)
		block.Hash = calculateHash(data)
	}

	tsEnd := time.Now()
	elapsedTime := tsEnd.Sub(tsStart)

	p := message.NewPrinter(language.English)
	p.Printf("Valid hash found in %s! (hash: %s, nonce: %d)\n", elapsedTime.String(), block.Hash, block.Nonce)
}

func calculateHash(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
