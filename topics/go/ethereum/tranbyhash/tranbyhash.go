package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/fea6924127a54566b8c1ce5e4cdfdb21")
	if err != nil {
		log.Fatalln("Dial:", err)
	}

	hash := common.HexToHash("0x30bb36153d1e6e4304eb2991f305a24be6da8565b185e7a7af2d22381c6f79b2")
	tx, pending, err := conn.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Fatalln("TBH:", err)
	}

	if pending {
		log.Println("Transaction Pending")
		return
	}

	fmt.Printf("%#v\n", tx)
}
