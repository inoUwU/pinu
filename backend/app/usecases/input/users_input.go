package input

// GetUsersInput ユーザー一覧取得の入力パラメータ
type GetUsersInput struct {
	// 現在は特にパラメータはないが、将来的にページング等に対応
}

type CreateUserInput struct {
	LoginId  string `json:"login_id"`
	PassWord string `json:"password"`
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserInput struct{}
type DeleteUserInput struct{}
