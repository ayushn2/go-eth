package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://holesky.infura.io/v3/72edaacccf8a4b6597626eb20859082d"
	// murl = "https://mainnet.infura.io/v3/72edaacccf8a4b6597626eb20859082d"
)

func main(){
	// ks := keystore.NewKeyStore("../wallet",keystore.StandardScryptN,keystore.StandardScryptP)
	// _,err := ks.NewAccount("password")

	// if err!= nil{
	// 	log.Fatal(err)
	// }

	// _,err = ks.NewAccount("password")

	// if err!= nil{
	// 	log.Fatal(err)
	// }

	// "9b7df42e1916bc68d0d45d26105d61bacc2b4d5b"
	// "e78b48f93f020a03941b9e1a1dd78cf35b3b34f0"

	client,err := ethclient.Dial(url)
	if err != nil{
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("9b7df42e1916bc68d0d45d26105d61bacc2b4d5b")// convert hex to address
	a2 := common.HexToAddress("e78b48f93f020a03941b9e1a1dd78cf35b3b34f0")
	
	bal_1 ,err :=client.BalanceAt(context.Background(),a1, nil)
	if err!= nil{
		log.Fatal(err)
	}

	bal_2 ,err :=client.BalanceAt(context.Background(),a2, nil)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(a1,a2)

	fmt.Println(bal_1)
	fmt.Println(bal_2)
}