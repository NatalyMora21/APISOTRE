package schema

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Uid   string `json:"uid"`
}

type Transaction struct {
	Id       string   `json:"id"`
	Buyer_id string   `json:"buyer_id"`
	Ip       string   `json:"ip"`
	Device   string   `json:"device"`
	Products []string `json:"products"`
}

type TransactionProduct struct {
	Id       string      `json:"id"`
	Buyer_id string      `json:"buyer_id"`
	Ip       string      `json:"ip"`
	Device   string      `json:"device"`
	Products []idProduct `json:"products"`
}

type idProduct struct {
	Uid string `json:"uid"`
}

//Info para crear los nods en Dgraph
//La info de los productos sería el id que se trae al crearlos en la Bd
type InfoBuyer struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Transactions []Transaction `json:"transactions"`
}
