package controller

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var c Client

func StartClient(a common.Address, cc *ethclient.Client, auth *bind.TransactOpts) http.Handler {
	router := c.RegisterRoutes(a, cc, auth)
	CorsRouter := cors.Default().Handler(router)
	return CorsRouter
}
