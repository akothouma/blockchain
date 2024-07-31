package blockChain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct{
	ID []byte
	Input []TxInput
	Output []TxOutput

}

type TxInput struct{
	ID []byte //ref transaction that output is part of
	OutID int
	Signature string

}

type TxOutput struct{
	Value int //tokens themselves
	PubKey string 
}

func Coinbase(to,data string)*Transaction{
if data==""{
	data=fmt.Sprintf("Coins to %s " ,to)
}
txOut:=TxOutput{100,to}
txIn:=TxInput{[]byte{},-1,data}

tx:=Transaction{nil,[]TxInput{txIn},[]TxOutput{txOut}}
tx.GenerateTxID()
return &tx
}

func (tx *Transaction) GenerateTxID(){
	var encoded bytes.Buffer
	var hash [32]byte

	encode:=gob.NewEncoder(&encoded)
	err:=encode.Encode(tx)
	if err!=nil{
		fmt.Println("Error")
	}
	hash=sha256.Sum256(encoded.Bytes())
	tx.ID=hash[:]
}

func(tx *Transaction)IsCoinbase()bool{
	return len(tx.Input)==1 && len(tx.Input[0].ID)==0 && tx.Input[0].OutID==-1 

}

func (in *TxInput)CanUnlock(data string)bool{
	return  in.Signature==data
}

func (out *TxOutput)CanBeUnlocked(data string)bool{
	return out.PubKey==data
}