package input

type GetUsersInput struct {
	UserId string `json:"user_id"`
}

type CreateUserInput struct {
	LoginId  string `json:"login_id"`
	PassWord string `json:"password"`
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserInput struct {
	LoginId  string `json:"login_id"`
	PassWord string `json:"password"`
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

type DeleteUserInput struct {
	UserId string `json:"user_id"`
}
