package entities

// TODO: Refer to the database schema. Now this is just sample.
type Category struct {
	Name string
}

func (c Category) GetTableName() string {
	return "categories"
}
