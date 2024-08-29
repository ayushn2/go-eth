package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
	pvk,err := crypto.GenerateKey() //This function will generate private key

	if err!= nil{
		log.Fatal(err)

		
	}

	pData := crypto.FromECDSA(pvk) //ECDSA algorithm
	encodedData := hexutil.Encode(pData)//Encode private data
	fmt.Println(encodedData)


	// pbk,err := pvk.PublicKey
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	publicKeyAddr := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	fmt.Println(publicKeyAddr)
}

