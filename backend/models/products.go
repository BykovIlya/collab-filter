package models

/**
	the struct of items
 */
type Items struct{
  Itemid_string string  `json:"itemid_string"`
  Itemid_count float64  `json:"itemid_count"`
}

type ItemsGlobal struct {
  Itemid string         `json:"itemid"`
  Count int64           `json:"count"`
}

type Product struct {
  id int64              `json:"product_id"`
  name string           `json:"product_name"`
  cathegory string      `json:"product_cathegory"`
  price float64         `json:"product_price"`
}


