package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
	// key := keystore.NewKeyStore("../wallet",keystore.StandardScryptN,keystore.StandardScryptP)
	password := "password"
	// a,err := key.NewAccount(password) 
	// if err!= nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(a.Address)

	b,err := os.ReadFile("../wallet/UTC--2024-09-05T23-42-56.696135000Z--20101a10859a706662f65297a2ffea5206bb32c5")
	if err != nil{
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err !=nil{
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private key ",hexutil.Encode(pData)) // printing private key 

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public ",hexutil.Encode(pData))

	fmt.Println("Add ",crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}