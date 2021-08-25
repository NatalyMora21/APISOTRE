package controllers

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func CreateSchema() {

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		id: string .
		price: int .
		age: int .
		buyer_id: string .
		ip: string .
		device: string .
		transactions: [uid] .
		products: [uid] .
		type Product {
			id: string
			name: string
			price: int
		}
		type Transaction {
			id: string
			buyer_id: string
			ip: string
			device: string
			products: [Product]
		}
		type Buyer {
			id: string
			name: string
			age: int
			transactions: [Transaction]
		}
	`

	ctx := context.Background()
	if err := C.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

}
