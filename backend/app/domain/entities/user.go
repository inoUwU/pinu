package entities

// TODO: Refer to the database schema. Now this is just sample.
type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type IUser interface {
	GetTableName() string
}

func (u User) GetTableName() string {
	return "users"
}
