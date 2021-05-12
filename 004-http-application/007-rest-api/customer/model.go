package customer

type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	DOB     string  `json:"dob"`
	Address Address `json:"address"`
}

type Address struct {
	StreetName string `json:"street_name"`
	Area       string `json:"area"`
	Pincode    string `json:"pincode"`
}
