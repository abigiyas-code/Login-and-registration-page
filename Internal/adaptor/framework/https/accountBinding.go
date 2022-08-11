package https

type SendAccountRequest struct {
	Firstname    string `json:"firstname" binding:"required"`
	Lastname     string `json:"lastname" binding:"required"`
	Emailaddress string `json:"emailaddress" binding:"required"`
	Username     string `json:"username" binding:"required_with=@ username"`
	Password     int64  `json:"password" binding:"min=4,max=10,"`
}
