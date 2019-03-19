/*
	Transacction processor
*/
package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/rlp"

	store "github.com/angadsharma1016/technica/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Latitude      string `json:"lat"`
	Longitude     string `json:"lng"`
	Timestamp     string `json:"time"`
	NearestPoleID string `json:"id"`
}

func (c Client) tester() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Client api working"))
	}
}

func (c Client) state(a common.Address, cc *ethclient.Client) http.HandlerFunc {
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

func (c *Client) create(a common.Address, cc *ethclient.Client, auth *bind.TransactOpts) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&c)
		log.Println(c)

		instance, err := store.NewStore(a, cc)
		if err != nil {
			log.Fatal(err)
			return
		}
		tx, err := instance.SetData(auth, c.NearestPoleID, c.Latitude, c.Longitude, c.Timestamp)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("tx sent: %s", tx.Hash().Hex())

		// result, err := instance.WatchDataSetter(nil, key)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(string(result[:]))

		json.NewEncoder(w).Encode(struct {
			Message string
		}{"Done"})
	}
}

func (c Client) RegisterRoutes(a common.Address, cc *ethclient.Client, auth *bind.TransactOpts) {
	http.HandleFunc("/api/v1/client/test", c.tester())
	http.HandleFunc("/api/v1/client/state", c.state(a, cc))
	http.HandleFunc("/api/v1/client/create", c.create(a, cc, auth))
}
