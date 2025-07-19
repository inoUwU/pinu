package entities

// TODO: Refer to the database schema. Now this is just sample.
type User struct {
	ID    string
	Name  string
	Email string
}

func (u User) GetTableName() string {
	return "users"
}
