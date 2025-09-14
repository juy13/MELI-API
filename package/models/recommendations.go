package models

/*
but actually imho maybe just have this enough:
	type Recommendations struct {
		Items []ItemShort `json:"items"`
	}
*/

// let's have this kind recommendations
type Recommendations struct {
	RelatedItems    []ItemShort `json:"related_items"`
	BuyAlso         []ItemShort `json:"buy_also"`
	CheaperProducts []ItemShort `json:"cheaper_products"`
}
