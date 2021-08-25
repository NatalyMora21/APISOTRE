package controllers

import (
	"apiStore/info"
	sc "apiStore/schema"
	"fmt"
)

//Structure the information related to its types
func RelationInfo() {
	//Buscar por cada usuario las transacciones relaciondas

	//Mutationproducts()

	buyersall := info.ReadBuyersJson()

	var buyerStore []sc.InfoBuyer

	for i := 0; i < len(buyersall); i++ {
		transactionsbuyer := foundTransaction(buyersall[i].Id)

		newBuyer := sc.InfoBuyer{
			Id:           buyersall[i].Id,
			Name:         buyersall[i].Name,
			Age:          buyersall[i].Age,
			Transactions: transactionsbuyer,
			DType:        []string{"Buyer"},
		}

		buyerStore = append(buyerStore, newBuyer)

	}

	fmt.Println(buyerStore)

	//Hacer la mutacion
	MutationInfoBuyers(buyerStore)
	fmt.Println("Mutation realizada")

}

//Funcion que actualiza la información de la transacción de los productos, agrega los id de la db para relacionar los nodos
func RelationProductos() []sc.TransactionProduct {

	//Traer las transacciones
	//Ese sería el nodo relacionado al usuario
	allTransactions := info.ReadInfoTxt()
	//Recorrer el arreglo y y comparar los id con los que se guardaron en l db
	var newTransaccionType []sc.TransactionProduct

	//Recorre las transacciones y los productos de esas transacciones que es una array []
	//Crear con la informacion un nuevo type para guardar las relaciones de los porductos
	for i := 0; i < len(allTransactions); i++ {
		var idsproducts []sc.IdProduct
		for j := 0; j < len(allTransactions[i].Products); j++ {
			uid := foundIdProduct(allTransactions[i].Products[j])
			idproductr := sc.IdProduct{
				Uid: uid,
			}
			idsproducts = append(idsproducts, idproductr)

		}

		transactionNewFormat := sc.TransactionProduct{
			Id:       allTransactions[i].Id,
			Buyer_id: allTransactions[i].Buyer_id,
			Ip:       allTransactions[i].Ip,
			Device:   allTransactions[i].Device,
			Products: idsproducts,
			DType:    []string{"Transaction"},
		}

		newTransaccionType = append(newTransaccionType, transactionNewFormat)

	}

	return (newTransaccionType)

	//fmt.Println(allTransactions)

}

func foundIdProduct(id string) string {

	//Trar los guardador en la db

	var dbproducts = Getproducts()

	fmt.Println("Getproducts", dbproducts)

	var idproductdb string

	for i := 0; i < len(dbproducts); i++ {
		if dbproducts[i].Id == id {
			println("Igual", dbproducts[i].Uid)
			idproductdb = dbproducts[i].Uid

		}
	}
	return (idproductdb)

}

func foundTransaction(idBuyer string) []sc.TransactionProduct {

	transactions := RelationProductos()

	var transactionxbuyer []sc.TransactionProduct

	for i := 0; i < len(transactions); i++ {

		if transactions[i].Buyer_id == idBuyer {
			transactionxbuyer = append(transactionxbuyer, transactions[i])

		}

	}

	return transactionxbuyer

}
