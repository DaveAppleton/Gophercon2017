package main

import (
	"math/big"
	"os"

	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//start deploy OMIT
func ownerTx(file string, password string) (res *bind.TransactOpts) {
	kr, _ := os.Open(file)
	res, _ = bind.NewTransactor(kr, password)
	return
}

var power *Power

func deploy() {
	var teller, tech common.Address
	endPoint := "/Users/daveappleton/Library/Ethereum/geth.ipc"
	client, _ := ethclient.Dial(endPoint)
	_, _, power, _ = DeployPower(ownerTx("owner.json", "password1"), client, teller, tech)
}

//end deploy OMIT

//start bal OMIT
func balance(user common.Address) *big.Int {
	b, _ := power.BalanceOf(nil, user)
	return b
}

//end bal OMIT

//start pay OMIT
func payBill(user common.Address, amount *big.Int) {
	tx, _ := power.PayBill(ownerTx("teller.json", "password2"), user, amount)
	fmt.Println("tx hash : ", tx.Hash().Hex())
}

//end pay OMIT
