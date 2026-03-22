package cart

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

type Cart struct {
	Items []Item `json:"items"`
}

//ergwergv
