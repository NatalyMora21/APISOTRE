package controllers

import (
	sc "apiStore/schema"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Getproducts() []sc.Product {

	txn := C.NewTxn()

	const q = `
	{
		products(func:has(name)){
		uid
		id
		name
		price
	  }
	}
	`

	type Root struct {
		Products []sc.Product `json:"products"`
	}

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r.Products

}

//Query all Buyers

func AllBuyers() {

	txn := C.NewTxn()

	const querybuyers = `
	{
		buyers(func: type(Buyer)) {
		   name
			 age
	 
	   }
	 }
	`

	type AllBuyers struct {
		Buyers []sc.Buyer `json:"buyers"`
	}

	var listBuyers AllBuyers

	resp, err := txn.Query(context.Background(), querybuyers)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Json, &listBuyers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listBuyers.Buyers)
	//return r.Products

}

func BuyerId() {

	txn := C.NewTxn()

	const querybuyers = `
	{
		buyer(func:uid(0xd7)) {
			uid
		  	name
			age
			transactions{
			 	ip
			 	device
			 	products{
					name
			    	price
			 }
		   }
	   }
	 }
	`

	type BuyerId struct {
		Buyer []sc.InfoBuyerj `json:"buyer"`
	}

	var buyer BuyerId

	resp, err := txn.Query(context.Background(), querybuyers)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Json, &buyer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buyer.Buyer[0].Transactions)
	//return r.Products
	out, _ := json.MarshalIndent(buyer, "", "\t")
	fmt.Printf("%s\n", out)

	//Compradores con la misma ip

}
