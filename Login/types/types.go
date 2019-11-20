package types

type User struct {
	Username  string `json:"username"`
	FullName string `json:"fullname"`
	Emailid   string `json:"emailid"`
	Password  string `json:"password"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
