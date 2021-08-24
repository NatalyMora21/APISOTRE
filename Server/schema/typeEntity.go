package schema

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Price int      `json:"price"`
	Uid   string   `json:"uid"`
	DType []string `json:"dgraph.type,omitempty"`
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
	Products []IdProduct `json:"products"`
	DType    []string    `json:"dgraph.type,omitempty"`
}

type Transactionj struct {
	Id       string    `json:"id"`
	Buyer_id string    `json:"buyer_id"`
	Ip       string    `json:"ip"`
	Device   string    `json:"device"`
	Products []Product `json:"products"`
	DType    []string  `json:"dgraph.type,omitempty"`
}

type IdProduct struct {
	Uid string `json:"uid"`
}

//Info para crear los nods en Dgraph
//La info de los productos sería el id que se trae al crearlos en la Bd
type InfoBuyer struct {
	Id           string               `json:"id"`
	Name         string               `json:"name"`
	Age          int                  `json:"age"`
	Transactions []TransactionProduct `json:"transactions"`
	DType        []string             `json:"dgraph.type,omitempty"`
}

type InfoBuyerj struct {
	Uid          string         `json:"uid"`
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Transactions []Transactionj `json:"transactions"`
	DType        []string       `json:"dgraph.type,omitempty"`
}
