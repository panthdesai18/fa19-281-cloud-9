package model

type Location struct {
	LocationId string `json:"locationId"`
	LocationName  string `json:"locationName"`
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type ResponseResult struct{
	Error  string `json:"error"`
	Result string `json:"result"`
}