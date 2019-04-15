package main

import (
	. "github.com/LowEntropyBody/basic-blockchain/blockchain"
)

func main() {
	bc := NewBlockchain()
	defer bc.Db.Close()

	cli := CLI{bc}
	cli.Run()
}
