package schema

import (
	"apiStore/db"
	"context"
	"log"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func CreateSchema() {

	op := &api.Operation{}
	op.Schema = `
		name: string .
		age: int .
		id: string.
		price: int.
		buyer_id: string.
		ip:string.
		device: string .
		transactions: [uid].
		products: [uid] .
		type Buyer {
			id: string
			name: string
			age: int
			transactions: [Transaction]

		}
		type Product {
			id: string
			name: string
			price: int
		}
		type Transaction {
			id: string
			buyer_id: string
			ip: string
			device string
			products: [Product]
		}
	`

	ctx := context.Background()
	if err := db.DgraphClient.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

}
