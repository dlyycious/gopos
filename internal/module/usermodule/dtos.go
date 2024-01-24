package usermodule

type UserDtos struct {
	Username  string `json:"username"`
	Password  string `json:"-"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Goid      uint   `json:"goid"`
	IsVerfied bool   `json:"is_verified"`
	RoleId    uint   `json:"-"`
	Role      string `json:"role"`
}
