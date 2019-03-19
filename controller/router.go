package controller

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var c Client

func StartClient(a common.Address, cc *ethclient.Client, auth *bind.TransactOpts) {
	c.RegisterRoutes(a, cc, auth)
}
