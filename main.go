package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/zujisoft/sl/cc/pkg"
)

func main() {
	slc := new(pkg.SlContract)
	cc, err := contractapi.NewChaincode(slc)
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
