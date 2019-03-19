package controller

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var c client

func StartClient(a common.Address, cc *ethclient.Client) {
	c.RegisterRoutes(a, cc)
}
