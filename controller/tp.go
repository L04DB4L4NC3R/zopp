/*
	Transacction processor
*/
package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type client struct {
	from          string
	to            string
	timestamp     string
	class         string
	ageSuggestive string
}

func (c client) tester() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("client api working"))
	}
}

func (c client) state(a common.Address, cc *ethclient.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		blk, err := cc.BlockByNumber(ctx, nil)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Some error occurred\n\n"))
			return
		}

		var y interface{}

		red := bytes.NewReader(blk.Bloom().Bytes())
		if err = rlp.Decode(red, &y); err != nil {
			log.Println(err)
			w.Write([]byte("Some error occurred\n\n"))
			return
		}
		log.Println(y)
		json.NewEncoder(w).Encode(blk.Body().Transactions)
		return
	}
}
func (c client) RegisterRoutes(a common.Address, cc *ethclient.Client) {
	http.HandleFunc("/api/v1/client/test", c.tester())
	http.HandleFunc("/api/v1/client/state", c.state(a, cc))
}
