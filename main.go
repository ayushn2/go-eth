package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
)



func main(){

	if err := godotenv.Load(".env.local"); err != nil {
        log.Fatalf("Error loading .env.local file: %v", err)
    }
	// infuralURL := os.Getenv("INFURAL_URL") // go infers the datatype itself :: using := for short varibale declaration
	ganacheURL := os.Getenv(("GANACHE_URL"))

	// Creating a client
	client,err := ethclient.DialContext(context.Background(),ganacheURL)//Creating a new ethereum client we can use ganache URL instead if we want to develop locally

	if err != nil{
		log.Fatal("Error to create a ether client:",err)
	}
	defer client.Close()
	// Getting a block by number
	block,err := client.BlockByNumber(context.Background(),nil) //second is the id of block : if nil is passed as second argument we automaticlly get the last block mined
	if err != nil{
		log.Fatal("Error to detect block:",err)
	}
	fmt.Println(block.Number())

}