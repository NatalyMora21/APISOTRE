package info

import (
	sc "apiStore/schema"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ReadBuyersJson() []sc.Buyer {

	jsonFile, err := os.Open("./info/buyers.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//Info

	var buyers []sc.Buyer
	// we unmarshal our byteArray which contains our jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &buyers)
	//fmt.Println(buyers)
	return buyers

}

func ReadProductsCsv() []sc.Product {

	dFile, err := os.Open("./info/products.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer dFile.Close()

	r := csv.NewReader(dFile)
	r.Comma = ','
	r.FieldsPerRecord = 3

	var products []sc.Product
	//Leer linea por línea
	for {
		record, err := r.Read()
		//Error al final del archivo
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error", err)
		}

		//Mirar vacio
		productNew := sc.Product{
			Id:    record[0],
			Name:  record[1],
			DType: []string{"Product"},
		}
		if record[2] != "" {
			price, err := strconv.Atoi(record[2])
			if err != nil {
				log.Println("Error price", err)
				continue
			}
			productNew.Price = price

		}

		products = append(products, productNew)

	}
	//fmt.Println(products)
	return products

}

func ReadInfoTxt() []sc.Transaction {

	transactionsAll := []sc.Transaction{

		{
			Id:       "#0000611ef080",
			Buyer_id: "62efa0df",
			Ip:       "126.238.179.254",
			Device:   "android",
			Products: []string{"33f009db", "815c8e5f"},
		},
		{
			Id:       "#0000611ef081",
			Buyer_id: "a95841b4",
			Ip:       "119.204.182.131",
			Device:   "ios",
			Products: []string{"33f009db", "815c8e5f", "ca23a9ba", "88f9107b"},
		},
	}

	//fmt.Println(transactionsAll)
	return transactionsAll

}
