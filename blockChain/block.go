package blockChain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

type BlockChain struct {
	Blocks []*Block
}

type Data struct {
	ContractName                 string
	ProjectGeographicPositioning string
	StipulatedDuration           string
	StartDate                    time.Time
	EndDate                      time.Time
	WholeProjectCost             string
	ProjectMileStones            []*MileStones
	CWB                          []*CompaniesWhoBid
	AwardedCompany               string
	POInvolved                   []*ProcurementOfficers
	EntityAwarding               string
	Party1                       string
	Party2                       string
	EnfocersSignature            string
	InDepthContract              string
}

type CompaniesWhoBid struct {
	Id               []byte
	CompanyName      string
	CompanyPortfolio string
}
type MileStones struct {
	Id           []byte
	StartDate    time.Time
	EndDate      time.Time
	Expectations string
	Completed     string
	
}
type ProcurementOfficers struct {
	Id          []byte
	Name        string
	Designation string
}

type Block struct {
	Hash     []byte
	Data     *Data
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data *Data, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	return block
}

func (mil *MileStones) GenerateMilID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(mil)
	if err != nil {
		fmt.Println("Error")
	}
	hash = sha256.Sum256(encoded.Bytes())
	mil.Id = hash[:]
}

func (b *Block) HashMilestones() []byte {
	var Milhashes [][]byte
	var Milhash [32]byte

	for _, mil := range b.Data.ProjectMileStones {
		Milhashes = append(Milhashes, mil.Id)
	}
	Milhash = sha256.Sum256(bytes.Join(Milhashes, []byte{}))
	return Milhash[:]
}
func (cwb *CompaniesWhoBid) GenerateCWBID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(cwb)
	if err != nil {
		fmt.Println("Error")
	}
	hash = sha256.Sum256(encoded.Bytes())
	cwb.Id = hash[:]
}

func (b *Block) HashCompaniesWhoBid() []byte {
	var CWBhashes [][]byte
	var CWBhash [32]byte

	for _, cwb := range b.Data.CWB {
		CWBhashes = append(CWBhashes, cwb.Id)
	}
	CWBhash = sha256.Sum256(bytes.Join(CWBhashes, []byte{}))
	return CWBhash[:]
}


func (po *ProcurementOfficers) GeneratePOID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(po)
	if err != nil {
		fmt.Println("Error")
	}
	hash = sha256.Sum256(encoded.Bytes())
	po.Id = hash[:]
}

func (b *Block) HashProcurementOfficer() []byte {
	var POhashes [][]byte
	var POhash [32]byte

	for _, po := range b.Data.POInvolved {
		POhashes = append(POhashes, po.Id)
	}
	POhash = sha256.Sum256(bytes.Join(POhashes, []byte{}))
	return POhash[:]
}

func (chain *BlockChain) AddBlock(data *Data) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func GenesisBlock() *Block {
	return CreateBlock(&Data{}, []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}
