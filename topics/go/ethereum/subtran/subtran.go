package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/fea6924127a54566b8c1ce5e4cdfdb21")
	if err != nil {
		log.Fatalln("Dial:", err)
	}

	ch := make(chan *types.Header)
	sub, err := conn.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		log.Fatalln("Sub:", err)
	}
	defer sub.Unsubscribe()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	for {
		select {
		case header := <-ch:
			fmt.Printf("%#v\n", header)
		case <-shutdown:
			fmt.Println("shutdown")
			return
		}
	}
}

/*
addr := common.BytesToAddress([]byte("0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"))

	q := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	ch := make(chan types.Log)
	sub, err := conn.SubscribeFilterLogs(context.Background(), q, ch)
	if err != nil {
		log.Fatalln("Sub:", err)
	}
	defer sub.Unsubscribe()
*/
