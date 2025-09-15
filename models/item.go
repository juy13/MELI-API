package models

type Item struct {
	ID                     string      `json:"id"` // suppose that id can be name of a product+id so better to keep string
	Title                  string      `json:"title"`
	Stars                  int         `json:"stars"`
	Price                  Price       `json:"price"`
	Shipping               Shipping    `json:"shipping"`
	Colors                 []string    `json:"colors"` // idk could be anything actually, size, color, size+color etc
	Available              bool        `json:"available"`
	Guarantee              string      `json:"guarantee"`
	ProductCharacteristics any         `json:"product_characteristics"` // just a json file with all characteristics
	Description            string      `json:"description"`
	Photos                 []string    `json:"photos"`     // ids of photos for request
	QASection              []QASection `json:"qa_section"` // questions and answers about the project
	Comments               []Comments  `json:"comments"`   // comments about the item
}

type Price struct {
	CurrencyID string  `json:"currency_id"` // can be usd or ars
	Amount     float64 `json:"amount"`
}

type Seller struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Sold     int         `json:"sold"`     // how much objects sold
	Products []ItemShort `json:"products"` // some recommended products from seller
	Icon     string      `json:"icon"`     // seller icon
}

type ItemShort struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"` // item title
	Price    Price    `json:"price"`
	Shipping Shipping `json:"shipping"` // when will be delivered
}

type QASection struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Comments struct {
	Comment string   `json:"comment"`
	Stars   int      `json:"Stars"`
	Photos  []string `json:"photos"` // ids of photos for request
}

type Shipping struct {
	Cost Price  `json:"cost"`
	Time string `json:"estimated_time"` // today, tomorrow, or a date
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// .. maybe some more stuff
}
