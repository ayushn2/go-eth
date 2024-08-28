package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)



func main(){

	if err := godotenv.Load(".env.local"); err != nil {
        log.Fatalf("Error loading .env.local file: %v", err)
    }
	infuralURL := os.Getenv("INFURAL_URL") // go infers the datatype itself :: using := for short varibale declaration
	// ganacheURL := os.Getenv(("GANACHE_URL"))

	// Creating a client
	client,err := ethclient.DialContext(context.Background(),infuralURL)//Creating a new ethereum client we can use ganache URL instead if we want to develop locally

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

	addr := "0x077D360f11D220E4d5D831430c81C26c9be7C4A4"
	address := common.HexToAddress(addr) //Converting hex into address

	balance,err := client.BalanceAt(context.Background(),address,nil)

	if err!=nil{
		log.Fatal("Error in getting the balance : ",err)
	}
	fmt.Println("The balance of the address : ",address," is :",balance)
	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)
	value := new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))
	fmt.Println(value.Text('f',18)," ETH")
}