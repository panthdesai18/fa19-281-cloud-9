package model

type OrderPayment struct {
	PaymentId string `json:"paymentId"`
	OrderId  string `json:"orderId"`
	Amount 	 float64 `json:"amount"`
	CardNumber 	string `json:"cardnumber"`
	Cvv 	string `json:"cvv"`
	ExpiprationDate string `json:"expiprationDate"`
	PostalCode string `json:"postalCode"`
	CountryCode string `json:"countryCode"`
}

type ResponseResult struct{
	Error  string `json:"error"`
	Result string `json:"result"`
}