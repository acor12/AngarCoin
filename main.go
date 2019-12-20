package main

import (
	"github.com/acor12/AngarCoin/node"
)

func main() {
	block := node.NewBlock([]byte{}, []byte{})
	println(block.Serialize())
	println(node.Deserialize(block.Serialize()).BlockDataHash)
}
