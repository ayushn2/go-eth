package main

import (
	"context"
	"fmt"

	"log"
	// "math/big"
	// "os"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
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
	a1 := common.HexToAddress("9b7df42e1916bc68d0d45d26105d61bacc2b4d5b")// convert hexa string to address
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

	// nonce,err := client.PendingNonceAt(context.Background(),a1)

	// if err!= nil{
	// 	log.Fatal(err)
	// }

	fmt.Println(bal_1)
	fmt.Println(bal_2)

	// 1 ether = 10^18 wei

	// amount := big.NewInt(100000000000000000)
	// gasPrice,err := client.SuggestGasPrice(context.Background())

	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// tx := types.NewTransaction(
	// 	nonce,
	// 	a2,
	// 	amount,
	// 	21000,//Gas limit
	// 	gasPrice,
	// 	nil,
	// )
	// chainID,err := client.NetworkID((context.Background()))
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// b,err := os.ReadFile("../wallet/UTC--2024-09-10T21-25-41.805659000Z--9b7df42e1916bc68d0d45d26105d61bacc2b4d5b")

	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// key,err := keystore.DecryptKey(b, "password")
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// tx,err = types.SignTx(tx,types.NewEIP155Signer(chainID),key.PrivateKey)
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// err = client.SendTransaction(context.Background(),tx)
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println("tx sent : %s",tx.Hash().Hex()) //Printing the transaction hash
	
// 0xe953543e5cc91d317aaf416ff8898ec691198bce688be1f4e9f33d1badaf6e27
}