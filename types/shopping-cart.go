package types

type Shopping_Cart struct {
	ID          int    `json:"id"`
	Customer_Id int    `json:"customer_id"`
	Order_Id    int    `json:"order_id"` // 'omitempty' ensures that if Order_Id is 0, it doesn't show in the JSON response.
	Store_Id    int    `json:"store_id"`
	Active      bool   `json:"active"`
	Address     string `json:"address"` // 'omitempty' ensures that if Address is an empty string, it doesn't show in the JSON response.
	Created_At  string `json:"created_at"`
}

type Create_Shopping_Cart struct {
	Customer_Id int `json:"customer_id"`
}

type Get_Shopping_Cart struct {
	Customer_Id int  `json:"customer_id"`
	Active      bool `json:"active"`
}

type Shopping_Cart_Details struct {
	Customer_Id int `json:"customer_id"`
	Cart_Id     int `json:"cart_id"`
}

type AssignSlot struct {
	Customer_Id int `json:"customer_id"`
	Cart_Id     int `json:"cart_id"`
	Slot_Id     int `json:slot_id`
}

type ApplyPromo struct {
	CartId int    `json:"cart_id"`
	Promo  string `json:"promo"`
}

func New_Shopping_Cart(customer_id int) (*Create_Shopping_Cart, error) {
	return &Create_Shopping_Cart{
		Customer_Id: customer_id,
	}, nil
}
