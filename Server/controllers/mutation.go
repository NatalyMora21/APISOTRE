package controllers

import (
	"apiStore/info"
	sc "apiStore/schema"
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func Mutationproducts() {

	ctx := context.Background()

	//Data Productsg

	products := info.ReadProductsCsv()
	println(products)

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	response, err := C.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	println(response)

}

func MutationInfoBuyers(info []sc.InfoBuyer) {

	ctx := context.Background()

	//Data Products

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	response, err := C.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	println(response)

}
