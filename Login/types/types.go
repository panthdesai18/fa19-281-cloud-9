package types

type User struct {
	Username  string `json:"username"`
	FullName string `json:"fullname"`
	Emailid   string `json:"emailid"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
