package main

import (
	"fmt"
	"strconv"
	"time"
	"hackerthon/blockChain"
)

func main() {
	chain := blockChain.InitBlockChain()

	data1:=&blockChain.MileStones{
	     StartDate:time.Date(2024, time.January, 17,0,0,0,0,time.Local),
		 EndDate: time.Date(2025, time.March, 14,0,0,0,0,time.Local),
		 Expectations: "All children under 5 years in all counties + refugees vaccinated on polio",
		
	}
		
	data2:=&blockChain.MileStones{
	     StartDate:time.Date(2025, time.November, 24,0,0,0,0,time.Local),
		 EndDate: time.Date(2026, time.March, 14,0,0,0,0,time.Local),
		 Expectations: "All girl children under 10 years in all counties + refugees vaccinated on HPV",
		
	}
	company1:=&blockChain.CompaniesWhoBid{
		CompanyName: "TATA",
		CompanyPortfolio: "link to company's previous transactions",
	}
	company2:=&blockChain.CompaniesWhoBid{
		CompanyName:"BICO",
		CompanyPortfolio: "link to company's previous transactions",
	}

	po1:=&blockChain.ProcurementOfficers{
		Name: "Lorna Akoth",
		Designation:" Procurement& store",
	}
	po2:=&blockChain.ProcurementOfficers{
		Name: "Raymond Ogwel",
		Designation:"supply-chain quality assurance",
	}

	chain.AddBlock(&blockChain.Data{
		 ContractName:"Vaccine Delivery in SubSaharan Africa\n",
		 ProjectGeographicPositioning: "Western Kenya\n",
		 StipulatedDuration: "2 years\n", 
		 StartDate:time.Now(),
		 EndDate: time.Date(2026, time.March, 14,0,0,0,0,time.Local),
		 WholeProjectCost: "\n2,000,0000,0000 USD",
		 ProjectMileStones:[] *blockChain.MileStones{data1,data2},
		 CWB: []*blockChain.CompaniesWhoBid{company1,company2},
		 AwardedCompany: "TATA",
		 POInvolved: []*blockChain.ProcurementOfficers{po1,po2},
		 EntityAwarding: "USAID",
		 Party1: "MOH-KENYA",
		 Party2:"LANCET",
		 EnfocersSignature:"WITNESS ADDRESS",
		 InDepthContract: "link to file",
		 	})
	

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash %x\n", block.PrevHash)
		fmt.Printf("Data in Block %v\n", block.Data)
		fmt.Printf("Current Hash %x\n", block.Hash)

		pow := blockChain.NewProofOfWork(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
