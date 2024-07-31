package blockChain

import (
	"bytes"
	"crypto/sha256"
	"hackerthon/utils"
	"math"
	"math/big"
)

const Difficulty =12// TO-DO algo to increment this value over a large period of ime

type ProofOfWorkStruct struct{
	Block *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWorkStruct{
	target:=big.NewInt(1)
	target.Lsh(target,uint(256-Difficulty))
	pow:=&ProofOfWorkStruct{b,target}
	return pow
}

func (pow *ProofOfWorkStruct)Run()(int,[]byte){
	var intHash big.Int
	var hash [32] byte
	nounce:=0

	for nounce<math.MaxInt64{
		data:=pow.InitData(nounce)
		hash=sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target)==-1{
           break
		}else{
			nounce++
		}
	}
return nounce,hash[:]
}

func (pow *ProofOfWorkStruct)InitData(nounce int)[]byte{
	data:=bytes.Join([][]byte{
		[]byte(pow.Block.Data.AwardedCompany),
		[]byte(pow.Block.Data.ContractName),
		[]byte(pow.Block.Data.ProjectGeographicPositioning),
		[]byte(pow.Block.Data.StipulatedDuration),
	    []byte(pow.Block.Data.StartDate.Local().String()),
        []byte(pow.Block.Data.EndDate.Local().String()),
		[]byte(pow.Block.Data.WholeProjectCost),
	    pow.Block.HashMilestones(),
		pow.Block.HashCompaniesWhoBid(),
	    []byte(pow.Block.Data.WholeProjectCost),
		[]byte(pow.Block.Data.AwardedCompany),
	    pow.Block.HashProcurementOfficer(),
		[]byte(pow.Block.Data.EntityAwarding),
		[]byte(pow.Block.Data.Party1),
		[]byte(pow.Block.Data.Party2),
		[]byte(pow.Block.Data.EnfocersSignature),
		[]byte(pow.Block.Data.InDepthContract),
		pow.Block.PrevHash,
		utils.ToHex(int64(nounce)),
		utils.ToHex(int64(Difficulty)),

	},
	[]byte{},
)
	return data
}

func (pow *ProofOfWorkStruct)Validate()bool{
	var intHash big.Int
	data:=pow.InitData(pow.Block.Nonce)
	hash:=sha256.Sum256(data)

	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target)==-1
}